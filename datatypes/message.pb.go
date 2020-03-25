// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package datatypes

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

type Arguments struct {
	Fnname               string   `protobuf:"bytes,1,opt,name=fnname,proto3" json:"fnname,omitempty"`
	Txid                 string   `protobuf:"bytes,2,opt,name=txid,proto3" json:"txid,omitempty"`
	Channelid            string   `protobuf:"bytes,3,opt,name=channelid,proto3" json:"channelid,omitempty"`
	Args                 []string `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Arguments) Reset()         { *m = Arguments{} }
func (m *Arguments) String() string { return proto.CompactTextString(m) }
func (*Arguments) ProtoMessage()    {}
func (*Arguments) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Arguments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Arguments.Unmarshal(m, b)
}
func (m *Arguments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Arguments.Marshal(b, m, deterministic)
}
func (m *Arguments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Arguments.Merge(m, src)
}
func (m *Arguments) XXX_Size() int {
	return xxx_messageInfo_Arguments.Size(m)
}
func (m *Arguments) XXX_DiscardUnknown() {
	xxx_messageInfo_Arguments.DiscardUnknown(m)
}

var xxx_messageInfo_Arguments proto.InternalMessageInfo

func (m *Arguments) GetFnname() string {
	if m != nil {
		return m.Fnname
	}
	return ""
}

func (m *Arguments) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *Arguments) GetChannelid() string {
	if m != nil {
		return m.Channelid
	}
	return ""
}

func (m *Arguments) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type Return struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Return) Reset()         { *m = Return{} }
func (m *Return) String() string { return proto.CompactTextString(m) }
func (*Return) ProtoMessage()    {}
func (*Return) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *Return) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Return.Unmarshal(m, b)
}
func (m *Return) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Return.Marshal(b, m, deterministic)
}
func (m *Return) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Return.Merge(m, src)
}
func (m *Return) XXX_Size() int {
	return xxx_messageInfo_Return.Size(m)
}
func (m *Return) XXX_DiscardUnknown() {
	xxx_messageInfo_Return.DiscardUnknown(m)
}

var xxx_messageInfo_Return proto.InternalMessageInfo

func (m *Return) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Return) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Arguments)(nil), "datatypes.Arguments")
	proto.RegisterType((*Return)(nil), "datatypes.Return")
}

func init() {
	proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd)
}

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x49, 0x2c, 0x49, 0x2c,
	0xa9, 0x2c, 0x48, 0x2d, 0x56, 0xca, 0xe4, 0xe2, 0x74, 0x2c, 0x4a, 0x2f, 0xcd, 0x4d, 0xcd, 0x2b,
	0x29, 0x16, 0x12, 0xe3, 0x62, 0x4b, 0xcb, 0xcb, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x2a, 0x32, 0x53, 0x24, 0x98, 0xc0, 0xa2,
	0x60, 0xb6, 0x90, 0x0c, 0x17, 0x67, 0x72, 0x46, 0x62, 0x5e, 0x5e, 0x6a, 0x4e, 0x66, 0x8a, 0x04,
	0x33, 0x58, 0x02, 0x21, 0x00, 0xd2, 0x91, 0x58, 0x94, 0x5e, 0x2c, 0xc1, 0xa2, 0xc0, 0x0c, 0xd2,
	0x01, 0x62, 0x2b, 0x19, 0x70, 0xb1, 0x05, 0xa5, 0x96, 0x94, 0x16, 0xe5, 0x81, 0x64, 0x93, 0xf3,
	0x53, 0x20, 0xb6, 0xb0, 0x06, 0x81, 0xd9, 0x20, 0x31, 0x90, 0xab, 0x60, 0x76, 0x80, 0xd8, 0x49,
	0x6c, 0x60, 0xe7, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x2f, 0xd2, 0xfb, 0xbf, 0x00,
	0x00, 0x00,
}
