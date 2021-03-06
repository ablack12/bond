// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/app_placeholder_field.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible values for App placeholder fields.
type AppPlaceholderFieldEnum_AppPlaceholderField int32

const (
	// Not specified.
	AppPlaceholderFieldEnum_UNSPECIFIED AppPlaceholderFieldEnum_AppPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	AppPlaceholderFieldEnum_UNKNOWN AppPlaceholderFieldEnum_AppPlaceholderField = 1
	// Data Type: INT64. The application store that the target application
	// belongs to. Valid values are: 1 = Apple iTunes Store; 2 = Google Play
	// Store.
	AppPlaceholderFieldEnum_STORE AppPlaceholderFieldEnum_AppPlaceholderField = 2
	// Data Type: STRING. The store-specific ID for the target application.
	AppPlaceholderFieldEnum_ID AppPlaceholderFieldEnum_AppPlaceholderField = 3
	// Data Type: STRING. The visible text displayed when the link is rendered
	// in an ad.
	AppPlaceholderFieldEnum_LINK_TEXT AppPlaceholderFieldEnum_AppPlaceholderField = 4
	// Data Type: STRING. The destination URL of the in-app link.
	AppPlaceholderFieldEnum_URL AppPlaceholderFieldEnum_AppPlaceholderField = 5
	// Data Type: URL_LIST. Final URLs for the in-app link when using Upgraded
	// URLs.
	AppPlaceholderFieldEnum_FINAL_URLS AppPlaceholderFieldEnum_AppPlaceholderField = 6
	// Data Type: URL_LIST. Final Mobile URLs for the in-app link when using
	// Upgraded URLs.
	AppPlaceholderFieldEnum_FINAL_MOBILE_URLS AppPlaceholderFieldEnum_AppPlaceholderField = 7
	// Data Type: URL. Tracking template for the in-app link when using Upgraded
	// URLs.
	AppPlaceholderFieldEnum_TRACKING_URL AppPlaceholderFieldEnum_AppPlaceholderField = 8
	// Data Type: STRING. Final URL suffix for the in-app link when using
	// parallel tracking.
	AppPlaceholderFieldEnum_FINAL_URL_SUFFIX AppPlaceholderFieldEnum_AppPlaceholderField = 9
)

var AppPlaceholderFieldEnum_AppPlaceholderField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "STORE",
	3: "ID",
	4: "LINK_TEXT",
	5: "URL",
	6: "FINAL_URLS",
	7: "FINAL_MOBILE_URLS",
	8: "TRACKING_URL",
	9: "FINAL_URL_SUFFIX",
}
var AppPlaceholderFieldEnum_AppPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":       0,
	"UNKNOWN":           1,
	"STORE":             2,
	"ID":                3,
	"LINK_TEXT":         4,
	"URL":               5,
	"FINAL_URLS":        6,
	"FINAL_MOBILE_URLS": 7,
	"TRACKING_URL":      8,
	"FINAL_URL_SUFFIX":  9,
}

func (x AppPlaceholderFieldEnum_AppPlaceholderField) String() string {
	return proto.EnumName(AppPlaceholderFieldEnum_AppPlaceholderField_name, int32(x))
}
func (AppPlaceholderFieldEnum_AppPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_app_placeholder_field_403a92f3c1d63244, []int{0, 0}
}

// Values for App placeholder fields.
type AppPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppPlaceholderFieldEnum) Reset()         { *m = AppPlaceholderFieldEnum{} }
func (m *AppPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*AppPlaceholderFieldEnum) ProtoMessage()    {}
func (*AppPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_placeholder_field_403a92f3c1d63244, []int{0}
}
func (m *AppPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *AppPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (dst *AppPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppPlaceholderFieldEnum.Merge(dst, src)
}
func (m *AppPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_AppPlaceholderFieldEnum.Size(m)
}
func (m *AppPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AppPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AppPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AppPlaceholderFieldEnum)(nil), "google.ads.googleads.v0.enums.AppPlaceholderFieldEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AppPlaceholderFieldEnum_AppPlaceholderField", AppPlaceholderFieldEnum_AppPlaceholderField_name, AppPlaceholderFieldEnum_AppPlaceholderField_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/app_placeholder_field.proto", fileDescriptor_app_placeholder_field_403a92f3c1d63244)
}

var fileDescriptor_app_placeholder_field_403a92f3c1d63244 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xdd, 0x6e, 0xa2, 0x40,
	0x1c, 0xc5, 0x17, 0x5c, 0x75, 0x1d, 0xf7, 0x63, 0x76, 0x76, 0x37, 0xbb, 0x37, 0x5e, 0xe8, 0x03,
	0x0c, 0x24, 0x7b, 0xd5, 0xe9, 0xd5, 0xa0, 0x60, 0x26, 0x52, 0x24, 0x7c, 0x58, 0xd3, 0x90, 0x10,
	0x2a, 0x94, 0x9a, 0x20, 0x43, 0x9c, 0xea, 0x03, 0xf5, 0xb2, 0x49, 0x1f, 0xa4, 0x7d, 0x94, 0x26,
	0x7d, 0x87, 0x06, 0xa8, 0xf4, 0xc6, 0xf6, 0x86, 0x1c, 0xfe, 0xe7, 0xfc, 0x26, 0xf3, 0x3f, 0x03,
	0x4e, 0x52, 0xce, 0xd3, 0x2c, 0x51, 0xa2, 0x58, 0x28, 0xb5, 0x2c, 0xd5, 0x5e, 0x55, 0x92, 0x7c,
	0xb7, 0x11, 0x4a, 0x54, 0x14, 0x61, 0x91, 0x45, 0xab, 0xe4, 0x9a, 0x67, 0x71, 0xb2, 0x0d, 0xaf,
	0xd6, 0x49, 0x16, 0xe3, 0x62, 0xcb, 0x6f, 0x38, 0x1a, 0xd4, 0x79, 0x1c, 0xc5, 0x02, 0x37, 0x28,
	0xde, 0xab, 0xb8, 0x42, 0x47, 0x0f, 0x12, 0xf8, 0x4b, 0x8b, 0xc2, 0x7e, 0xa3, 0x8d, 0x12, 0xd6,
	0xf3, 0xdd, 0x66, 0x74, 0x2f, 0x81, 0x5f, 0x47, 0x3c, 0xf4, 0x03, 0xf4, 0x7d, 0xcb, 0xb5, 0xf5,
	0x31, 0x33, 0x98, 0x3e, 0x81, 0x9f, 0x50, 0x1f, 0x74, 0x7d, 0x6b, 0x66, 0xcd, 0xcf, 0x2d, 0x28,
	0xa1, 0x1e, 0x68, 0xbb, 0xde, 0xdc, 0xd1, 0xa1, 0x8c, 0x3a, 0x40, 0x66, 0x13, 0xd8, 0x42, 0xdf,
	0x40, 0xcf, 0x64, 0xd6, 0x2c, 0xf4, 0xf4, 0xa5, 0x07, 0x3f, 0xa3, 0x2e, 0x68, 0xf9, 0x8e, 0x09,
	0xdb, 0xe8, 0x3b, 0x00, 0x06, 0xb3, 0xa8, 0x19, 0xfa, 0x8e, 0xe9, 0xc2, 0x0e, 0xfa, 0x03, 0x7e,
	0xd6, 0xff, 0x67, 0x73, 0x8d, 0x99, 0x7a, 0x3d, 0xee, 0x22, 0x08, 0xbe, 0x7a, 0x0e, 0x1d, 0xcf,
	0x98, 0x35, 0x2d, 0x47, 0xf0, 0x0b, 0xfa, 0x0d, 0x60, 0x03, 0x86, 0xae, 0x6f, 0x18, 0x6c, 0x09,
	0x7b, 0xda, 0xb3, 0x04, 0x86, 0x2b, 0xbe, 0xc1, 0x1f, 0x6e, 0xac, 0xfd, 0x3b, 0xb2, 0x92, 0x5d,
	0x56, 0x65, 0x4b, 0x17, 0xda, 0x2b, 0x9a, 0xf2, 0x2c, 0xca, 0x53, 0xcc, 0xb7, 0xa9, 0x92, 0x26,
	0x79, 0x55, 0xe4, 0xa1, 0xf7, 0x62, 0x2d, 0xde, 0x79, 0x86, 0xd3, 0xea, 0x7b, 0x2b, 0xb7, 0xa6,
	0x94, 0xde, 0xc9, 0x83, 0x69, 0x7d, 0x14, 0x8d, 0x05, 0xae, 0x65, 0xa9, 0x16, 0x2a, 0x2e, 0xab,
	0x15, 0x8f, 0x07, 0x3f, 0xa0, 0xb1, 0x08, 0x1a, 0x3f, 0x58, 0xa8, 0x41, 0xe5, 0x3f, 0xc9, 0xc3,
	0x7a, 0x48, 0x08, 0x8d, 0x05, 0x21, 0x4d, 0x82, 0x90, 0x85, 0x4a, 0x48, 0x95, 0xb9, 0xec, 0x54,
	0x17, 0xfb, 0xff, 0x12, 0x00, 0x00, 0xff, 0xff, 0x11, 0x1c, 0x65, 0x10, 0x1e, 0x02, 0x00, 0x00,
}
