// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_http2_req.proto

package clientH2Req

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

type ClientMsg struct {
	//用户名
	UserName *string `protobuf:"bytes,1,req,name=userName" json:"userName,omitempty"`
	//请求类型
	ActionCode *uint32 `protobuf:"varint,2,req,name=actionCode" json:"actionCode,omitempty"`
	//请求时间（用于与服务器进行校对）
	ClntReqYear          *int32   `protobuf:"varint,3,opt,name=clntReqYear" json:"clntReqYear,omitempty"`
	ClntReqMnth          *int32   `protobuf:"varint,4,opt,name=clntReqMnth" json:"clntReqMnth,omitempty"`
	ClntReqDay           *int32   `protobuf:"varint,5,opt,name=clntReqDay" json:"clntReqDay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientMsg) Reset()         { *m = ClientMsg{} }
func (m *ClientMsg) String() string { return proto.CompactTextString(m) }
func (*ClientMsg) ProtoMessage()    {}
func (*ClientMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_0108030a2e5322fc, []int{0}
}

func (m *ClientMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientMsg.Unmarshal(m, b)
}
func (m *ClientMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientMsg.Marshal(b, m, deterministic)
}
func (m *ClientMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientMsg.Merge(m, src)
}
func (m *ClientMsg) XXX_Size() int {
	return xxx_messageInfo_ClientMsg.Size(m)
}
func (m *ClientMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ClientMsg proto.InternalMessageInfo

func (m *ClientMsg) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *ClientMsg) GetActionCode() uint32 {
	if m != nil && m.ActionCode != nil {
		return *m.ActionCode
	}
	return 0
}

func (m *ClientMsg) GetClntReqYear() int32 {
	if m != nil && m.ClntReqYear != nil {
		return *m.ClntReqYear
	}
	return 0
}

func (m *ClientMsg) GetClntReqMnth() int32 {
	if m != nil && m.ClntReqMnth != nil {
		return *m.ClntReqMnth
	}
	return 0
}

func (m *ClientMsg) GetClntReqDay() int32 {
	if m != nil && m.ClntReqDay != nil {
		return *m.ClntReqDay
	}
	return 0
}

func init() {
	proto.RegisterType((*ClientMsg)(nil), "clientH2Req.clientMsg")
}

func init() { proto.RegisterFile("client_http2_req.proto", fileDescriptor_0108030a2e5322fc) }

var fileDescriptor_0108030a2e5322fc = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0x89, 0xcf, 0x28, 0x29, 0x29, 0x30, 0x8a, 0x2f, 0x4a, 0x2d, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x88, 0x7b, 0x18, 0x05, 0xa5, 0x16, 0x2a, 0xad, 0x66, 0xe4, 0xe2,
	0x84, 0xf0, 0x7d, 0x8b, 0xd3, 0x85, 0xa4, 0xb8, 0x38, 0x4a, 0x8b, 0x53, 0x8b, 0xfc, 0x12, 0x73,
	0x53, 0x25, 0x18, 0x15, 0x98, 0x34, 0x38, 0x83, 0xe0, 0x7c, 0x21, 0x39, 0x2e, 0xae, 0xc4, 0xe4,
	0x92, 0xcc, 0xfc, 0x3c, 0xe7, 0xfc, 0x94, 0x54, 0x09, 0x26, 0x05, 0x26, 0x0d, 0xde, 0x20, 0x24,
	0x11, 0x21, 0x05, 0x2e, 0xee, 0xe4, 0x9c, 0xbc, 0x92, 0xa0, 0xd4, 0xc2, 0xc8, 0xd4, 0xc4, 0x22,
	0x09, 0x66, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x64, 0x21, 0x24, 0x15, 0xbe, 0x79, 0x25, 0x19, 0x12,
	0x2c, 0x28, 0x2a, 0x40, 0x42, 0x20, 0x3b, 0xa0, 0x5c, 0x97, 0xc4, 0x4a, 0x09, 0x56, 0xb0, 0x02,
	0x24, 0x11, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0xf3, 0x3d, 0x07, 0xd3, 0x00, 0x00, 0x00,
}
