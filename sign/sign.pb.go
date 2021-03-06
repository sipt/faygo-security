// Code generated by protoc-gen-go.
// source: sign.proto
// DO NOT EDIT!

/*
Package sign is a generated protocol buffer package.

It is generated from these files:
	sign.proto

It has these top-level messages:
	ClientInfo
*/
package sign

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

type ClientInfo struct {
	Secret   string `protobuf:"bytes,1,opt,name=secret" json:"secret,omitempty"`
	IsTest   bool   `protobuf:"varint,2,opt,name=isTest" json:"isTest,omitempty"`
	Platform string `protobuf:"bytes,3,opt,name=platform" json:"platform,omitempty"`
}

func (m *ClientInfo) Reset()                    { *m = ClientInfo{} }
func (m *ClientInfo) String() string            { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()               {}
func (*ClientInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientInfo) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *ClientInfo) GetIsTest() bool {
	if m != nil {
		return m.IsTest
	}
	return false
}

func (m *ClientInfo) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientInfo)(nil), "ClientInfo")
}

func init() { proto.RegisterFile("sign.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 108 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xce, 0x4c, 0xcf,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x8a, 0xe0, 0xe2, 0x72, 0xce, 0xc9, 0x4c, 0xcd, 0x2b,
	0xf1, 0xcc, 0x4b, 0xcb, 0x17, 0x12, 0xe3, 0x62, 0x2b, 0x4e, 0x4d, 0x2e, 0x4a, 0x2d, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x40, 0xe2, 0x99, 0xc5, 0x21, 0xa9, 0xc5, 0x25, 0x12,
	0x4c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x50, 0x9e, 0x90, 0x14, 0x17, 0x47, 0x41, 0x4e, 0x62, 0x49,
	0x5a, 0x7e, 0x51, 0xae, 0x04, 0x33, 0x58, 0x07, 0x9c, 0x9f, 0xc4, 0x06, 0xb6, 0xc0, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xe4, 0xad, 0xf4, 0x61, 0x6e, 0x00, 0x00, 0x00,
}
