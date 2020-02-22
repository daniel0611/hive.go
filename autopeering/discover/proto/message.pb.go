// Code generated by protoc-gen-go. DO NOT EDIT.
// source: discover/proto/message.proto

package proto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	proto2 "github.com/iotaledger/hive.go/autopeering/peer/proto"
	proto1 "github.com/iotaledger/hive.go/autopeering/peer/service/proto"
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

type Ping struct {
	// protocol version number
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// string form of the return address (e.g. "192.0.2.1:25", "[2001:db8::1]:80")
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	// string form of the recipient address
	To string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	// unix time
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_43f14146485f66eb, []int{0}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ping) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Ping) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Ping) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Pong struct {
	// hash of the ping packet
	PingHash []byte `protobuf:"bytes,1,opt,name=ping_hash,json=pingHash,proto3" json:"ping_hash,omitempty"`
	// string form of the recipient address
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// services supported by the sender
	Services             *proto1.ServiceMap `protobuf:"bytes,3,opt,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_43f14146485f66eb, []int{1}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetPingHash() []byte {
	if m != nil {
		return m.PingHash
	}
	return nil
}

func (m *Pong) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Pong) GetServices() *proto1.ServiceMap {
	if m != nil {
		return m.Services
	}
	return nil
}

type DiscoveryRequest struct {
	// string form of the recipient address
	To string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	// unix time
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryRequest) Reset()         { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()    {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43f14146485f66eb, []int{2}
}

func (m *DiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryRequest.Unmarshal(m, b)
}
func (m *DiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryRequest.Marshal(b, m, deterministic)
}
func (m *DiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryRequest.Merge(m, src)
}
func (m *DiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoveryRequest.Size(m)
}
func (m *DiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryRequest proto.InternalMessageInfo

func (m *DiscoveryRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *DiscoveryRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type DiscoveryResponse struct {
	// hash of the corresponding request
	ReqHash []byte `protobuf:"bytes,1,opt,name=req_hash,json=reqHash,proto3" json:"req_hash,omitempty"`
	// list of peers
	Peers                []*proto2.Peer `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DiscoveryResponse) Reset()         { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()    {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_43f14146485f66eb, []int{3}
}

func (m *DiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryResponse.Unmarshal(m, b)
}
func (m *DiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *DiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryResponse.Merge(m, src)
}
func (m *DiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DiscoveryResponse.Size(m)
}
func (m *DiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryResponse proto.InternalMessageInfo

func (m *DiscoveryResponse) GetReqHash() []byte {
	if m != nil {
		return m.ReqHash
	}
	return nil
}

func (m *DiscoveryResponse) GetPeers() []*proto2.Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func init() {
	proto.RegisterType((*Ping)(nil), "proto.Ping")
	proto.RegisterType((*Pong)(nil), "proto.Pong")
	proto.RegisterType((*DiscoveryRequest)(nil), "proto.DiscoveryRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "proto.DiscoveryResponse")
}

func init() { proto.RegisterFile("discover/proto/message.proto", fileDescriptor_43f14146485f66eb) }

var fileDescriptor_43f14146485f66eb = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x59, 0xd7, 0xb9, 0xf6, 0x4d, 0xc5, 0x05, 0x84, 0x3a, 0x77, 0x98, 0x3b, 0x79, 0x59,
	0x0b, 0x2a, 0x9e, 0x45, 0x3c, 0x78, 0x11, 0x66, 0xbc, 0x79, 0x91, 0xb4, 0x8b, 0x6d, 0xc0, 0x36,
	0x59, 0x93, 0x0e, 0xfc, 0xef, 0x7d, 0x4d, 0x63, 0x9d, 0xe2, 0x29, 0xef, 0x7d, 0xef, 0xf1, 0x7d,
	0xbf, 0x24, 0x30, 0xdf, 0x08, 0x9d, 0xc9, 0x1d, 0xaf, 0x13, 0x55, 0x4b, 0x23, 0x93, 0x92, 0x6b,
	0xcd, 0x72, 0x1e, 0xdb, 0x8e, 0x8c, 0xec, 0x31, 0x3b, 0x55, 0xbc, 0x5f, 0x68, 0xcb, 0x6e, 0x3a,
	0x5b, 0x58, 0x59, 0xf3, 0x7a, 0x27, 0x32, 0xee, 0xc6, 0xae, 0xeb, 0x36, 0x96, 0x29, 0xf8, 0x6b,
	0x51, 0xe5, 0x24, 0x82, 0x31, 0x46, 0x68, 0x21, 0xab, 0x68, 0xb0, 0x18, 0x5c, 0x1e, 0xd1, 0xef,
	0x96, 0x10, 0xf0, 0xdf, 0x6b, 0x59, 0x46, 0x1e, 0xca, 0x21, 0xb5, 0x35, 0x39, 0x06, 0xcf, 0xc8,
	0x68, 0x68, 0x15, 0xac, 0xc8, 0x1c, 0x42, 0x23, 0x10, 0xcc, 0xb0, 0x52, 0x45, 0x3e, 0xca, 0x43,
	0xfa, 0x23, 0xd8, 0x0c, 0x89, 0x19, 0xe7, 0x10, 0x2a, 0xcc, 0x7a, 0x2b, 0x98, 0x2e, 0x6c, 0xca,
	0x21, 0x0d, 0x5a, 0xe1, 0x11, 0x7b, 0x67, 0xe9, 0xf5, 0x96, 0x2b, 0x08, 0x1c, 0xa9, 0xb6, 0x41,
	0x93, 0xab, 0x69, 0x87, 0x1c, 0xbf, 0x74, 0xf2, 0x13, 0x53, 0xb4, 0x5f, 0x59, 0xde, 0xc1, 0xc9,
	0x83, 0x7b, 0xa7, 0x4f, 0xca, 0xb7, 0x0d, 0x46, 0x3b, 0xcb, 0xc1, 0xff, 0x94, 0xde, 0x5f, 0xca,
	0x67, 0x98, 0xee, 0x39, 0x68, 0x25, 0x2b, 0xcd, 0xc9, 0x19, 0x04, 0x35, 0xdf, 0xee, 0x13, 0x8f,
	0xb1, 0xb7, 0xc0, 0x17, 0x30, 0x6a, 0x5f, 0x57, 0xa3, 0xd3, 0x10, 0xe9, 0x26, 0x8e, 0x6e, 0x8d,
	0x1a, 0xed, 0x26, 0xf7, 0xb7, 0xaf, 0x37, 0xb9, 0x30, 0x45, 0x93, 0xc6, 0x99, 0x2c, 0x13, 0x21,
	0x0d, 0xfb, 0xe0, 0x9b, 0x1c, 0x7f, 0x84, 0x35, 0x46, 0xb6, 0x2b, 0x78, 0xf9, 0x95, 0x16, 0x65,
	0xf2, 0xfb, 0x8b, 0xd3, 0x03, 0x7b, 0x5c, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xa6, 0x1f,
	0xce, 0xfb, 0x01, 0x00, 0x00,
}