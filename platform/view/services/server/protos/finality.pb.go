// Code generated by protoc-gen-go. DO NOT EDIT.
// source: finality.proto

package protos

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

type IsTxFinal struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Txid                 string   `protobuf:"bytes,2,opt,name=txid,proto3" json:"txid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsTxFinal) Reset()         { *m = IsTxFinal{} }
func (m *IsTxFinal) String() string { return proto.CompactTextString(m) }
func (*IsTxFinal) ProtoMessage()    {}
func (*IsTxFinal) Descriptor() ([]byte, []int) {
	return fileDescriptor_0144d353a635b215, []int{0}
}

func (m *IsTxFinal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsTxFinal.Unmarshal(m, b)
}
func (m *IsTxFinal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsTxFinal.Marshal(b, m, deterministic)
}
func (m *IsTxFinal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsTxFinal.Merge(m, src)
}
func (m *IsTxFinal) XXX_Size() int {
	return xxx_messageInfo_IsTxFinal.Size(m)
}
func (m *IsTxFinal) XXX_DiscardUnknown() {
	xxx_messageInfo_IsTxFinal.DiscardUnknown(m)
}

var xxx_messageInfo_IsTxFinal proto.InternalMessageInfo

func (m *IsTxFinal) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *IsTxFinal) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

type IsTxFinalResponse struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsTxFinalResponse) Reset()         { *m = IsTxFinalResponse{} }
func (m *IsTxFinalResponse) String() string { return proto.CompactTextString(m) }
func (*IsTxFinalResponse) ProtoMessage()    {}
func (*IsTxFinalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0144d353a635b215, []int{1}
}

func (m *IsTxFinalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsTxFinalResponse.Unmarshal(m, b)
}
func (m *IsTxFinalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsTxFinalResponse.Marshal(b, m, deterministic)
}
func (m *IsTxFinalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsTxFinalResponse.Merge(m, src)
}
func (m *IsTxFinalResponse) XXX_Size() int {
	return xxx_messageInfo_IsTxFinalResponse.Size(m)
}
func (m *IsTxFinalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsTxFinalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsTxFinalResponse proto.InternalMessageInfo

func (m *IsTxFinalResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*IsTxFinal)(nil), "protos.IsTxFinal")
	proto.RegisterType((*IsTxFinalResponse)(nil), "protos.IsTxFinalResponse")
}

func init() { proto.RegisterFile("finality.proto", fileDescriptor_0144d353a635b215) }

var fileDescriptor_0144d353a635b215 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcb, 0xcc, 0x4b,
	0xcc, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a,
	0x96, 0x5c, 0x9c, 0x9e, 0xc5, 0x21, 0x15, 0x6e, 0x20, 0x59, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0x8c,
	0xc4, 0xbc, 0xbc, 0xd4, 0x1c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x48, 0x88,
	0x8b, 0xa5, 0xa4, 0x22, 0x33, 0x45, 0x82, 0x09, 0x2c, 0x0c, 0x66, 0x2b, 0xe9, 0x72, 0x09, 0xc2,
	0xb5, 0x06, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x82, 0x8c, 0x28, 0x48, 0xac, 0xcc, 0xc9,
	0x4f, 0x4c, 0x01, 0x1b, 0xc1, 0x13, 0x04, 0xe3, 0x3a, 0x71, 0x47, 0x41, 0xed, 0x6c, 0x60, 0x64,
	0x4c, 0x82, 0x30, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xc1, 0x00, 0xe3, 0x97, 0x00,
	0x00, 0x00,
}
