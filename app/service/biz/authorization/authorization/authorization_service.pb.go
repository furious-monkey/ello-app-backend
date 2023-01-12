// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: authorization_service.proto

package authorization

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

type MData struct {
	ServerId             string   `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ClientAddr           string   `protobuf:"bytes,2,opt,name=client_addr,json=clientAddr,proto3" json:"client_addr,omitempty"`
	AuthId               int64    `protobuf:"varint,3,opt,name=auth_id,json=authId,proto3" json:"auth_id,omitempty"`
	SessionId            int64    `protobuf:"varint,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	ReceiveTime          int64    `protobuf:"varint,5,opt,name=receive_time,json=receiveTime,proto3" json:"receive_time,omitempty"`
	UserId               int64    `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClientMsgId          int64    `protobuf:"varint,7,opt,name=client_msg_id,json=clientMsgId,proto3" json:"client_msg_id,omitempty"`
	IsBot                bool     `protobuf:"varint,8,opt,name=is_bot,json=isBot,proto3" json:"is_bot,omitempty"`
	Layer                int32    `protobuf:"varint,9,opt,name=layer,proto3" json:"layer,omitempty"`
	Client               string   `protobuf:"bytes,10,opt,name=client,proto3" json:"client,omitempty"`
	IsAdmin              bool     `protobuf:"varint,11,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	Langpack             string   `protobuf:"bytes,13,opt,name=langpack,proto3" json:"langpack,omitempty"`
	PermAuthKeyId        int64    `protobuf:"varint,14,opt,name=perm_auth_key_id,json=permAuthKeyId,proto3" json:"perm_auth_key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MData) Reset()         { *m = MData{} }
func (m *MData) String() string { return proto.CompactTextString(m) }
func (*MData) ProtoMessage()    {}
func (*MData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e27dc8f5aa0b42d7, []int{0}
}
func (m *MData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MData.Unmarshal(m, b)
}
func (m *MData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MData.Marshal(b, m, deterministic)
}
func (m *MData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MData.Merge(m, src)
}
func (m *MData) XXX_Size() int {
	return xxx_messageInfo_MData.Size(m)
}
func (m *MData) XXX_DiscardUnknown() {
	xxx_messageInfo_MData.DiscardUnknown(m)
}

var xxx_messageInfo_MData proto.InternalMessageInfo

func (m *MData) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *MData) GetClientAddr() string {
	if m != nil {
		return m.ClientAddr
	}
	return ""
}

func (m *MData) GetAuthId() int64 {
	if m != nil {
		return m.AuthId
	}
	return 0
}

func (m *MData) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *MData) GetReceiveTime() int64 {
	if m != nil {
		return m.ReceiveTime
	}
	return 0
}

func (m *MData) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *MData) GetClientMsgId() int64 {
	if m != nil {
		return m.ClientMsgId
	}
	return 0
}

func (m *MData) GetIsBot() bool {
	if m != nil {
		return m.IsBot
	}
	return false
}

func (m *MData) GetLayer() int32 {
	if m != nil {
		return m.Layer
	}
	return 0
}

func (m *MData) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *MData) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *MData) GetLangpack() string {
	if m != nil {
		return m.Langpack
	}
	return ""
}

func (m *MData) GetPermAuthKeyId() int64 {
	if m != nil {
		return m.PermAuthKeyId
	}
	return 0
}

type AuthSignUpRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Gender               string   `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	DateOfBirth          string   `protobuf:"bytes,4,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	CountryCode          string   `protobuf:"bytes,7,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Avatar               string   `protobuf:"bytes,8,opt,name=avatar,proto3" json:"avatar,omitempty"`
	FirstName            string   `protobuf:"bytes,9,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,10,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	MData                *MData   `protobuf:"bytes,11,opt,name=m_data,json=mData,proto3" json:"m_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthSignUpRequest) Reset()         { *m = AuthSignUpRequest{} }
func (m *AuthSignUpRequest) String() string { return proto.CompactTextString(m) }
func (*AuthSignUpRequest) ProtoMessage()    {}
func (*AuthSignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e27dc8f5aa0b42d7, []int{1}
}
func (m *AuthSignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthSignUpRequest.Unmarshal(m, b)
}
func (m *AuthSignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthSignUpRequest.Marshal(b, m, deterministic)
}
func (m *AuthSignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthSignUpRequest.Merge(m, src)
}
func (m *AuthSignUpRequest) XXX_Size() int {
	return xxx_messageInfo_AuthSignUpRequest.Size(m)
}
func (m *AuthSignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthSignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthSignUpRequest proto.InternalMessageInfo

func (m *AuthSignUpRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthSignUpRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthSignUpRequest) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *AuthSignUpRequest) GetDateOfBirth() string {
	if m != nil {
		return m.DateOfBirth
	}
	return ""
}

func (m *AuthSignUpRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthSignUpRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *AuthSignUpRequest) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *AuthSignUpRequest) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *AuthSignUpRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *AuthSignUpRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *AuthSignUpRequest) GetMData() *MData {
	if m != nil {
		return m.MData
	}
	return nil
}

type AuthSignInRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	MData                *MData   `protobuf:"bytes,3,opt,name=m_data,json=mData,proto3" json:"m_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthSignInRequest) Reset()         { *m = AuthSignInRequest{} }
func (m *AuthSignInRequest) String() string { return proto.CompactTextString(m) }
func (*AuthSignInRequest) ProtoMessage()    {}
func (*AuthSignInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e27dc8f5aa0b42d7, []int{2}
}
func (m *AuthSignInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthSignInRequest.Unmarshal(m, b)
}
func (m *AuthSignInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthSignInRequest.Marshal(b, m, deterministic)
}
func (m *AuthSignInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthSignInRequest.Merge(m, src)
}
func (m *AuthSignInRequest) XXX_Size() int {
	return xxx_messageInfo_AuthSignInRequest.Size(m)
}
func (m *AuthSignInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthSignInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthSignInRequest proto.InternalMessageInfo

func (m *AuthSignInRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthSignInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthSignInRequest) GetMData() *MData {
	if m != nil {
		return m.MData
	}
	return nil
}

type AuthSignUpRsp struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	ConfirmationExpire   int64    `protobuf:"varint,2,opt,name=confirmation_expire,json=confirmationExpire,proto3" json:"confirmation_expire,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthSignUpRsp) Reset()         { *m = AuthSignUpRsp{} }
func (m *AuthSignUpRsp) String() string { return proto.CompactTextString(m) }
func (*AuthSignUpRsp) ProtoMessage()    {}
func (*AuthSignUpRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e27dc8f5aa0b42d7, []int{3}
}
func (m *AuthSignUpRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthSignUpRsp.Unmarshal(m, b)
}
func (m *AuthSignUpRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthSignUpRsp.Marshal(b, m, deterministic)
}
func (m *AuthSignUpRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthSignUpRsp.Merge(m, src)
}
func (m *AuthSignUpRsp) XXX_Size() int {
	return xxx_messageInfo_AuthSignUpRsp.Size(m)
}
func (m *AuthSignUpRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthSignUpRsp.DiscardUnknown(m)
}

var xxx_messageInfo_AuthSignUpRsp proto.InternalMessageInfo

func (m *AuthSignUpRsp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthSignUpRsp) GetConfirmationExpire() int64 {
	if m != nil {
		return m.ConfirmationExpire
	}
	return 0
}

type ConfirmationRequest struct {
	UsernameOrEmail      string   `protobuf:"bytes,1,opt,name=username_or_email,json=usernameOrEmail,proto3" json:"username_or_email,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmationRequest) Reset()         { *m = ConfirmationRequest{} }
func (m *ConfirmationRequest) String() string { return proto.CompactTextString(m) }
func (*ConfirmationRequest) ProtoMessage()    {}
func (*ConfirmationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e27dc8f5aa0b42d7, []int{4}
}
func (m *ConfirmationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmationRequest.Unmarshal(m, b)
}
func (m *ConfirmationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmationRequest.Marshal(b, m, deterministic)
}
func (m *ConfirmationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmationRequest.Merge(m, src)
}
func (m *ConfirmationRequest) XXX_Size() int {
	return xxx_messageInfo_ConfirmationRequest.Size(m)
}
func (m *ConfirmationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmationRequest proto.InternalMessageInfo

func (m *ConfirmationRequest) GetUsernameOrEmail() string {
	if m != nil {
		return m.UsernameOrEmail
	}
	return ""
}

func (m *ConfirmationRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ConfirmationResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmationResponse) Reset()         { *m = ConfirmationResponse{} }
func (m *ConfirmationResponse) String() string { return proto.CompactTextString(m) }
func (*ConfirmationResponse) ProtoMessage()    {}
func (*ConfirmationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e27dc8f5aa0b42d7, []int{5}
}
func (m *ConfirmationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmationResponse.Unmarshal(m, b)
}
func (m *ConfirmationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmationResponse.Marshal(b, m, deterministic)
}
func (m *ConfirmationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmationResponse.Merge(m, src)
}
func (m *ConfirmationResponse) XXX_Size() int {
	return xxx_messageInfo_ConfirmationResponse.Size(m)
}
func (m *ConfirmationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmationResponse proto.InternalMessageInfo

func (m *ConfirmationResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MData)(nil), "authorization.MData")
	proto.RegisterType((*AuthSignUpRequest)(nil), "authorization.AuthSignUpRequest")
	proto.RegisterType((*AuthSignInRequest)(nil), "authorization.AuthSignInRequest")
	proto.RegisterType((*AuthSignUpRsp)(nil), "authorization.AuthSignUpRsp")
	proto.RegisterType((*ConfirmationRequest)(nil), "authorization.ConfirmationRequest")
	proto.RegisterType((*ConfirmationResponse)(nil), "authorization.ConfirmationResponse")
}

func init() { proto.RegisterFile("authorization_service.proto", fileDescriptor_e27dc8f5aa0b42d7) }

var fileDescriptor_e27dc8f5aa0b42d7 = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x52, 0x13, 0x31,
	0x18, 0x9d, 0xa5, 0xf4, 0x67, 0xbf, 0x52, 0x85, 0x50, 0x74, 0x29, 0x3a, 0xd6, 0x7a, 0x61, 0x47,
	0x67, 0x8a, 0x83, 0x4f, 0xd0, 0x02, 0x8e, 0x55, 0x11, 0x66, 0x11, 0x9d, 0xe1, 0x66, 0x27, 0x6d,
	0xd2, 0x6d, 0x86, 0x6e, 0xb2, 0x26, 0x29, 0x52, 0x9f, 0xc0, 0x37, 0xf4, 0x71, 0x74, 0x92, 0xec,
	0xd2, 0x16, 0x41, 0x2f, 0xbc, 0x6a, 0xcf, 0x39, 0xd9, 0xef, 0xe7, 0x9c, 0xdd, 0xc0, 0x0e, 0x9e,
	0xea, 0xb1, 0x90, 0xec, 0x3b, 0xd6, 0x4c, 0xf0, 0x48, 0x51, 0x79, 0xc9, 0x86, 0xb4, 0x93, 0x4a,
	0xa1, 0x05, 0xaa, 0x2d, 0x89, 0x8d, 0xed, 0x58, 0x88, 0x78, 0x42, 0x77, 0xad, 0x38, 0x98, 0x8e,
	0x76, 0x31, 0x9f, 0xb9, 0x93, 0xad, 0x1f, 0x05, 0x28, 0x1e, 0x1d, 0x60, 0x8d, 0xd1, 0x0e, 0xf8,
	0xa6, 0x08, 0x95, 0x11, 0x23, 0x81, 0xd7, 0xf4, 0xda, 0x7e, 0x58, 0x71, 0x44, 0x9f, 0xa0, 0x27,
	0x50, 0x1d, 0x4e, 0x18, 0xe5, 0x3a, 0xc2, 0x84, 0xc8, 0x60, 0xc5, 0xca, 0xe0, 0xa8, 0x2e, 0x21,
	0x12, 0x3d, 0x84, 0xb2, 0xe9, 0x69, 0x9e, 0x2d, 0x34, 0xbd, 0x76, 0x21, 0x2c, 0x19, 0xd8, 0x27,
	0xe8, 0x31, 0x80, 0xa2, 0x4a, 0x99, 0x19, 0x19, 0x09, 0x56, 0xad, 0xe6, 0x67, 0x4c, 0x9f, 0xa0,
	0xa7, 0xb0, 0x26, 0xe9, 0x90, 0xb2, 0x4b, 0x1a, 0x69, 0x96, 0xd0, 0xa0, 0x68, 0x0f, 0x54, 0x33,
	0xee, 0x13, 0x4b, 0xa8, 0x29, 0x3d, 0x55, 0x6e, 0xac, 0x92, 0x2b, 0x6d, 0x60, 0x9f, 0xa0, 0x16,
	0xd4, 0xb2, 0xa1, 0x12, 0x15, 0x1b, 0xb9, 0xec, 0x1e, 0x76, 0xe4, 0x91, 0x8a, 0xfb, 0x04, 0x6d,
	0x41, 0x89, 0xa9, 0x68, 0x20, 0x74, 0x50, 0x69, 0x7a, 0xed, 0x4a, 0x58, 0x64, 0xaa, 0x27, 0x34,
	0xaa, 0x43, 0x71, 0x82, 0x67, 0x54, 0x06, 0x7e, 0xd3, 0x6b, 0x17, 0x43, 0x07, 0xd0, 0x03, 0x28,
	0xb9, 0x67, 0x03, 0xb0, 0x0b, 0x66, 0x08, 0x6d, 0x43, 0x85, 0xa9, 0x08, 0x93, 0x84, 0xf1, 0xa0,
	0x6a, 0xcb, 0x94, 0x99, 0xea, 0x1a, 0x88, 0x1a, 0x50, 0x99, 0x60, 0x1e, 0xa7, 0x78, 0x78, 0x11,
	0xd4, 0x9c, 0x69, 0x39, 0x46, 0xcf, 0x61, 0x3d, 0xa5, 0x32, 0x89, 0xac, 0x31, 0x17, 0x74, 0x66,
	0x46, 0xbc, 0x67, 0x47, 0xac, 0x19, 0xbe, 0x3b, 0xd5, 0xe3, 0xf7, 0x74, 0xd6, 0x27, 0xef, 0x56,
	0x2b, 0x6b, 0xeb, 0xb5, 0xd6, 0xcf, 0x15, 0xd8, 0x30, 0xdc, 0x29, 0x8b, 0xf9, 0x59, 0x1a, 0xd2,
	0xaf, 0x53, 0xaa, 0xb4, 0x69, 0x60, 0xd6, 0xe5, 0x38, 0xa1, 0x79, 0x2a, 0x39, 0x36, 0x5a, 0x8a,
	0x95, 0xfa, 0x26, 0x24, 0xc9, 0x22, 0xb9, 0xc6, 0x66, 0x97, 0x98, 0x72, 0x42, 0xa5, 0xcd, 0xc3,
	0x0f, 0x33, 0x64, 0x4c, 0x23, 0x58, 0xd3, 0x48, 0x8c, 0xa2, 0x01, 0x93, 0x7a, 0x6c, 0x23, 0xf1,
	0xc3, 0xaa, 0x21, 0x8f, 0x47, 0x3d, 0x43, 0x19, 0x77, 0x68, 0x82, 0xd9, 0xc4, 0xa6, 0xe1, 0x87,
	0x0e, 0x18, 0x36, 0x1d, 0x0b, 0x4e, 0x6d, 0x0a, 0x7e, 0xe8, 0x80, 0x09, 0x70, 0x28, 0xa6, 0x5c,
	0xcb, 0x59, 0x34, 0x14, 0x84, 0xda, 0x0c, 0xfc, 0xb0, 0x9a, 0x71, 0xfb, 0x82, 0x50, 0x33, 0x0a,
	0xbe, 0xc4, 0x1a, 0x4b, 0x9b, 0x81, 0x1f, 0x66, 0xc8, 0xbc, 0x1a, 0x23, 0x26, 0x95, 0x8e, 0xec,
	0x72, 0xbe, 0xd5, 0x7c, 0xcb, 0x7c, 0x34, 0xdb, 0xed, 0x80, 0x3f, 0xc1, 0xb9, 0x0a, 0xb9, 0xb7,
	0x99, 0xf8, 0x12, 0x4a, 0x49, 0x44, 0xb0, 0xc6, 0x36, 0x90, 0xea, 0x5e, 0xbd, 0xb3, 0xf4, 0xca,
	0x77, 0xec, 0x3b, 0x1d, 0x16, 0x13, 0xf3, 0xd3, 0xba, 0x9a, 0x1b, 0xdb, 0xe7, 0xff, 0x6b, 0xec,
	0xbc, 0x73, 0xe1, 0xdf, 0x9d, 0x3f, 0x43, 0x6d, 0x21, 0x52, 0x95, 0xce, 0xad, 0xf5, 0x16, 0xad,
	0xdd, 0x85, 0xcd, 0xa1, 0xe0, 0x23, 0x26, 0x13, 0xf7, 0x35, 0xd3, 0xab, 0x94, 0x49, 0x6a, 0x5b,
	0x17, 0x42, 0xb4, 0x28, 0x1d, 0x5a, 0xa5, 0x75, 0x06, 0x9b, 0xfb, 0x0b, 0x6c, 0xbe, 0xd3, 0x0b,
	0xd8, 0xc8, 0x77, 0x88, 0x84, 0x8c, 0x16, 0x3b, 0xdd, 0xcf, 0x85, 0x63, 0x79, 0x68, 0x7b, 0x22,
	0x58, 0xb5, 0x81, 0xb9, 0xfd, 0xec, 0xff, 0xd6, 0x2b, 0xa8, 0x2f, 0x97, 0x55, 0xa9, 0xe0, 0x8a,
	0xa2, 0x00, 0xca, 0x09, 0x55, 0x0a, 0xc7, 0xb9, 0x55, 0x39, 0xdc, 0xfb, 0xe5, 0xc1, 0x7a, 0x78,
	0xb2, 0xdf, 0x5d, 0xb4, 0x00, 0x1d, 0x00, 0xcc, 0xfd, 0x46, 0xcd, 0x1b, 0x06, 0xfd, 0x11, 0x45,
	0xa3, 0xde, 0x71, 0x17, 0x54, 0x27, 0xbf, 0xa0, 0x3a, 0x5d, 0x3e, 0x43, 0x1f, 0xe6, 0x55, 0xce,
	0xd2, 0x3b, 0xab, 0x5c, 0x7f, 0x29, 0x8d, 0x47, 0x77, 0x9f, 0x50, 0x29, 0xfa, 0x02, 0x6b, 0x8b,
	0xab, 0xa1, 0xd6, 0x8d, 0xd3, 0xb7, 0xd8, 0xd9, 0x78, 0xf6, 0xd7, 0x33, 0xce, 0x9b, 0xde, 0x1b,
	0xd8, 0xba, 0xf5, 0x2a, 0xee, 0x35, 0xce, 0x4f, 0xcc, 0x3a, 0x4b, 0xd6, 0x9c, 0x3a, 0xed, 0xed,
	0xca, 0x89, 0x77, 0xbe, 0x7c, 0x49, 0x0f, 0x4a, 0x76, 0xf9, 0xd7, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xc6, 0xcc, 0xaf, 0xef, 0xd9, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCAuthorizationClient is the client API for RPCAuthorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCAuthorizationClient interface {
	AuthSignIn(ctx context.Context, in *AuthSignInRequest, opts ...grpc.CallOption) (*types.Any, error)
	AuthSignUp(ctx context.Context, in *AuthSignUpRequest, opts ...grpc.CallOption) (*AuthSignUpRsp, error)
	Confirmation(ctx context.Context, in *ConfirmationRequest, opts ...grpc.CallOption) (*ConfirmationResponse, error)
}

type rPCAuthorizationClient struct {
	cc *grpc.ClientConn
}

func NewRPCAuthorizationClient(cc *grpc.ClientConn) RPCAuthorizationClient {
	return &rPCAuthorizationClient{cc}
}

func (c *rPCAuthorizationClient) AuthSignIn(ctx context.Context, in *AuthSignInRequest, opts ...grpc.CallOption) (*types.Any, error) {
	out := new(types.Any)
	err := c.cc.Invoke(ctx, "/authorization.RPCAuthorization/AuthSignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthorizationClient) AuthSignUp(ctx context.Context, in *AuthSignUpRequest, opts ...grpc.CallOption) (*AuthSignUpRsp, error) {
	out := new(AuthSignUpRsp)
	err := c.cc.Invoke(ctx, "/authorization.RPCAuthorization/AuthSignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthorizationClient) Confirmation(ctx context.Context, in *ConfirmationRequest, opts ...grpc.CallOption) (*ConfirmationResponse, error) {
	out := new(ConfirmationResponse)
	err := c.cc.Invoke(ctx, "/authorization.RPCAuthorization/Confirmation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCAuthorizationServer is the server API for RPCAuthorization service.
type RPCAuthorizationServer interface {
	AuthSignIn(context.Context, *AuthSignInRequest) (*types.Any, error)
	AuthSignUp(context.Context, *AuthSignUpRequest) (*AuthSignUpRsp, error)
	Confirmation(context.Context, *ConfirmationRequest) (*ConfirmationResponse, error)
}

// UnimplementedRPCAuthorizationServer can be embedded to have forward compatible implementations.
type UnimplementedRPCAuthorizationServer struct {
}

func (*UnimplementedRPCAuthorizationServer) AuthSignIn(ctx context.Context, req *AuthSignInRequest) (*types.Any, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthSignIn not implemented")
}
func (*UnimplementedRPCAuthorizationServer) AuthSignUp(ctx context.Context, req *AuthSignUpRequest) (*AuthSignUpRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthSignUp not implemented")
}
func (*UnimplementedRPCAuthorizationServer) Confirmation(ctx context.Context, req *ConfirmationRequest) (*ConfirmationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confirmation not implemented")
}

func RegisterRPCAuthorizationServer(s *grpc.Server, srv RPCAuthorizationServer) {
	s.RegisterService(&_RPCAuthorization_serviceDesc, srv)
}

func _RPCAuthorization_AuthSignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthSignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthorizationServer).AuthSignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.RPCAuthorization/AuthSignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthorizationServer).AuthSignIn(ctx, req.(*AuthSignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthorization_AuthSignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthSignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthorizationServer).AuthSignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.RPCAuthorization/AuthSignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthorizationServer).AuthSignUp(ctx, req.(*AuthSignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthorization_Confirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthorizationServer).Confirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.RPCAuthorization/Confirmation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthorizationServer).Confirmation(ctx, req.(*ConfirmationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCAuthorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authorization.RPCAuthorization",
	HandlerType: (*RPCAuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthSignIn",
			Handler:    _RPCAuthorization_AuthSignIn_Handler,
		},
		{
			MethodName: "AuthSignUp",
			Handler:    _RPCAuthorization_AuthSignUp_Handler,
		},
		{
			MethodName: "Confirmation",
			Handler:    _RPCAuthorization_Confirmation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorization_service.proto",
}
