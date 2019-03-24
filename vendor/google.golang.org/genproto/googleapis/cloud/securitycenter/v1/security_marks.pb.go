// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1/security_marks.proto

package securitycenter // import "google.golang.org/genproto/googleapis/cloud/securitycenter/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// User specified security marks that are attached to the parent Cloud Security
// Command Center (Cloud SCC) resource. Security marks are scoped within a Cloud
// SCC organization -- they can be modified and viewed by all users who have
// proper permissions on the organization.
type SecurityMarks struct {
	// The relative resource name of the SecurityMarks. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Examples:
	// "organizations/123/assets/456/securityMarks"
	// "organizations/123/sources/456/findings/789/securityMarks".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Mutable user specified security marks belonging to the parent resource.
	// Constraints are as follows:
	//   - Keys and values are treated as case insensitive
	//   - Keys must be between 1 - 256 characters (inclusive)
	//   - Keys must be letters, numbers, underscores, or dashes
	//   - Values have leading and trailing whitespace trimmed, remaining
	//     characters must be between 1 - 4096 characters (inclusive)
	Marks                map[string]string `protobuf:"bytes,2,rep,name=marks,proto3" json:"marks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SecurityMarks) Reset()         { *m = SecurityMarks{} }
func (m *SecurityMarks) String() string { return proto.CompactTextString(m) }
func (*SecurityMarks) ProtoMessage()    {}
func (*SecurityMarks) Descriptor() ([]byte, []int) {
	return fileDescriptor_security_marks_0acd0834d3f61043, []int{0}
}
func (m *SecurityMarks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityMarks.Unmarshal(m, b)
}
func (m *SecurityMarks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityMarks.Marshal(b, m, deterministic)
}
func (dst *SecurityMarks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityMarks.Merge(dst, src)
}
func (m *SecurityMarks) XXX_Size() int {
	return xxx_messageInfo_SecurityMarks.Size(m)
}
func (m *SecurityMarks) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityMarks.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityMarks proto.InternalMessageInfo

func (m *SecurityMarks) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SecurityMarks) GetMarks() map[string]string {
	if m != nil {
		return m.Marks
	}
	return nil
}

func init() {
	proto.RegisterType((*SecurityMarks)(nil), "google.cloud.securitycenter.v1.SecurityMarks")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.securitycenter.v1.SecurityMarks.MarksEntry")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1/security_marks.proto", fileDescriptor_security_marks_0acd0834d3f61043)
}

var fileDescriptor_security_marks_0acd0834d3f61043 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x49, 0xe7, 0x04, 0x9f, 0x08, 0x12, 0x3c, 0x94, 0x21, 0x32, 0x76, 0xda, 0x29, 0x61,
	0xee, 0x52, 0xf4, 0x26, 0x78, 0x53, 0x91, 0x79, 0xf3, 0x22, 0xb1, 0x3e, 0x42, 0x59, 0xfb, 0x5e,
	0x49, 0xd3, 0x42, 0x3f, 0x94, 0xdf, 0x51, 0x96, 0x04, 0xa5, 0x07, 0xf5, 0x12, 0x5e, 0xfe, 0xc9,
	0xef, 0xff, 0x23, 0x81, 0xad, 0x65, 0xb6, 0x35, 0xea, 0xb2, 0xe6, 0xfe, 0x43, 0x77, 0x58, 0xf6,
	0xae, 0xf2, 0x63, 0x89, 0xe4, 0xd1, 0xe9, 0x61, 0xf3, 0x9d, 0xbc, 0x35, 0xc6, 0xed, 0x3b, 0xd5,
	0x3a, 0xf6, 0x2c, 0xaf, 0x22, 0xa4, 0x02, 0xa4, 0xa6, 0x90, 0x1a, 0x36, 0x8b, 0xcb, 0x54, 0x6a,
	0xda, 0x4a, 0x1b, 0x22, 0xf6, 0xc6, 0x57, 0x4c, 0x89, 0x5e, 0x7d, 0x0a, 0x38, 0x7b, 0x49, 0xcc,
	0xe3, 0xa1, 0x55, 0x4a, 0x38, 0x22, 0xd3, 0x60, 0x2e, 0x96, 0x62, 0x7d, 0xb2, 0x0b, 0xb3, 0x7c,
	0x82, 0x79, 0x50, 0xe6, 0xd9, 0x72, 0xb6, 0x3e, 0xbd, 0x2e, 0xd4, 0xdf, 0x4e, 0x35, 0x69, 0x54,
	0x61, 0xbd, 0x27, 0xef, 0xc6, 0x5d, 0xac, 0x59, 0x14, 0x00, 0x3f, 0xa1, 0x3c, 0x87, 0xd9, 0x1e,
	0xc7, 0x24, 0x3c, 0x8c, 0xf2, 0x02, 0xe6, 0x83, 0xa9, 0x7b, 0xcc, 0xb3, 0x90, 0xc5, 0xcd, 0x4d,
	0x56, 0x88, 0x3b, 0x0f, 0xab, 0x92, 0x9b, 0x7f, 0xfc, 0xcf, 0xe2, 0xf5, 0x21, 0xdd, 0xb0, 0x5c,
	0x1b, 0xb2, 0x8a, 0x9d, 0xd5, 0x16, 0x29, 0xbc, 0x59, 0xc7, 0x23, 0xd3, 0x56, 0xdd, 0x6f, 0x3f,
	0x7d, 0x3b, 0x4d, 0xde, 0x8f, 0x03, 0xb8, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xde, 0x72, 0xc1,
	0x00, 0xa1, 0x01, 0x00, 0x00,
}