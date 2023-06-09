// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/server/server.proto

package server

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type TestTask struct {
	Data                 string            `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Kind                 string            `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	CaseName             string            `protobuf:"bytes,3,opt,name=caseName,proto3" json:"caseName,omitempty"`
	Level                string            `protobuf:"bytes,4,opt,name=level,proto3" json:"level,omitempty"`
	Env                  map[string]string `protobuf:"bytes,5,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TestTask) Reset()         { *m = TestTask{} }
func (m *TestTask) String() string { return proto.CompactTextString(m) }
func (*TestTask) ProtoMessage()    {}
func (*TestTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_36fb7b77b8f76c98, []int{0}
}

func (m *TestTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestTask.Unmarshal(m, b)
}
func (m *TestTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestTask.Marshal(b, m, deterministic)
}
func (m *TestTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestTask.Merge(m, src)
}
func (m *TestTask) XXX_Size() int {
	return xxx_messageInfo_TestTask.Size(m)
}
func (m *TestTask) XXX_DiscardUnknown() {
	xxx_messageInfo_TestTask.DiscardUnknown(m)
}

var xxx_messageInfo_TestTask proto.InternalMessageInfo

func (m *TestTask) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *TestTask) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *TestTask) GetCaseName() string {
	if m != nil {
		return m.CaseName
	}
	return ""
}

func (m *TestTask) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *TestTask) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_36fb7b77b8f76c98, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_36fb7b77b8f76c98, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TestTask)(nil), "server.TestTask")
	proto.RegisterMapType((map[string]string)(nil), "server.TestTask.EnvEntry")
	proto.RegisterType((*HelloReply)(nil), "server.HelloReply")
	proto.RegisterType((*Empty)(nil), "server.Empty")
}

func init() {
	proto.RegisterFile("pkg/server/server.proto", fileDescriptor_36fb7b77b8f76c98)
}

var fileDescriptor_36fb7b77b8f76c98 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xdf, 0x34, 0xfd, 0xf7, 0x8e, 0x08, 0x65, 0x10, 0x8c, 0x3d, 0x95, 0x9c, 0x0a, 0xda,
	0x04, 0x2b, 0x88, 0x88, 0x27, 0xa1, 0xe8, 0xc9, 0x43, 0x28, 0x1e, 0xbc, 0x6d, 0xdb, 0x21, 0xc6,
	0x6c, 0x36, 0x61, 0x77, 0x13, 0xcc, 0x47, 0xf4, 0x5b, 0xc9, 0x6e, 0xb2, 0x16, 0xc4, 0xd3, 0xce,
	0xef, 0xd9, 0x99, 0x79, 0x1e, 0x18, 0x38, 0xaf, 0xf2, 0x34, 0x56, 0x24, 0x1b, 0x92, 0xfd, 0x13,
	0x55, 0xb2, 0xd4, 0x25, 0x8e, 0x3b, 0x0a, 0xbf, 0x3c, 0x98, 0x6e, 0x49, 0xe9, 0x2d, 0x53, 0x39,
	0x22, 0x0c, 0x0f, 0x4c, 0xb3, 0xc0, 0x5b, 0x78, 0xcb, 0xff, 0x89, 0xad, 0x8d, 0x96, 0x67, 0xe2,
	0x10, 0x0c, 0x3a, 0xcd, 0xd4, 0x38, 0x87, 0xe9, 0x9e, 0x29, 0x7a, 0x61, 0x05, 0x05, 0xbe, 0xd5,
	0x7f, 0x18, 0xcf, 0x60, 0xc4, 0xa9, 0x21, 0x1e, 0x0c, 0xed, 0x47, 0x07, 0x78, 0x09, 0x3e, 0x89,
	0x26, 0x18, 0x2d, 0xfc, 0xe5, 0xc9, 0xfa, 0x22, 0xea, 0xa3, 0x38, 0xe3, 0x68, 0x23, 0x9a, 0x8d,
	0xd0, 0xb2, 0x4d, 0x4c, 0xd7, 0xfc, 0x16, 0xa6, 0x4e, 0xc0, 0x19, 0xf8, 0x39, 0xb5, 0x7d, 0x22,
	0x53, 0x1a, 0x83, 0x86, 0xf1, 0x9a, 0xfa, 0x44, 0x1d, 0xdc, 0x0f, 0xee, 0xbc, 0xf0, 0x01, 0xe0,
	0x99, 0x38, 0x2f, 0x13, 0xaa, 0x78, 0x8b, 0x01, 0x4c, 0x0a, 0x52, 0x8a, 0xa5, 0xd4, 0x4f, 0x3b,
	0x34, 0x1b, 0x48, 0xca, 0x52, 0xba, 0x0d, 0x16, 0xc2, 0x09, 0x8c, 0x36, 0x45, 0xa5, 0xdb, 0xf5,
	0x07, 0x8c, 0x93, 0x5a, 0x08, 0x92, 0xb8, 0x02, 0x3f, 0xa9, 0x05, 0xce, 0x7e, 0xe7, 0x9d, 0xa3,
	0x53, 0x8e, 0x7e, 0xe1, 0x3f, 0xbc, 0x06, 0x78, 0x22, 0xfd, 0x4a, 0x52, 0x65, 0xa5, 0xc0, 0x53,
	0xd7, 0x63, 0xb7, 0xfe, 0x3d, 0xf2, 0x18, 0xbd, 0x5d, 0xa5, 0x99, 0x7e, 0xaf, 0x77, 0xd1, 0xbe,
	0x2c, 0x62, 0x9e, 0x89, 0xfa, 0x53, 0xd5, 0x92, 0x44, 0xcc, 0xaa, 0x6c, 0xa5, 0x49, 0xe9, 0x4c,
	0xa4, 0xf1, 0xf1, 0x86, 0xbb, 0xb1, 0xbd, 0xde, 0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30,
	0xde, 0x43, 0x18, 0xd8, 0x01, 0x00, 0x00,
}
