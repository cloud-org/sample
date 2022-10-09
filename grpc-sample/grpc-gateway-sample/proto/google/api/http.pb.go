// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/http.proto

package google_api

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Defines the HTTP configuration for a service. It contains a list of
// [HttpRule][google.api.HttpRule], each specifying the mapping of an RPC method
// to one or more HTTP REST API methods.
type Http struct {
	// A list of HTTP rules for configuring the HTTP REST API methods.
	Rules                []*HttpRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Http) Reset()         { *m = Http{} }
func (m *Http) String() string { return proto.CompactTextString(m) }
func (*Http) ProtoMessage()    {}
func (*Http) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9994be407cdcc9, []int{0}
}

func (m *Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Http.Unmarshal(m, b)
}
func (m *Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Http.Marshal(b, m, deterministic)
}
func (m *Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Http.Merge(m, src)
}
func (m *Http) XXX_Size() int {
	return xxx_messageInfo_Http.Size(m)
}
func (m *Http) XXX_DiscardUnknown() {
	xxx_messageInfo_Http.DiscardUnknown(m)
}

var xxx_messageInfo_Http proto.InternalMessageInfo

func (m *Http) GetRules() []*HttpRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Use CustomHttpPattern to specify any HTTP method that is not included in the
// `pattern` field, such as HEAD, or "*" to leave the HTTP method unspecified for
// a given URL path rule. The wild-card rule is useful for services that provide
// content to Web (HTML) clients.
type HttpRule struct {
	// Selects methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Determines the URL pattern is matched by this rules. This pattern can be
	// used with any of the {get|put|post|delete|patch} methods. A custom method
	// can be defined using the 'custom' field.
	//
	// Types that are valid to be assigned to Pattern:
	//	*HttpRule_Get
	//	*HttpRule_Put
	//	*HttpRule_Post
	//	*HttpRule_Delete
	//	*HttpRule_Patch
	//	*HttpRule_Custom
	Pattern isHttpRule_Pattern `protobuf_oneof:"pattern"`
	// The name of the request field whose value is mapped to the HTTP body, or
	// `*` for mapping all fields not captured by the path pattern to the HTTP
	// body. NOTE: the referred field must not be a repeated field.
	Body string `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	// Additional HTTP bindings for the selector. Nested bindings must
	// not contain an `additional_bindings` field themselves (that is,
	// the nesting may only be one level deep).
	AdditionalBindings   []*HttpRule `protobuf:"bytes,11,rep,name=additional_bindings,json=additionalBindings,proto3" json:"additional_bindings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HttpRule) Reset()         { *m = HttpRule{} }
func (m *HttpRule) String() string { return proto.CompactTextString(m) }
func (*HttpRule) ProtoMessage()    {}
func (*HttpRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9994be407cdcc9, []int{1}
}

func (m *HttpRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpRule.Unmarshal(m, b)
}
func (m *HttpRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpRule.Marshal(b, m, deterministic)
}
func (m *HttpRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpRule.Merge(m, src)
}
func (m *HttpRule) XXX_Size() int {
	return xxx_messageInfo_HttpRule.Size(m)
}
func (m *HttpRule) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpRule.DiscardUnknown(m)
}

var xxx_messageInfo_HttpRule proto.InternalMessageInfo

func (m *HttpRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

type isHttpRule_Pattern interface {
	isHttpRule_Pattern()
}

type HttpRule_Get struct {
	Get string `protobuf:"bytes,2,opt,name=get,proto3,oneof"`
}

type HttpRule_Put struct {
	Put string `protobuf:"bytes,3,opt,name=put,proto3,oneof"`
}

type HttpRule_Post struct {
	Post string `protobuf:"bytes,4,opt,name=post,proto3,oneof"`
}

type HttpRule_Delete struct {
	Delete string `protobuf:"bytes,5,opt,name=delete,proto3,oneof"`
}

type HttpRule_Patch struct {
	Patch string `protobuf:"bytes,6,opt,name=patch,proto3,oneof"`
}

type HttpRule_Custom struct {
	Custom *CustomHttpPattern `protobuf:"bytes,8,opt,name=custom,proto3,oneof"`
}

func (*HttpRule_Get) isHttpRule_Pattern() {}

func (*HttpRule_Put) isHttpRule_Pattern() {}

func (*HttpRule_Post) isHttpRule_Pattern() {}

func (*HttpRule_Delete) isHttpRule_Pattern() {}

func (*HttpRule_Patch) isHttpRule_Pattern() {}

func (*HttpRule_Custom) isHttpRule_Pattern() {}

func (m *HttpRule) GetPattern() isHttpRule_Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *HttpRule) GetGet() string {
	if x, ok := m.GetPattern().(*HttpRule_Get); ok {
		return x.Get
	}
	return ""
}

func (m *HttpRule) GetPut() string {
	if x, ok := m.GetPattern().(*HttpRule_Put); ok {
		return x.Put
	}
	return ""
}

func (m *HttpRule) GetPost() string {
	if x, ok := m.GetPattern().(*HttpRule_Post); ok {
		return x.Post
	}
	return ""
}

func (m *HttpRule) GetDelete() string {
	if x, ok := m.GetPattern().(*HttpRule_Delete); ok {
		return x.Delete
	}
	return ""
}

func (m *HttpRule) GetPatch() string {
	if x, ok := m.GetPattern().(*HttpRule_Patch); ok {
		return x.Patch
	}
	return ""
}

func (m *HttpRule) GetCustom() *CustomHttpPattern {
	if x, ok := m.GetPattern().(*HttpRule_Custom); ok {
		return x.Custom
	}
	return nil
}

func (m *HttpRule) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *HttpRule) GetAdditionalBindings() []*HttpRule {
	if m != nil {
		return m.AdditionalBindings
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpRule) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpRule_Get)(nil),
		(*HttpRule_Put)(nil),
		(*HttpRule_Post)(nil),
		(*HttpRule_Delete)(nil),
		(*HttpRule_Patch)(nil),
		(*HttpRule_Custom)(nil),
	}
}

// A custom pattern is used for defining custom HTTP verb.
type CustomHttpPattern struct {
	// The name of this custom HTTP verb.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// The path matched by this custom verb.
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomHttpPattern) Reset()         { *m = CustomHttpPattern{} }
func (m *CustomHttpPattern) String() string { return proto.CompactTextString(m) }
func (*CustomHttpPattern) ProtoMessage()    {}
func (*CustomHttpPattern) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9994be407cdcc9, []int{2}
}

func (m *CustomHttpPattern) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomHttpPattern.Unmarshal(m, b)
}
func (m *CustomHttpPattern) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomHttpPattern.Marshal(b, m, deterministic)
}
func (m *CustomHttpPattern) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomHttpPattern.Merge(m, src)
}
func (m *CustomHttpPattern) XXX_Size() int {
	return xxx_messageInfo_CustomHttpPattern.Size(m)
}
func (m *CustomHttpPattern) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomHttpPattern.DiscardUnknown(m)
}

var xxx_messageInfo_CustomHttpPattern proto.InternalMessageInfo

func (m *CustomHttpPattern) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *CustomHttpPattern) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*Http)(nil), "google.api.Http")
	proto.RegisterType((*HttpRule)(nil), "google.api.HttpRule")
	proto.RegisterType((*CustomHttpPattern)(nil), "google.api.CustomHttpPattern")
}

func init() {
	proto.RegisterFile("google/api/http.proto", fileDescriptor_ff9994be407cdcc9)
}

var fileDescriptor_ff9994be407cdcc9 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0xbf, 0x69, 0xd3, 0xb4, 0x3d, 0x85, 0x0f, 0x3c, 0x56, 0x19, 0x04, 0x21, 0x74, 0x55,
	0x5c, 0xa4, 0x50, 0x17, 0x2e, 0xdc, 0x45, 0x84, 0x2e, 0x4b, 0x6e, 0x40, 0xa6, 0xc9, 0x90, 0x0c,
	0xa6, 0x99, 0x21, 0x39, 0x59, 0x78, 0x61, 0xde, 0x9b, 0x4b, 0x99, 0x9f, 0xda, 0x82, 0xe0, 0xee,
	0xbc, 0xcf, 0xbc, 0x73, 0x7e, 0xe1, 0xa6, 0xd2, 0xba, 0x6a, 0xe4, 0x46, 0x18, 0xb5, 0xa9, 0x89,
	0x4c, 0x6a, 0x3a, 0x4d, 0x1a, 0xc1, 0xe3, 0x54, 0x18, 0xb5, 0xda, 0x42, 0xb4, 0x23, 0x32, 0xf8,
	0x00, 0x93, 0x6e, 0x68, 0x64, 0xcf, 0x59, 0x32, 0x5e, 0x2f, 0xb6, 0xcb, 0xf4, 0xec, 0x49, 0xad,
	0x21, 0x1f, 0x1a, 0x99, 0x7b, 0xcb, 0xea, 0x73, 0x04, 0xb3, 0x13, 0xc3, 0x3b, 0x98, 0xf5, 0xb2,
	0x91, 0x05, 0xe9, 0x8e, 0xb3, 0x84, 0xad, 0xe7, 0xf9, 0x8f, 0x46, 0x84, 0x71, 0x25, 0x89, 0x8f,
	0x2c, 0xde, 0xfd, 0xcb, 0xad, 0xb0, 0xcc, 0x0c, 0xc4, 0xc7, 0x27, 0x66, 0x06, 0xc2, 0x25, 0x44,
	0x46, 0xf7, 0xc4, 0xa3, 0x00, 0x9d, 0x42, 0x0e, 0x71, 0x29, 0x1b, 0x49, 0x92, 0x4f, 0x02, 0x0f,
	0x1a, 0x6f, 0x61, 0x62, 0x04, 0x15, 0x35, 0x8f, 0xc3, 0x83, 0x97, 0xf8, 0x04, 0x71, 0x31, 0xf4,
	0xa4, 0x8f, 0x7c, 0x96, 0xb0, 0xf5, 0x62, 0x7b, 0x7f, 0x39, 0xc5, 0x8b, 0x7b, 0xb1, 0x7d, 0xef,
	0x05, 0x91, 0xec, 0x5a, 0x9b, 0xd0, 0xdb, 0x11, 0x21, 0x3a, 0xe8, 0xf2, 0x83, 0x4f, 0xdd, 0x00,
	0x2e, 0xc6, 0x57, 0xb8, 0x16, 0x65, 0xa9, 0x48, 0xe9, 0x56, 0x34, 0x6f, 0x07, 0xd5, 0x96, 0xaa,
	0xad, 0x7a, 0xbe, 0xf8, 0x63, 0x3f, 0x78, 0xfe, 0x90, 0x05, 0x7f, 0x36, 0x87, 0xa9, 0xf1, 0xf5,
	0x56, 0xcf, 0x70, 0xf5, 0xab, 0x09, 0x5b, 0xfa, 0x5d, 0xb5, 0x65, 0xd8, 0x9d, 0x8b, 0x2d, 0x33,
	0x82, 0x6a, 0xbf, 0xb8, 0xdc, 0xc5, 0x59, 0x02, 0xff, 0x0b, 0x7d, 0xbc, 0x28, 0x9b, 0xcd, 0x5d,
	0x1a, 0x7b, 0xd1, 0x3d, 0xfb, 0x62, 0xec, 0x10, 0xbb, 0xeb, 0x3e, 0x7e, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x12, 0x3f, 0x0a, 0x0d, 0xf6, 0x01, 0x00, 0x00,
}
