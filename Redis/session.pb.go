// Code generated by protoc-gen-go.
// source: session.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	session.proto

It has these top-level messages:
	Session
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Session struct {
	UserName string `protobuf:"bytes,1,opt,name=userName" json:"userName,omitempty"`
	LoggedIn bool   `protobuf:"varint,2,opt,name=loggedIn" json:"loggedIn,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Session)(nil), "main.session")
}

var fileDescriptor0 = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2e,
	0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53,
	0x72, 0xe4, 0x62, 0x87, 0x0a, 0x0b, 0x49, 0x71, 0x71, 0x94, 0x16, 0xa7, 0x16, 0xf9, 0x25, 0xe6,
	0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x20, 0xb9, 0x9c, 0xfc, 0xf4, 0xf4,
	0xd4, 0x14, 0xcf, 0x3c, 0x09, 0x26, 0xa0, 0x1c, 0x47, 0x10, 0x9c, 0x9f, 0xc4, 0x06, 0x36, 0xcf,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xe0, 0xb7, 0xd3, 0x60, 0x00, 0x00, 0x00,
}