// Code generated by protoc-gen-go.
// source: msg.proto
// DO NOT EDIT!

/*
Package demo is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	Msg
*/
package demo

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

// go中导出结构体、方法必须大写
type Msg struct {
	MsgType          *int32  `protobuf:"varint,1,req,name=MsgType" json:"MsgType,omitempty"`
	MsgInfo          *string `protobuf:"bytes,2,req,name=MsgInfo" json:"MsgInfo,omitempty"`
	MsgFrom          *string `protobuf:"bytes,3,req,name=MsgFrom" json:"MsgFrom,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Msg) GetMsgType() int32 {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return 0
}

func (m *Msg) GetMsgInfo() string {
	if m != nil && m.MsgInfo != nil {
		return *m.MsgInfo
	}
	return ""
}

func (m *Msg) GetMsgFrom() string {
	if m != nil && m.MsgFrom != nil {
		return *m.MsgFrom
	}
	return ""
}

func init() {
	proto.RegisterType((*Msg)(nil), "demo.Msg")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 87 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x49, 0xcd, 0xcd, 0x57, 0xb2, 0xe0, 0x62, 0xf6,
	0x2d, 0x4e, 0x17, 0xe2, 0xe7, 0x62, 0xf7, 0x2d, 0x4e, 0x0f, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54,
	0x60, 0xd2, 0x60, 0x85, 0x0a, 0x78, 0xe6, 0xa5, 0xe5, 0x4b, 0x30, 0x29, 0x30, 0x69, 0x70, 0x42,
	0x05, 0xdc, 0x8a, 0xf2, 0x73, 0x25, 0x98, 0x41, 0x02, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f,
	0xec, 0x33, 0xee, 0x4b, 0x00, 0x00, 0x00,
}