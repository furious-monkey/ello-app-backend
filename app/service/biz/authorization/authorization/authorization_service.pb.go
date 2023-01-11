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

type AuthSignUpRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Gender               string   `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	DateOfBirth          string   `protobuf:"bytes,4,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	CountryCode          string   `protobuf:"bytes,7,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Avatar               string   `protobuf:"bytes,8,opt,name=avatar,proto3" json:"avatar,omitempty"`
	FirstName            string   `protobuf:"bytes,9,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,10,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Username             string   `protobuf:"bytes,11,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthSignUpRequest) Reset()         { *m = AuthSignUpRequest{} }
func (m *AuthSignUpRequest) String() string { return proto.CompactTextString(m) }
func (*AuthSignUpRequest) ProtoMessage()    {}
func (*AuthSignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e27dc8f5aa0b42d7, []int{0}
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

func (m *AuthSignUpRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
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

func (m *AuthSignUpRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type AuthSignInRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthSignInRequest) Reset()         { *m = AuthSignInRequest{} }
func (m *AuthSignInRequest) String() string { return proto.CompactTextString(m) }
func (*AuthSignInRequest) ProtoMessage()    {}
func (*AuthSignInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e27dc8f5aa0b42d7, []int{1}
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
	return fileDescriptor_e27dc8f5aa0b42d7, []int{2}
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
	return fileDescriptor_e27dc8f5aa0b42d7, []int{3}
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
	return fileDescriptor_e27dc8f5aa0b42d7, []int{4}
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
	proto.RegisterType((*AuthSignUpRequest)(nil), "authorization.AuthSignUpRequest")
	proto.RegisterType((*AuthSignInRequest)(nil), "authorization.AuthSignInRequest")
	proto.RegisterType((*AuthSignUpRsp)(nil), "authorization.AuthSignUpRsp")
	proto.RegisterType((*ConfirmationRequest)(nil), "authorization.ConfirmationRequest")
	proto.RegisterType((*ConfirmationResponse)(nil), "authorization.ConfirmationResponse")
}

func init() { proto.RegisterFile("authorization_service.proto", fileDescriptor_e27dc8f5aa0b42d7) }

var fileDescriptor_e27dc8f5aa0b42d7 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0xda, 0xad, 0x1f, 0x77, 0xab, 0x60, 0xde, 0x18, 0x56, 0x0a, 0x52, 0x09, 0x2f, 0x13,
	0x0f, 0x29, 0x82, 0x5f, 0xb0, 0x56, 0x9d, 0x98, 0x40, 0xac, 0xea, 0x54, 0x90, 0xf6, 0x12, 0xb9,
	0xcd, 0x6d, 0x6a, 0xa9, 0x89, 0x8d, 0xed, 0x0c, 0xc6, 0x1b, 0xbf, 0x8d, 0x3f, 0x86, 0x62, 0x2f,
	0x6b, 0xc2, 0x3e, 0xde, 0x72, 0xce, 0xb9, 0xbe, 0xbe, 0xf7, 0xe4, 0x18, 0xfa, 0x2c, 0x37, 0x6b,
	0xa1, 0xf8, 0x6f, 0x66, 0xb8, 0xc8, 0x22, 0x8d, 0xea, 0x9a, 0x2f, 0x31, 0x94, 0x4a, 0x18, 0x41,
	0x7a, 0x35, 0xd1, 0xef, 0x27, 0x42, 0x24, 0x1b, 0x1c, 0x5a, 0x71, 0x91, 0xaf, 0x86, 0x98, 0x4a,
	0x73, 0xe3, 0x6a, 0x83, 0xbf, 0x0d, 0x38, 0x38, 0xcd, 0xcd, 0xfa, 0x92, 0x27, 0xd9, 0x5c, 0xce,
	0xf0, 0x47, 0x8e, 0xda, 0x90, 0x97, 0xd0, 0xce, 0x35, 0xaa, 0x88, 0xc7, 0xd4, 0x1b, 0x78, 0x27,
	0xcd, 0x59, 0xab, 0x80, 0xe7, 0x31, 0xf1, 0xa1, 0x23, 0x99, 0xd6, 0x3f, 0x85, 0x8a, 0x69, 0x63,
	0xe0, 0x9d, 0x74, 0x67, 0x77, 0x98, 0x1c, 0x43, 0x2b, 0xc1, 0x2c, 0x46, 0x45, 0x9b, 0x56, 0xb9,
	0x45, 0x24, 0x80, 0x5e, 0xcc, 0x0c, 0x46, 0x62, 0x15, 0x2d, 0xb8, 0x32, 0x6b, 0xba, 0x63, 0xe5,
	0xbd, 0x82, 0xbc, 0x58, 0x8d, 0x0a, 0x8a, 0x1c, 0xc1, 0x2e, 0xa6, 0x8c, 0x6f, 0xe8, 0xae, 0xd5,
	0x1c, 0x28, 0x58, 0xb9, 0x16, 0x19, 0xd2, 0x96, 0x63, 0x2d, 0x20, 0x6f, 0x60, 0x7f, 0x29, 0xf2,
	0xcc, 0xa8, 0x9b, 0x68, 0x29, 0x62, 0xa4, 0x6d, 0xd7, 0xee, 0x96, 0x1b, 0x8b, 0x18, 0x8b, 0x51,
	0xd8, 0x35, 0x33, 0x4c, 0xd1, 0x8e, 0x1b, 0xc5, 0x21, 0xf2, 0x1a, 0x60, 0xc5, 0x95, 0x36, 0x51,
	0xc6, 0x52, 0xa4, 0x5d, 0xab, 0x75, 0x2d, 0xf3, 0x95, 0xa5, 0x48, 0xfa, 0xd0, 0xdd, 0xb0, 0x52,
	0x05, 0xb7, 0x5e, 0x41, 0x58, 0xd1, 0x87, 0x4e, 0x61, 0x82, 0xd5, 0xf6, 0x9c, 0x56, 0xe2, 0xe0,
	0xf3, 0xd6, 0xc4, 0xf3, 0xac, 0x34, 0xb1, 0x7a, 0xc0, 0xab, 0x1f, 0x78, 0xca, 0xc7, 0xe0, 0x1b,
	0xf4, 0x2a, 0x7f, 0x44, 0xcb, 0xad, 0x39, 0x5e, 0xd5, 0x9c, 0x21, 0x1c, 0x2e, 0x45, 0xb6, 0xe2,
	0x2a, 0x75, 0x19, 0xc0, 0x5f, 0x92, 0x2b, 0xb4, 0xdd, 0x9a, 0x33, 0x52, 0x95, 0x26, 0x56, 0x09,
	0xe6, 0x70, 0x38, 0xae, 0xb0, 0xe5, 0x98, 0xef, 0xe0, 0xa0, 0x1c, 0x2b, 0x12, 0x2a, 0xaa, 0xde,
	0xf4, 0xac, 0x14, 0x2e, 0xd4, 0xc4, 0xde, 0x49, 0x60, 0xc7, 0x5a, 0xee, 0x46, 0xb6, 0xdf, 0xc1,
	0x7b, 0x38, 0xaa, 0xb7, 0xd5, 0x52, 0x64, 0x1a, 0x09, 0x85, 0x76, 0x8a, 0x5a, 0xb3, 0xa4, 0xdc,
	0xbe, 0x84, 0x1f, 0xfe, 0x34, 0xe0, 0xf9, 0x6c, 0x3a, 0x3e, 0xad, 0xa6, 0x94, 0x9c, 0x01, 0x6c,
	0x2d, 0x24, 0x83, 0xb0, 0x96, 0xe1, 0xf0, 0x9e, 0xbb, 0xfe, 0x71, 0xe8, 0x62, 0x1d, 0x96, 0xb1,
	0x0e, 0x27, 0x45, 0xac, 0xc9, 0x97, 0x6d, 0x9f, 0xb9, 0x7c, 0xb4, 0xcf, 0x5d, 0xd4, 0xfd, 0x57,
	0x8f, 0x57, 0x68, 0x49, 0xbe, 0xc3, 0x7e, 0x75, 0x39, 0x12, 0xfc, 0x57, 0xfd, 0x80, 0xa1, 0xfe,
	0xdb, 0x27, 0x6b, 0x9c, 0x3b, 0xa3, 0x33, 0x78, 0xf1, 0xe0, 0x13, 0x1e, 0xf9, 0x57, 0xd3, 0x62,
	0xa1, 0x9a, 0x39, 0x97, 0x4e, 0xfb, 0xd4, 0x98, 0x7a, 0x57, 0xf5, 0xc7, 0xbd, 0x68, 0xd9, 0xf5,
	0x3f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x35, 0x53, 0xbe, 0xff, 0x11, 0x04, 0x00, 0x00,
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
	AuthSignIn(ctx context.Context, in *AuthSignInRequest, opts ...grpc.CallOption) (*types.Empty, error)
	AuthSignUp(ctx context.Context, in *AuthSignUpRequest, opts ...grpc.CallOption) (*AuthSignUpRsp, error)
	Confirmation(ctx context.Context, in *ConfirmationRequest, opts ...grpc.CallOption) (*ConfirmationResponse, error)
}

type rPCAuthorizationClient struct {
	cc *grpc.ClientConn
}

func NewRPCAuthorizationClient(cc *grpc.ClientConn) RPCAuthorizationClient {
	return &rPCAuthorizationClient{cc}
}

func (c *rPCAuthorizationClient) AuthSignIn(ctx context.Context, in *AuthSignInRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
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
	AuthSignIn(context.Context, *AuthSignInRequest) (*types.Empty, error)
	AuthSignUp(context.Context, *AuthSignUpRequest) (*AuthSignUpRsp, error)
	Confirmation(context.Context, *ConfirmationRequest) (*ConfirmationResponse, error)
}

// UnimplementedRPCAuthorizationServer can be embedded to have forward compatible implementations.
type UnimplementedRPCAuthorizationServer struct {
}

func (*UnimplementedRPCAuthorizationServer) AuthSignIn(ctx context.Context, req *AuthSignInRequest) (*types.Empty, error) {
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
