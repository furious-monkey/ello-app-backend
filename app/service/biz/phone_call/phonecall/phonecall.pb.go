// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: phonecall.proto

package phonecall

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	mtproto "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PhoneCallSession struct {
	Id                    int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AdminId               int64    `protobuf:"varint,2,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	AdminAccessHash       int64    `protobuf:"varint,3,opt,name=admin_access_hash,json=adminAccessHash,proto3" json:"admin_access_hash,omitempty"`
	ParticipantId         int64    `protobuf:"varint,4,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	ParticipantAccessHash int64    `protobuf:"varint,5,opt,name=participant_access_hash,json=participantAccessHash,proto3" json:"participant_access_hash,omitempty"`
	UdpP2P                bool     `protobuf:"varint,6,opt,name=udp_p2p,json=udpP2p,proto3" json:"udp_p2p,omitempty"`
	UdpReflector          bool     `protobuf:"varint,7,opt,name=udp_reflector,json=udpReflector,proto3" json:"udp_reflector,omitempty"`
	MinLayer              int32    `protobuf:"varint,8,opt,name=min_layer,json=minLayer,proto3" json:"min_layer,omitempty"`
	MaxLayer              int32    `protobuf:"varint,9,opt,name=max_layer,json=maxLayer,proto3" json:"max_layer,omitempty"`
	GA                    []byte   `protobuf:"bytes,10,opt,name=g_a,json=gA,proto3" json:"g_a,omitempty"`
	GB                    []byte   `protobuf:"bytes,11,opt,name=g_b,json=gB,proto3" json:"g_b,omitempty"`
	State                 int32    `protobuf:"varint,12,opt,name=state,proto3" json:"state,omitempty"`
	Date                  int64    `protobuf:"varint,13,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *PhoneCallSession) Reset()         { *m = PhoneCallSession{} }
func (m *PhoneCallSession) String() string { return proto.CompactTextString(m) }
func (*PhoneCallSession) ProtoMessage()    {}
func (*PhoneCallSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e245dd4547dc78, []int{0}
}
func (m *PhoneCallSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhoneCallSession.Unmarshal(m, b)
}
func (m *PhoneCallSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhoneCallSession.Marshal(b, m, deterministic)
}
func (m *PhoneCallSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhoneCallSession.Merge(m, src)
}
func (m *PhoneCallSession) XXX_Size() int {
	return xxx_messageInfo_PhoneCallSession.Size(m)
}
func (m *PhoneCallSession) XXX_DiscardUnknown() {
	xxx_messageInfo_PhoneCallSession.DiscardUnknown(m)
}

var xxx_messageInfo_PhoneCallSession proto.InternalMessageInfo

func (m *PhoneCallSession) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PhoneCallSession) GetAdminId() int64 {
	if m != nil {
		return m.AdminId
	}
	return 0
}

func (m *PhoneCallSession) GetAdminAccessHash() int64 {
	if m != nil {
		return m.AdminAccessHash
	}
	return 0
}

func (m *PhoneCallSession) GetParticipantId() int64 {
	if m != nil {
		return m.ParticipantId
	}
	return 0
}

func (m *PhoneCallSession) GetParticipantAccessHash() int64 {
	if m != nil {
		return m.ParticipantAccessHash
	}
	return 0
}

func (m *PhoneCallSession) GetUdpP2P() bool {
	if m != nil {
		return m.UdpP2P
	}
	return false
}

func (m *PhoneCallSession) GetUdpReflector() bool {
	if m != nil {
		return m.UdpReflector
	}
	return false
}

func (m *PhoneCallSession) GetMinLayer() int32 {
	if m != nil {
		return m.MinLayer
	}
	return 0
}

func (m *PhoneCallSession) GetMaxLayer() int32 {
	if m != nil {
		return m.MaxLayer
	}
	return 0
}

func (m *PhoneCallSession) GetGA() []byte {
	if m != nil {
		return m.GA
	}
	return nil
}

func (m *PhoneCallSession) GetGB() []byte {
	if m != nil {
		return m.GB
	}
	return nil
}

func (m *PhoneCallSession) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *PhoneCallSession) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

type TLMakePhoneCallSession struct {
	AdminId              int64                      `protobuf:"varint,1,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	ParticipantId        int64                      `protobuf:"varint,2,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	GaHash               []byte                     `protobuf:"bytes,3,opt,name=ga_hash,json=gaHash,proto3" json:"ga_hash,omitempty"`
	Protocol             *mtproto.PhoneCallProtocol `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	RandomId             int32                      `protobuf:"varint,5,opt,name=random_id,json=randomId,proto3" json:"random_id,omitempty"`
	IsVideo              bool                       `protobuf:"varint,6,opt,name=is_video,json=isVideo,proto3" json:"is_video,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TLMakePhoneCallSession) Reset()         { *m = TLMakePhoneCallSession{} }
func (m *TLMakePhoneCallSession) String() string { return proto.CompactTextString(m) }
func (*TLMakePhoneCallSession) ProtoMessage()    {}
func (*TLMakePhoneCallSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e245dd4547dc78, []int{1}
}
func (m *TLMakePhoneCallSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLMakePhoneCallSession.Unmarshal(m, b)
}
func (m *TLMakePhoneCallSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLMakePhoneCallSession.Marshal(b, m, deterministic)
}
func (m *TLMakePhoneCallSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLMakePhoneCallSession.Merge(m, src)
}
func (m *TLMakePhoneCallSession) XXX_Size() int {
	return xxx_messageInfo_TLMakePhoneCallSession.Size(m)
}
func (m *TLMakePhoneCallSession) XXX_DiscardUnknown() {
	xxx_messageInfo_TLMakePhoneCallSession.DiscardUnknown(m)
}

var xxx_messageInfo_TLMakePhoneCallSession proto.InternalMessageInfo

func (m *TLMakePhoneCallSession) GetAdminId() int64 {
	if m != nil {
		return m.AdminId
	}
	return 0
}

func (m *TLMakePhoneCallSession) GetParticipantId() int64 {
	if m != nil {
		return m.ParticipantId
	}
	return 0
}

func (m *TLMakePhoneCallSession) GetGaHash() []byte {
	if m != nil {
		return m.GaHash
	}
	return nil
}

func (m *TLMakePhoneCallSession) GetProtocol() *mtproto.PhoneCallProtocol {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func (m *TLMakePhoneCallSession) GetRandomId() int32 {
	if m != nil {
		return m.RandomId
	}
	return 0
}

func (m *TLMakePhoneCallSession) GetIsVideo() bool {
	if m != nil {
		return m.IsVideo
	}
	return false
}

type TLMakePhoneCallSessionByLoad struct {
	SessionId            int64    `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLMakePhoneCallSessionByLoad) Reset()         { *m = TLMakePhoneCallSessionByLoad{} }
func (m *TLMakePhoneCallSessionByLoad) String() string { return proto.CompactTextString(m) }
func (*TLMakePhoneCallSessionByLoad) ProtoMessage()    {}
func (*TLMakePhoneCallSessionByLoad) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e245dd4547dc78, []int{2}
}
func (m *TLMakePhoneCallSessionByLoad) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLMakePhoneCallSessionByLoad.Unmarshal(m, b)
}
func (m *TLMakePhoneCallSessionByLoad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLMakePhoneCallSessionByLoad.Marshal(b, m, deterministic)
}
func (m *TLMakePhoneCallSessionByLoad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLMakePhoneCallSessionByLoad.Merge(m, src)
}
func (m *TLMakePhoneCallSessionByLoad) XXX_Size() int {
	return xxx_messageInfo_TLMakePhoneCallSessionByLoad.Size(m)
}
func (m *TLMakePhoneCallSessionByLoad) XXX_DiscardUnknown() {
	xxx_messageInfo_TLMakePhoneCallSessionByLoad.DiscardUnknown(m)
}

var xxx_messageInfo_TLMakePhoneCallSessionByLoad proto.InternalMessageInfo

func (m *TLMakePhoneCallSessionByLoad) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

type TLSetGB struct {
	SessionId            int64    `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Gb                   []byte   `protobuf:"bytes,2,opt,name=gb,proto3" json:"gb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLSetGB) Reset()         { *m = TLSetGB{} }
func (m *TLSetGB) String() string { return proto.CompactTextString(m) }
func (*TLSetGB) ProtoMessage()    {}
func (*TLSetGB) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e245dd4547dc78, []int{3}
}
func (m *TLSetGB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLSetGB.Unmarshal(m, b)
}
func (m *TLSetGB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLSetGB.Marshal(b, m, deterministic)
}
func (m *TLSetGB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSetGB.Merge(m, src)
}
func (m *TLSetGB) XXX_Size() int {
	return xxx_messageInfo_TLSetGB.Size(m)
}
func (m *TLSetGB) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSetGB.DiscardUnknown(m)
}

var xxx_messageInfo_TLSetGB proto.InternalMessageInfo

func (m *TLSetGB) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *TLSetGB) GetGb() []byte {
	if m != nil {
		return m.Gb
	}
	return nil
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e245dd4547dc78, []int{4}
}
func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PhoneCallSession)(nil), "phonecall.PhoneCallSession")
	proto.RegisterType((*TLMakePhoneCallSession)(nil), "phonecall.TLMakePhoneCallSession")
	proto.RegisterType((*TLMakePhoneCallSessionByLoad)(nil), "phonecall.TLMakePhoneCallSessionByLoad")
	proto.RegisterType((*TLSetGB)(nil), "phonecall.TLSetGB")
	proto.RegisterType((*Void)(nil), "phonecall.Void")
}

func init() { proto.RegisterFile("phonecall.proto", fileDescriptor_68e245dd4547dc78) }

var fileDescriptor_68e245dd4547dc78 = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x4f, 0xdb, 0x30,
	0x14, 0x55, 0x5a, 0xfa, 0x75, 0x69, 0x29, 0xb3, 0x60, 0x84, 0xc0, 0xa4, 0xd2, 0x69, 0x5a, 0x35,
	0x4d, 0x7d, 0x28, 0x12, 0xda, 0xcb, 0x1e, 0x80, 0x87, 0xad, 0x52, 0x27, 0x55, 0x01, 0xf1, 0x1a,
	0xb9, 0xb6, 0xd7, 0x58, 0xa4, 0xb1, 0x65, 0x1b, 0x44, 0xff, 0xc3, 0xfe, 0xe5, 0xfe, 0xc4, 0x1e,
	0x27, 0x3b, 0x21, 0x4d, 0x69, 0x05, 0x6f, 0xbd, 0xe7, 0x9c, 0x7b, 0x6f, 0xef, 0x39, 0x56, 0xa0,
	0x2b, 0x63, 0x91, 0x32, 0x82, 0x93, 0x64, 0x28, 0x95, 0x30, 0x02, 0xb5, 0x0a, 0x20, 0x38, 0x54,
	0x92, 0x44, 0x4c, 0x29, 0xa1, 0x22, 0x22, 0x28, 0xd3, 0x99, 0x22, 0x08, 0x34, 0x89, 0xd9, 0x02,
	0x0f, 0x4d, 0x32, 0x24, 0x42, 0xb1, 0xc8, 0x2c, 0x65, 0xc1, 0x1d, 0x96, 0x38, 0x45, 0xce, 0x47,
	0x39, 0x7c, 0xbc, 0x82, 0x63, 0x9c, 0x52, 0x1d, 0xe3, 0x7b, 0x96, 0x53, 0x67, 0x5b, 0xa8, 0x48,
	0x33, 0xf5, 0xc8, 0xc9, 0xb3, 0xe4, 0x60, 0x25, 0xd1, 0xcb, 0x94, 0xe4, 0xe8, 0xe9, 0x3a, 0xfa,
	0xa2, 0xa7, 0xb4, 0xd1, 0x28, 0x9c, 0x6a, 0x29, 0x94, 0xd9, 0xdc, 0x58, 0x50, 0x2f, 0xba, 0xf7,
	0x9f, 0x4b, 0x93, 0xdb, 0x12, 0x74, 0x16, 0x4b, 0x77, 0x68, 0x56, 0xf6, 0xff, 0x54, 0x61, 0x7f,
	0x6a, 0x8d, 0xba, 0xc6, 0x49, 0x72, 0xc3, 0xb4, 0xe6, 0x22, 0x45, 0x7b, 0x50, 0xe1, 0xd4, 0xf7,
	0x7a, 0xde, 0xa0, 0x1a, 0x56, 0x38, 0x45, 0xc7, 0xd0, 0xc4, 0x74, 0xc1, 0xd3, 0x88, 0x53, 0xbf,
	0xe2, 0xd0, 0x86, 0xab, 0xc7, 0x14, 0x7d, 0x81, 0x77, 0x19, 0x85, 0x09, 0x61, 0x5a, 0x47, 0x31,
	0xd6, 0xb1, 0x5f, 0x75, 0x9a, 0xae, 0x23, 0x2e, 0x1d, 0xfe, 0x13, 0xeb, 0x18, 0x7d, 0x82, 0x3d,
	0x89, 0x95, 0xe1, 0x84, 0x4b, 0x9c, 0x1a, 0x3b, 0x6c, 0xc7, 0x09, 0x3b, 0x25, 0x74, 0x4c, 0xd1,
	0x05, 0x1c, 0x95, 0x65, 0xe5, 0xc1, 0x35, 0xa7, 0x3f, 0x2c, 0xd1, 0xa5, 0xf1, 0x47, 0xd0, 0x78,
	0xa0, 0x32, 0x92, 0x23, 0xe9, 0xd7, 0x7b, 0xde, 0xa0, 0x19, 0xd6, 0x1f, 0xa8, 0x9c, 0x8e, 0x24,
	0xfa, 0x08, 0x1d, 0x4b, 0x28, 0xf6, 0x3b, 0x61, 0xc4, 0x08, 0xe5, 0x37, 0x1c, 0xdd, 0x7e, 0xa0,
	0x32, 0x7c, 0xc6, 0xd0, 0x09, 0xb4, 0xec, 0x19, 0x09, 0x5e, 0x32, 0xe5, 0x37, 0x7b, 0xde, 0xa0,
	0x16, 0x36, 0x17, 0x3c, 0x9d, 0xd8, 0xda, 0x91, 0xf8, 0x29, 0x27, 0x5b, 0x39, 0x89, 0x9f, 0x32,
	0xb2, 0x0b, 0xd5, 0x79, 0x84, 0x7d, 0xe8, 0x79, 0x83, 0x76, 0x58, 0x99, 0x5f, 0x66, 0xc0, 0xcc,
	0xdf, 0xcd, 0x81, 0x2b, 0x74, 0x00, 0x35, 0x6d, 0xb0, 0x61, 0x7e, 0xdb, 0xb5, 0x66, 0x05, 0x42,
	0xb0, 0x43, 0x2d, 0xd8, 0x71, 0x47, 0xb9, 0xdf, 0xfd, 0xbf, 0x1e, 0xbc, 0xbf, 0x9d, 0xfc, 0xc2,
	0xf7, 0x6c, 0x23, 0x94, 0x72, 0x08, 0xde, 0x7a, 0x08, 0x9b, 0xc6, 0x56, 0xb6, 0x19, 0x7b, 0x04,
	0x8d, 0x39, 0x5e, 0x25, 0xd4, 0x0e, 0xeb, 0x73, 0xec, 0x9c, 0xbb, 0x80, 0xa6, 0x7b, 0x0d, 0x44,
	0x24, 0x2e, 0x92, 0xdd, 0x51, 0x30, 0x5c, 0x18, 0x07, 0x0d, 0x8b, 0xff, 0x31, 0xcd, 0x15, 0x61,
	0xa1, 0xb5, 0xb6, 0x28, 0x9c, 0x52, 0xb1, 0xb0, 0x2b, 0x6b, 0x99, 0x2d, 0x19, 0x30, 0x76, 0x8f,
	0x86, 0xeb, 0xe8, 0x91, 0x53, 0x26, 0xf2, 0x3c, 0x1a, 0x5c, 0xdf, 0xd9, 0xb2, 0xff, 0x1d, 0x4e,
	0xb7, 0x1f, 0x79, 0xb5, 0x9c, 0x08, 0x4c, 0xd1, 0x07, 0x00, 0x9d, 0x01, 0xab, 0x63, 0x5b, 0x39,
	0x32, 0xa6, 0xfd, 0x6f, 0xd0, 0xb8, 0x9d, 0xdc, 0x30, 0xf3, 0xe3, 0xea, 0x0d, 0xa5, 0x7d, 0xc8,
	0xf3, 0x99, 0x33, 0xc3, 0x06, 0x31, 0xeb, 0xd7, 0x61, 0xe7, 0x4e, 0x70, 0x3a, 0xfa, 0xe7, 0x41,
	0x3b, 0x9c, 0x5e, 0x17, 0xeb, 0xd1, 0x57, 0xa8, 0x65, 0x03, 0xd1, 0x70, 0xf5, 0x1d, 0xc9, 0x97,
	0x04, 0xdd, 0x12, 0x66, 0xdb, 0xd1, 0x1d, 0x1c, 0x6c, 0x8d, 0xe8, 0x6c, 0xad, 0x79, 0x9b, 0x24,
	0x38, 0x29, 0x49, 0x36, 0xfa, 0x67, 0x10, 0xbc, 0xe2, 0xca, 0xe7, 0x37, 0xa7, 0x67, 0xc2, 0x57,
	0x77, 0xcc, 0xea, 0x2e, 0xbd, 0xf3, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xbe, 0xf1, 0x98,
	0x30, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCPhoneCallClient is the client API for RPCPhoneCall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCPhoneCallClient interface {
	SetGB(ctx context.Context, in *TLSetGB, opts ...grpc.CallOption) (*Void, error)
	MakePhoneCallSession(ctx context.Context, in *TLMakePhoneCallSession, opts ...grpc.CallOption) (*PhoneCallSession, error)
	MakePhoneCallSessionByLoad(ctx context.Context, in *TLMakePhoneCallSessionByLoad, opts ...grpc.CallOption) (*PhoneCallSession, error)
}

type rPCPhoneCallClient struct {
	cc *grpc.ClientConn
}

func NewRPCPhoneCallClient(cc *grpc.ClientConn) RPCPhoneCallClient {
	return &rPCPhoneCallClient{cc}
}

func (c *rPCPhoneCallClient) SetGB(ctx context.Context, in *TLSetGB, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/phonecall.RPCPhoneCall/SetGB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCPhoneCallClient) MakePhoneCallSession(ctx context.Context, in *TLMakePhoneCallSession, opts ...grpc.CallOption) (*PhoneCallSession, error) {
	out := new(PhoneCallSession)
	err := c.cc.Invoke(ctx, "/phonecall.RPCPhoneCall/MakePhoneCallSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCPhoneCallClient) MakePhoneCallSessionByLoad(ctx context.Context, in *TLMakePhoneCallSessionByLoad, opts ...grpc.CallOption) (*PhoneCallSession, error) {
	out := new(PhoneCallSession)
	err := c.cc.Invoke(ctx, "/phonecall.RPCPhoneCall/MakePhoneCallSessionByLoad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCPhoneCallServer is the server API for RPCPhoneCall service.
type RPCPhoneCallServer interface {
	SetGB(context.Context, *TLSetGB) (*Void, error)
	MakePhoneCallSession(context.Context, *TLMakePhoneCallSession) (*PhoneCallSession, error)
	MakePhoneCallSessionByLoad(context.Context, *TLMakePhoneCallSessionByLoad) (*PhoneCallSession, error)
}

// UnimplementedRPCPhoneCallServer can be embedded to have forward compatible implementations.
type UnimplementedRPCPhoneCallServer struct {
}

func (*UnimplementedRPCPhoneCallServer) SetGB(ctx context.Context, req *TLSetGB) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGB not implemented")
}
func (*UnimplementedRPCPhoneCallServer) MakePhoneCallSession(ctx context.Context, req *TLMakePhoneCallSession) (*PhoneCallSession, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakePhoneCallSession not implemented")
}
func (*UnimplementedRPCPhoneCallServer) MakePhoneCallSessionByLoad(ctx context.Context, req *TLMakePhoneCallSessionByLoad) (*PhoneCallSession, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakePhoneCallSessionByLoad not implemented")
}

func RegisterRPCPhoneCallServer(s *grpc.Server, srv RPCPhoneCallServer) {
	s.RegisterService(&_RPCPhoneCall_serviceDesc, srv)
}

func _RPCPhoneCall_SetGB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSetGB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCPhoneCallServer).SetGB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phonecall.RPCPhoneCall/SetGB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCPhoneCallServer).SetGB(ctx, req.(*TLSetGB))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCPhoneCall_MakePhoneCallSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLMakePhoneCallSession)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCPhoneCallServer).MakePhoneCallSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phonecall.RPCPhoneCall/MakePhoneCallSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCPhoneCallServer).MakePhoneCallSession(ctx, req.(*TLMakePhoneCallSession))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCPhoneCall_MakePhoneCallSessionByLoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLMakePhoneCallSessionByLoad)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCPhoneCallServer).MakePhoneCallSessionByLoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phonecall.RPCPhoneCall/MakePhoneCallSessionByLoad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCPhoneCallServer).MakePhoneCallSessionByLoad(ctx, req.(*TLMakePhoneCallSessionByLoad))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCPhoneCall_serviceDesc = grpc.ServiceDesc{
	ServiceName: "phonecall.RPCPhoneCall",
	HandlerType: (*RPCPhoneCallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetGB",
			Handler:    _RPCPhoneCall_SetGB_Handler,
		},
		{
			MethodName: "MakePhoneCallSession",
			Handler:    _RPCPhoneCall_MakePhoneCallSession_Handler,
		},
		{
			MethodName: "MakePhoneCallSessionByLoad",
			Handler:    _RPCPhoneCall_MakePhoneCallSessionByLoad_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "phonecall.proto",
}
