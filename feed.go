package bond

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type ArtifactsFeed struct {
	Versions []*ArtifactVersion

	mutex sync.RWMutex
	table map[string]*ArtifactVersion
	dir   string
	path  string
}

type ArtifactDownload struct {
	Arch    MongoDBArch
	Edition MongoDBEdition
	Target  string
	Archive struct {
		Sha1   string
		Sha256 string
		Url    string
	}
	Msi      string
	Packages []string
}

const day = time.Hour * 24

func NewArtifactsFeed(path string) (*ArtifactsFeed, error) {
	f := &ArtifactsFeed{
		table: make(map[string]*ArtifactVersion),
		path:  path,
	}

	if path == "" {
		// no value for feed, let's write it to the tempDir
		tmpDir, err := ioutil.TempDir("", "mongodb-downloads")
		if err != nil {
			return nil, err
		}

		f.dir = tmpDir
		f.path = filepath.Join(tmpDir, "full.json")
	} else if strings.HasSuffix(path, ".json") {
		f.dir = filepath.Dir(path)
		f.path = path
	} else {
		f.dir = path
		f.path = filepath.Join(f.dir, "full.json")
	}

	if stat, err := os.Stat(f.path); !os.IsNotExist(err) && stat.IsDir() {
		// if the thing we think should be the json file
		// exists but isn't a file (i.e. directory,) then this
		// should be an error.
		return nil, errors.Errorf("path %s not a json file  directory", path)
	}

	return f, nil
}

func (feed *ArtifactsFeed) Populate() error {
	data, err := CacheDownload(day*2, "http://downloads.mongodb.org/full.json", feed.path, false)

	if err != nil {
		return errors.Wrap(err, "problem getting feed data")
	}

	if err = feed.Reload(data); err != nil {
		return errors.Wrap(err, "problem reloading feed")
	}

	return nil
}

func (feed *ArtifactsFeed) Reload(data []byte) error {
	// file exists, remove it if it's more than 48 hours old.
	feed.mutex.Lock()
	defer feed.mutex.Unlock()

	err := json.Unmarshal(data, feed)
	if err != nil {
		return errors.Wrap(err, "problem converting data from json")
	}

	// this is a reload rather than a new load, and we shoiuld
	if len(feed.table) > 0 {
		feed.table = make(map[string]*ArtifactVersion)
	}

	for _, version := range feed.Versions {
		feed.table[version.Version] = version
		version.refresh()
	}

	return err
}

func (feed *ArtifactsFeed) GetVersion(release string) (*ArtifactVersion, bool) {
	feed.mutex.RLock()
	defer feed.mutex.RUnlock()

	version, ok := feed.table[release]
	return version, ok
}

func (feed *ArtifactsFeed) GetLatestArchive(series string, edition MongoDBEdition, arch MongoDBArch, target string) (string, error) {
	if len(series) != 3 || string(series[1]) != "." {
		return "", errors.Errorf("series '%s' is not a valid version series", series)
	}

	version, ok := feed.GetVersion(series + ".0")
	if !ok {
		return "", errors.Errorf("there is no .0 release for series '%s' in the feed", series)
	}

	dl, err := version.GetDownload(edition, arch, target)
	if err != nil {
		return "", errors.Wrapf(err, "problem fetching download information for series '%s'", series)
	}

	return strings.Replace(dl.Archive.Url, series+".0", series+"-latest", -1), nil
}

func (feed *ArtifactsFeed) GetArchives(releases []string, edition MongoDBEdition, arch MongoDBArch, target string) (<-chan string, <-chan []error) {
	output := make(chan string)
	errReturn := make(chan []error)

	go func() {
		var errs []error
		for _, rel := range releases {
			// this is a series, have to handle it differently
			if len(rel) == 3 {
				url, err := feed.GetLatestArchive(rel, edition, arch, target)
				if err != nil {
					errs = append(errs, err)
					continue
				}
				output <- url
				continue
			}

			version, ok := feed.GetVersion(rel)
			if !ok {
				errs = append(errs, errors.Errorf("no version defined for %s", rel))
				continue
			}
			dl, err := version.GetDownload(edition, arch, target)
			if err != nil {
				errs = append(errs, err)
				continue
			}

			output <- dl.Archive.Url
		}
		close(output)
		errReturn <- errs
		close(errReturn)
	}()

	return output, errReturn
}