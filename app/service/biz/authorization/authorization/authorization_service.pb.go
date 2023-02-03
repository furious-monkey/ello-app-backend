// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: authorization_service.proto

package authorization

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
	Type                 string   `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Kind                 string   `protobuf:"bytes,8,opt,name=kind,proto3" json:"kind,omitempty"`
	CountryCode          string   `protobuf:"bytes,9,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Avatar               string   `protobuf:"bytes,10,opt,name=avatar,proto3" json:"avatar,omitempty"`
	FirstName            string   `protobuf:"bytes,11,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,12,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	MData                *MData   `protobuf:"bytes,13,opt,name=m_data,json=mData,proto3" json:"m_data,omitempty"`
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

func (m *AuthSignUpRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AuthSignUpRequest) GetKind() string {
	if m != nil {
		return m.Kind
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

type AuthSignInResponse struct {
	PredicateName         string     `protobuf:"bytes,1,opt,name=predicate_name,json=predicateName,proto3" json:"predicate_name,omitempty"`
	Constructor           int32      `protobuf:"varint,2,opt,name=constructor,proto3" json:"constructor,omitempty"`
	OtherwiseReloginDays  int32      `protobuf:"varint,3,opt,name=otherwise_relogin_days,json=otherwiseReloginDays,proto3" json:"otherwise_relogin_days,omitempty"`
	TmpSessions           *types.Any `protobuf:"bytes,4,opt,name=tmp_sessions,json=tmpSessions,proto3" json:"tmp_sessions,omitempty"`
	TermsOfService        *types.Any `protobuf:"bytes,5,opt,name=terms_of_service,json=termsOfService,proto3" json:"terms_of_service,omitempty"`
	User                  *types.Any `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	SetupPasswordRequired bool       `protobuf:"varint,7,opt,name=setup_password_required,json=setupPasswordRequired,proto3" json:"setup_password_required,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}   `json:"-"`
	XXX_unrecognized      []byte     `json:"-"`
	XXX_sizecache         int32      `json:"-"`
}

func (m *AuthSignInResponse) Reset()         { *m = AuthSignInResponse{} }
func (m *AuthSignInResponse) String() string { return proto.CompactTextString(m) }
func (*AuthSignInResponse) ProtoMessage()    {}
func (*AuthSignInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e27dc8f5aa0b42d7, []int{3}
}
func (m *AuthSignInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthSignInResponse.Unmarshal(m, b)
}
func (m *AuthSignInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthSignInResponse.Marshal(b, m, deterministic)
}
func (m *AuthSignInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthSignInResponse.Merge(m, src)
}
func (m *AuthSignInResponse) XXX_Size() int {
	return xxx_messageInfo_AuthSignInResponse.Size(m)
}
func (m *AuthSignInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthSignInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthSignInResponse proto.InternalMessageInfo

func (m *AuthSignInResponse) GetPredicateName() string {
	if m != nil {
		return m.PredicateName
	}
	return ""
}

func (m *AuthSignInResponse) GetConstructor() int32 {
	if m != nil {
		return m.Constructor
	}
	return 0
}

func (m *AuthSignInResponse) GetOtherwiseReloginDays() int32 {
	if m != nil {
		return m.OtherwiseReloginDays
	}
	return 0
}

func (m *AuthSignInResponse) GetTmpSessions() *types.Any {
	if m != nil {
		return m.TmpSessions
	}
	return nil
}

func (m *AuthSignInResponse) GetTermsOfService() *types.Any {
	if m != nil {
		return m.TermsOfService
	}
	return nil
}

func (m *AuthSignInResponse) GetUser() *types.Any {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *AuthSignInResponse) GetSetupPasswordRequired() bool {
	if m != nil {
		return m.SetupPasswordRequired
	}
	return false
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
	return fileDescriptor_e27dc8f5aa0b42d7, []int{4}
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
	return fileDescriptor_e27dc8f5aa0b42d7, []int{5}
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
	return fileDescriptor_e27dc8f5aa0b42d7, []int{6}
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
	proto.RegisterType((*AuthSignInResponse)(nil), "authorization.AuthSignInResponse")
	proto.RegisterType((*AuthSignUpRsp)(nil), "authorization.AuthSignUpRsp")
	proto.RegisterType((*ConfirmationRequest)(nil), "authorization.ConfirmationRequest")
	proto.RegisterType((*ConfirmationResponse)(nil), "authorization.ConfirmationResponse")
}

func init() { proto.RegisterFile("authorization_service.proto", fileDescriptor_e27dc8f5aa0b42d7) }

var fileDescriptor_e27dc8f5aa0b42d7 = []byte{
	// 907 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x51, 0x6f, 0x23, 0x35,
	0x10, 0x56, 0xda, 0x24, 0xcd, 0x4e, 0x92, 0xa3, 0xe7, 0x6b, 0xef, 0xf6, 0x52, 0x10, 0x21, 0x08,
	0x11, 0x81, 0x94, 0xa2, 0x03, 0xc1, 0x1b, 0x52, 0xda, 0x3b, 0x44, 0x38, 0x8e, 0x56, 0x2e, 0x05,
	0xe9, 0x5e, 0x2c, 0x77, 0xed, 0x24, 0x56, 0xb3, 0xeb, 0x3d, 0xdb, 0xe9, 0xdd, 0xf2, 0x03, 0x10,
	0xbf, 0x96, 0xbf, 0x00, 0xf2, 0x78, 0xb7, 0xd9, 0x1c, 0x57, 0x78, 0xe0, 0x29, 0x3b, 0xdf, 0x37,
	0xb6, 0x67, 0xbe, 0xf9, 0xec, 0xc0, 0x11, 0x5f, 0xbb, 0xa5, 0x36, 0xea, 0x37, 0xee, 0x94, 0xce,
	0x98, 0x95, 0xe6, 0x46, 0x25, 0x72, 0x92, 0x1b, 0xed, 0x34, 0xe9, 0x6f, 0x91, 0x83, 0xc7, 0x0b,
	0xad, 0x17, 0x2b, 0x79, 0x8c, 0xe4, 0xd5, 0x7a, 0x7e, 0xcc, 0xb3, 0x22, 0x64, 0x0e, 0xfa, 0x69,
	0xc1, 0x5c, 0x91, 0x97, 0x0b, 0x47, 0x7f, 0xec, 0x42, 0xeb, 0xc5, 0x53, 0xee, 0x38, 0x39, 0x82,
	0xc8, 0xef, 0x29, 0x0d, 0x53, 0x22, 0x6e, 0x0c, 0x1b, 0xe3, 0x88, 0x76, 0x02, 0x30, 0x13, 0xe4,
	0x43, 0xe8, 0x26, 0x2b, 0x25, 0x33, 0xc7, 0xb8, 0x10, 0x26, 0xde, 0x41, 0x1a, 0x02, 0x34, 0x15,
	0xc2, 0x90, 0x47, 0xb0, 0xe7, 0x4b, 0xf0, 0x6b, 0x77, 0x87, 0x8d, 0xf1, 0x2e, 0x6d, 0xfb, 0x70,
	0x26, 0xc8, 0x07, 0x00, 0x56, 0x5a, 0xeb, 0x4b, 0x56, 0x22, 0x6e, 0x22, 0x17, 0x95, 0xc8, 0x4c,
	0x90, 0x8f, 0xa0, 0x67, 0x64, 0x22, 0xd5, 0x8d, 0x64, 0x4e, 0xa5, 0x32, 0x6e, 0x61, 0x42, 0xb7,
	0xc4, 0x7e, 0x56, 0xa9, 0xf4, 0x5b, 0xaf, 0x6d, 0x28, 0xab, 0x1d, 0xb6, 0xf6, 0xe1, 0x4c, 0x90,
	0x11, 0xf4, 0xcb, 0xa2, 0x52, 0xbb, 0xf0, 0xf4, 0x5e, 0x58, 0x1c, 0xc0, 0x17, 0x76, 0x31, 0x13,
	0xe4, 0x10, 0xda, 0xca, 0xb2, 0x2b, 0xed, 0xe2, 0xce, 0xb0, 0x31, 0xee, 0xd0, 0x96, 0xb2, 0x27,
	0xda, 0x91, 0x03, 0x68, 0xad, 0x78, 0x21, 0x4d, 0x1c, 0x0d, 0x1b, 0xe3, 0x16, 0x0d, 0x01, 0x79,
	0x08, 0xed, 0xb0, 0x36, 0x06, 0x6c, 0xb0, 0x8c, 0xc8, 0x63, 0xe8, 0x28, 0xcb, 0xb8, 0x48, 0x55,
	0x16, 0x77, 0x71, 0x9b, 0x3d, 0x65, 0xa7, 0x3e, 0x24, 0x03, 0xe8, 0xac, 0x78, 0xb6, 0xc8, 0x79,
	0x72, 0x1d, 0xf7, 0x83, 0x68, 0x55, 0x4c, 0x3e, 0x85, 0xfd, 0x5c, 0x9a, 0x94, 0xa1, 0x30, 0xd7,
	0xb2, 0xf0, 0x25, 0xde, 0xc3, 0x12, 0xfb, 0x1e, 0x9f, 0xae, 0xdd, 0xf2, 0xb9, 0x2c, 0x66, 0xe2,
	0x87, 0x66, 0xa7, 0xb7, 0xdf, 0x1f, 0xfd, 0xb5, 0x03, 0xf7, 0x3d, 0x76, 0xa1, 0x16, 0xd9, 0x65,
	0x4e, 0xe5, 0xab, 0xb5, 0xb4, 0xce, 0x1f, 0xe0, 0xdb, 0xcd, 0x78, 0x2a, 0xab, 0xa9, 0x54, 0xb1,
	0xe7, 0x72, 0x6e, 0xed, 0x6b, 0x6d, 0x44, 0x39, 0x92, 0xdb, 0xd8, 0xf7, 0xb2, 0x90, 0x99, 0x90,
	0x06, 0xe7, 0x11, 0xd1, 0x32, 0xf2, 0xa2, 0x09, 0xee, 0x24, 0xd3, 0x73, 0x76, 0xa5, 0x8c, 0x5b,
	0xe2, 0x48, 0x22, 0xda, 0xf5, 0xe0, 0xd9, 0xfc, 0xc4, 0x43, 0x5e, 0x1d, 0x99, 0x72, 0xb5, 0xc2,
	0x69, 0x44, 0x34, 0x04, 0x1e, 0xcd, 0x97, 0x3a, 0x93, 0x38, 0x85, 0x88, 0x86, 0x80, 0x10, 0x68,
	0x7a, 0x3b, 0xa1, 0xf6, 0x11, 0xc5, 0x6f, 0x8f, 0x5d, 0xab, 0x4c, 0xa0, 0xe4, 0x11, 0xc5, 0x6f,
	0x3f, 0xe8, 0x44, 0xaf, 0x33, 0x67, 0x0a, 0x96, 0x68, 0x21, 0x51, 0xf8, 0x88, 0x76, 0x4b, 0xec,
	0x54, 0x0b, 0xe9, 0x4b, 0xe6, 0x37, 0xdc, 0x71, 0x53, 0xc9, 0x1f, 0x22, 0x6f, 0xa1, 0xb9, 0x32,
	0xd6, 0x31, 0x14, 0xa1, 0x8b, 0x5c, 0x84, 0xc8, 0x4f, 0x5e, 0x85, 0x23, 0x88, 0x56, 0xbc, 0x62,
	0x7b, 0xd5, 0x0c, 0x4a, 0xf2, 0x73, 0x68, 0xa7, 0x4c, 0x70, 0xc7, 0x71, 0x3a, 0xdd, 0x27, 0x07,
	0x93, 0xad, 0x9b, 0x32, 0x41, 0xef, 0xd3, 0x56, 0xea, 0x7f, 0x46, 0x6f, 0x36, 0x03, 0x98, 0x65,
	0xff, 0x77, 0x00, 0x9b, 0x93, 0x77, 0xff, 0xfb, 0xe4, 0x3f, 0x77, 0x80, 0xd4, 0x8f, 0xb6, 0xb9,
	0xce, 0xac, 0x24, 0x9f, 0xc0, 0xbd, 0xdc, 0x48, 0xa1, 0x12, 0x3f, 0xb1, 0x5a, 0x05, 0xfd, 0x5b,
	0x14, 0x9b, 0x1c, 0x42, 0x37, 0xd1, 0x99, 0x75, 0x66, 0x9d, 0x38, 0x1d, 0x6e, 0x67, 0x8b, 0xd6,
	0x21, 0xf2, 0x15, 0x3c, 0xd4, 0x6e, 0x29, 0xcd, 0x6b, 0x65, 0x25, 0x33, 0x72, 0xa5, 0x17, 0x2a,
	0x63, 0x82, 0x17, 0x16, 0x8b, 0x6b, 0xd1, 0x83, 0x5b, 0x96, 0x06, 0xf2, 0x29, 0x2f, 0x2c, 0xf9,
	0x06, 0x7a, 0x2e, 0xcd, 0x59, 0x79, 0x5b, 0x2d, 0x5a, 0xc5, 0x37, 0x12, 0x5e, 0x97, 0x49, 0xf5,
	0xba, 0x4c, 0xa6, 0x59, 0x41, 0xbb, 0x2e, 0xcd, 0x2f, 0xca, 0x44, 0xf2, 0x2d, 0xec, 0x3b, 0x69,
	0x52, 0xeb, 0x5d, 0x56, 0x3e, 0x54, 0xe8, 0xa5, 0xbb, 0x16, 0xdf, 0xc3, 0xec, 0xb3, 0xf9, 0x45,
	0xc8, 0x25, 0x63, 0x68, 0x7a, 0x8d, 0xd1, 0x69, 0x77, 0xad, 0xc1, 0x0c, 0xf2, 0x35, 0x3c, 0xb2,
	0xd2, 0xad, 0x73, 0x56, 0xe9, 0xce, 0x8c, 0x7c, 0xb5, 0x56, 0x46, 0x86, 0xd7, 0xa0, 0x43, 0x0f,
	0x91, 0x3e, 0x2f, 0x59, 0x5a, 0x92, 0xa3, 0x5f, 0xa0, 0x5f, 0xbb, 0x6b, 0x36, 0xdf, 0x78, 0xbe,
	0x51, 0xf7, 0xfc, 0x31, 0x3c, 0x48, 0x74, 0x36, 0x57, 0x26, 0x0d, 0xaf, 0xae, 0x7c, 0x93, 0x2b,
	0x23, 0x51, 0xe1, 0x5d, 0x4a, 0xea, 0xd4, 0x33, 0x64, 0x46, 0x97, 0xf0, 0xe0, 0xb4, 0x86, 0x56,
	0x26, 0xfa, 0x0c, 0xee, 0x57, 0xa6, 0x61, 0xda, 0xb0, 0xfa, 0x49, 0xef, 0x55, 0xc4, 0x99, 0x79,
	0x86, 0x67, 0x12, 0x68, 0xe2, 0x0d, 0x09, 0x86, 0xc2, 0xef, 0xd1, 0x17, 0x70, 0xb0, 0xbd, 0x6d,
	0x69, 0x90, 0x18, 0xf6, 0x52, 0x69, 0x2d, 0x5f, 0x54, 0xce, 0xa8, 0xc2, 0x27, 0xbf, 0xef, 0xc0,
	0x3e, 0x3d, 0x3f, 0x9d, 0xd6, 0x3d, 0x47, 0x9e, 0x03, 0x6c, 0x5c, 0x46, 0x86, 0x6f, 0x39, 0xf2,
	0x1f, 0xde, 0x1f, 0x1c, 0x4d, 0x52, 0x87, 0xa2, 0x23, 0xc7, 0xb6, 0x37, 0xfb, 0x71, 0xb3, 0xd9,
	0x65, 0x7e, 0xe7, 0x66, 0xb7, 0x2f, 0xd9, 0xe0, 0xfd, 0xbb, 0x33, 0x6c, 0x4e, 0x7e, 0x85, 0x5e,
	0xbd, 0x43, 0x32, 0x7a, 0x2b, 0xfb, 0x1d, 0xaa, 0x0e, 0x3e, 0xfe, 0xd7, 0x9c, 0x20, 0xd1, 0xc9,
	0x77, 0x70, 0xf8, 0xce, 0x7f, 0xce, 0x93, 0xc1, 0xcb, 0x73, 0xdf, 0xda, 0x56, 0x53, 0xa5, 0x01,
	0xbf, 0xdf, 0x39, 0x6f, 0xbc, 0xdc, 0xfe, 0x4f, 0xbd, 0x6a, 0xa3, 0x10, 0x5f, 0xfe, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x37, 0x9c, 0x75, 0xd3, 0x88, 0x07, 0x00, 0x00,
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
	AuthSignIn(ctx context.Context, in *AuthSignInRequest, opts ...grpc.CallOption) (*mtproto.Auth_Authorization, error)
	AuthSignUp(ctx context.Context, in *AuthSignUpRequest, opts ...grpc.CallOption) (*AuthSignUpRsp, error)
	Confirmation(ctx context.Context, in *ConfirmationRequest, opts ...grpc.CallOption) (*ConfirmationResponse, error)
}

type rPCAuthorizationClient struct {
	cc *grpc.ClientConn
}

func NewRPCAuthorizationClient(cc *grpc.ClientConn) RPCAuthorizationClient {
	return &rPCAuthorizationClient{cc}
}

func (c *rPCAuthorizationClient) AuthSignIn(ctx context.Context, in *AuthSignInRequest, opts ...grpc.CallOption) (*mtproto.Auth_Authorization, error) {
	out := new(mtproto.Auth_Authorization)
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
	AuthSignIn(context.Context, *AuthSignInRequest) (*mtproto.Auth_Authorization, error)
	AuthSignUp(context.Context, *AuthSignUpRequest) (*AuthSignUpRsp, error)
	Confirmation(context.Context, *ConfirmationRequest) (*ConfirmationResponse, error)
}

// UnimplementedRPCAuthorizationServer can be embedded to have forward compatible implementations.
type UnimplementedRPCAuthorizationServer struct {
}

func (*UnimplementedRPCAuthorizationServer) AuthSignIn(ctx context.Context, req *AuthSignInRequest) (*mtproto.Auth_Authorization, error) {
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
