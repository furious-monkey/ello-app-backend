// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: accounts.proto

package accounts

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type DeleteAccountReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAccountReq) Reset()         { *m = DeleteAccountReq{} }
func (m *DeleteAccountReq) String() string { return proto.CompactTextString(m) }
func (*DeleteAccountReq) ProtoMessage()    {}
func (*DeleteAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{0}
}
func (m *DeleteAccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAccountReq.Unmarshal(m, b)
}
func (m *DeleteAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAccountReq.Marshal(b, m, deterministic)
}
func (m *DeleteAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAccountReq.Merge(m, src)
}
func (m *DeleteAccountReq) XXX_Size() int {
	return xxx_messageInfo_DeleteAccountReq.Size(m)
}
func (m *DeleteAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAccountReq proto.InternalMessageInfo

func (m *DeleteAccountReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type DeleteAccountResp struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAccountResp) Reset()         { *m = DeleteAccountResp{} }
func (m *DeleteAccountResp) String() string { return proto.CompactTextString(m) }
func (*DeleteAccountResp) ProtoMessage()    {}
func (*DeleteAccountResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{1}
}
func (m *DeleteAccountResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAccountResp.Unmarshal(m, b)
}
func (m *DeleteAccountResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAccountResp.Marshal(b, m, deterministic)
}
func (m *DeleteAccountResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAccountResp.Merge(m, src)
}
func (m *DeleteAccountResp) XXX_Size() int {
	return xxx_messageInfo_DeleteAccountResp.Size(m)
}
func (m *DeleteAccountResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAccountResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAccountResp proto.InternalMessageInfo

func (m *DeleteAccountResp) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *DeleteAccountResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ChangeAccountPasswordReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PrevPass             string   `protobuf:"bytes,2,opt,name=prev_pass,json=prevPass,proto3" json:"prev_pass,omitempty"`
	NewPass              string   `protobuf:"bytes,3,opt,name=new_pass,json=newPass,proto3" json:"new_pass,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeAccountPasswordReq) Reset()         { *m = ChangeAccountPasswordReq{} }
func (m *ChangeAccountPasswordReq) String() string { return proto.CompactTextString(m) }
func (*ChangeAccountPasswordReq) ProtoMessage()    {}
func (*ChangeAccountPasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{2}
}
func (m *ChangeAccountPasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeAccountPasswordReq.Unmarshal(m, b)
}
func (m *ChangeAccountPasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeAccountPasswordReq.Marshal(b, m, deterministic)
}
func (m *ChangeAccountPasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeAccountPasswordReq.Merge(m, src)
}
func (m *ChangeAccountPasswordReq) XXX_Size() int {
	return xxx_messageInfo_ChangeAccountPasswordReq.Size(m)
}
func (m *ChangeAccountPasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeAccountPasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAccountPasswordReq proto.InternalMessageInfo

func (m *ChangeAccountPasswordReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ChangeAccountPasswordReq) GetPrevPass() string {
	if m != nil {
		return m.PrevPass
	}
	return ""
}

func (m *ChangeAccountPasswordReq) GetNewPass() string {
	if m != nil {
		return m.NewPass
	}
	return ""
}

type ChangeAccountPasswordResp struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeAccountPasswordResp) Reset()         { *m = ChangeAccountPasswordResp{} }
func (m *ChangeAccountPasswordResp) String() string { return proto.CompactTextString(m) }
func (*ChangeAccountPasswordResp) ProtoMessage()    {}
func (*ChangeAccountPasswordResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{3}
}
func (m *ChangeAccountPasswordResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeAccountPasswordResp.Unmarshal(m, b)
}
func (m *ChangeAccountPasswordResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeAccountPasswordResp.Marshal(b, m, deterministic)
}
func (m *ChangeAccountPasswordResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeAccountPasswordResp.Merge(m, src)
}
func (m *ChangeAccountPasswordResp) XXX_Size() int {
	return xxx_messageInfo_ChangeAccountPasswordResp.Size(m)
}
func (m *ChangeAccountPasswordResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeAccountPasswordResp.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAccountPasswordResp proto.InternalMessageInfo

func (m *ChangeAccountPasswordResp) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *ChangeAccountPasswordResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ChangeAccountEmailReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NewEmail             string   `protobuf:"bytes,2,opt,name=new_email,json=newEmail,proto3" json:"new_email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeAccountEmailReq) Reset()         { *m = ChangeAccountEmailReq{} }
func (m *ChangeAccountEmailReq) String() string { return proto.CompactTextString(m) }
func (*ChangeAccountEmailReq) ProtoMessage()    {}
func (*ChangeAccountEmailReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{4}
}
func (m *ChangeAccountEmailReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeAccountEmailReq.Unmarshal(m, b)
}
func (m *ChangeAccountEmailReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeAccountEmailReq.Marshal(b, m, deterministic)
}
func (m *ChangeAccountEmailReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeAccountEmailReq.Merge(m, src)
}
func (m *ChangeAccountEmailReq) XXX_Size() int {
	return xxx_messageInfo_ChangeAccountEmailReq.Size(m)
}
func (m *ChangeAccountEmailReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeAccountEmailReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAccountEmailReq proto.InternalMessageInfo

func (m *ChangeAccountEmailReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ChangeAccountEmailReq) GetNewEmail() string {
	if m != nil {
		return m.NewEmail
	}
	return ""
}

type ChangeAccountEmailResp struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeAccountEmailResp) Reset()         { *m = ChangeAccountEmailResp{} }
func (m *ChangeAccountEmailResp) String() string { return proto.CompactTextString(m) }
func (*ChangeAccountEmailResp) ProtoMessage()    {}
func (*ChangeAccountEmailResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{5}
}
func (m *ChangeAccountEmailResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeAccountEmailResp.Unmarshal(m, b)
}
func (m *ChangeAccountEmailResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeAccountEmailResp.Marshal(b, m, deterministic)
}
func (m *ChangeAccountEmailResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeAccountEmailResp.Merge(m, src)
}
func (m *ChangeAccountEmailResp) XXX_Size() int {
	return xxx_messageInfo_ChangeAccountEmailResp.Size(m)
}
func (m *ChangeAccountEmailResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeAccountEmailResp.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAccountEmailResp proto.InternalMessageInfo

func (m *ChangeAccountEmailResp) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *ChangeAccountEmailResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ChangeAccountInfoReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	UserName             string   `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Bio                  string   `protobuf:"bytes,5,opt,name=bio,proto3" json:"bio,omitempty"`
	Gender               string   `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthday             string   `protobuf:"bytes,7,opt,name=birthday,proto3" json:"birthday,omitempty"`
	CountryCode          string   `protobuf:"bytes,8,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeAccountInfoReq) Reset()         { *m = ChangeAccountInfoReq{} }
func (m *ChangeAccountInfoReq) String() string { return proto.CompactTextString(m) }
func (*ChangeAccountInfoReq) ProtoMessage()    {}
func (*ChangeAccountInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{6}
}
func (m *ChangeAccountInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeAccountInfoReq.Unmarshal(m, b)
}
func (m *ChangeAccountInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeAccountInfoReq.Marshal(b, m, deterministic)
}
func (m *ChangeAccountInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeAccountInfoReq.Merge(m, src)
}
func (m *ChangeAccountInfoReq) XXX_Size() int {
	return xxx_messageInfo_ChangeAccountInfoReq.Size(m)
}
func (m *ChangeAccountInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeAccountInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAccountInfoReq proto.InternalMessageInfo

func (m *ChangeAccountInfoReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ChangeAccountInfoReq) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *ChangeAccountInfoReq) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *ChangeAccountInfoReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *ChangeAccountInfoReq) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

func (m *ChangeAccountInfoReq) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *ChangeAccountInfoReq) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *ChangeAccountInfoReq) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

type ChangeAccountInfoResp struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeAccountInfoResp) Reset()         { *m = ChangeAccountInfoResp{} }
func (m *ChangeAccountInfoResp) String() string { return proto.CompactTextString(m) }
func (*ChangeAccountInfoResp) ProtoMessage()    {}
func (*ChangeAccountInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{7}
}
func (m *ChangeAccountInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeAccountInfoResp.Unmarshal(m, b)
}
func (m *ChangeAccountInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeAccountInfoResp.Marshal(b, m, deterministic)
}
func (m *ChangeAccountInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeAccountInfoResp.Merge(m, src)
}
func (m *ChangeAccountInfoResp) XXX_Size() int {
	return xxx_messageInfo_ChangeAccountInfoResp.Size(m)
}
func (m *ChangeAccountInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeAccountInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAccountInfoResp proto.InternalMessageInfo

func (m *ChangeAccountInfoResp) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *ChangeAccountInfoResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ForgotAccountPassReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NewPass              string   `protobuf:"bytes,2,opt,name=new_pass,json=newPass,proto3" json:"new_pass,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgotAccountPassReq) Reset()         { *m = ForgotAccountPassReq{} }
func (m *ForgotAccountPassReq) String() string { return proto.CompactTextString(m) }
func (*ForgotAccountPassReq) ProtoMessage()    {}
func (*ForgotAccountPassReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{8}
}
func (m *ForgotAccountPassReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgotAccountPassReq.Unmarshal(m, b)
}
func (m *ForgotAccountPassReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgotAccountPassReq.Marshal(b, m, deterministic)
}
func (m *ForgotAccountPassReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgotAccountPassReq.Merge(m, src)
}
func (m *ForgotAccountPassReq) XXX_Size() int {
	return xxx_messageInfo_ForgotAccountPassReq.Size(m)
}
func (m *ForgotAccountPassReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgotAccountPassReq.DiscardUnknown(m)
}

var xxx_messageInfo_ForgotAccountPassReq proto.InternalMessageInfo

func (m *ForgotAccountPassReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ForgotAccountPassReq) GetNewPass() string {
	if m != nil {
		return m.NewPass
	}
	return ""
}

type ForgotAccountPassResp struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgotAccountPassResp) Reset()         { *m = ForgotAccountPassResp{} }
func (m *ForgotAccountPassResp) String() string { return proto.CompactTextString(m) }
func (*ForgotAccountPassResp) ProtoMessage()    {}
func (*ForgotAccountPassResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{9}
}
func (m *ForgotAccountPassResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgotAccountPassResp.Unmarshal(m, b)
}
func (m *ForgotAccountPassResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgotAccountPassResp.Marshal(b, m, deterministic)
}
func (m *ForgotAccountPassResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgotAccountPassResp.Merge(m, src)
}
func (m *ForgotAccountPassResp) XXX_Size() int {
	return xxx_messageInfo_ForgotAccountPassResp.Size(m)
}
func (m *ForgotAccountPassResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgotAccountPassResp.DiscardUnknown(m)
}

var xxx_messageInfo_ForgotAccountPassResp proto.InternalMessageInfo

func (m *ForgotAccountPassResp) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *ForgotAccountPassResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteAccountReq)(nil), "accounts.DeleteAccountReq")
	proto.RegisterType((*DeleteAccountResp)(nil), "accounts.DeleteAccountResp")
	proto.RegisterType((*ChangeAccountPasswordReq)(nil), "accounts.ChangeAccountPasswordReq")
	proto.RegisterType((*ChangeAccountPasswordResp)(nil), "accounts.ChangeAccountPasswordResp")
	proto.RegisterType((*ChangeAccountEmailReq)(nil), "accounts.ChangeAccountEmailReq")
	proto.RegisterType((*ChangeAccountEmailResp)(nil), "accounts.ChangeAccountEmailResp")
	proto.RegisterType((*ChangeAccountInfoReq)(nil), "accounts.ChangeAccountInfoReq")
	proto.RegisterType((*ChangeAccountInfoResp)(nil), "accounts.ChangeAccountInfoResp")
	proto.RegisterType((*ForgotAccountPassReq)(nil), "accounts.ForgotAccountPassReq")
	proto.RegisterType((*ForgotAccountPassResp)(nil), "accounts.ForgotAccountPassResp")
}

func init() { proto.RegisterFile("accounts.proto", fileDescriptor_e1e7723af4c007b7) }

var fileDescriptor_e1e7723af4c007b7 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x13, 0x48, 0x9c, 0x29, 0x82, 0xb0, 0x4a, 0x52, 0xd7, 0x15, 0x34, 0x98, 0x4b, 0x25,
	0xa4, 0x1e, 0xe0, 0xc0, 0x99, 0x86, 0x22, 0x12, 0xa9, 0x28, 0x32, 0x27, 0x72, 0xb1, 0x36, 0xf6,
	0x24, 0xb5, 0x48, 0xbc, 0x66, 0xd7, 0x69, 0xd4, 0xaf, 0xe4, 0x5b, 0xf8, 0x03, 0xb4, 0x6b, 0xaf,
	0xeb, 0x58, 0xb1, 0x2b, 0xf9, 0xe6, 0x79, 0x6f, 0xe6, 0xed, 0xf3, 0xec, 0xcc, 0xc2, 0x4b, 0xea,
	0xfb, 0x6c, 0x17, 0x25, 0xe2, 0x2a, 0xe6, 0x2c, 0x61, 0xc4, 0xd4, 0xb1, 0xf3, 0x01, 0xfa, 0x5f,
	0x71, 0x83, 0x09, 0x7e, 0x49, 0x11, 0x17, 0xff, 0x90, 0x53, 0xe8, 0xee, 0x04, 0x72, 0x2f, 0x0c,
	0x2c, 0x63, 0x6c, 0x5c, 0xb6, 0xdd, 0x8e, 0x0c, 0xa7, 0x81, 0x73, 0x03, 0xaf, 0x4b, 0xc9, 0x22,
	0x26, 0x23, 0xe8, 0x88, 0x84, 0x26, 0x3b, 0xa1, 0x92, 0x4d, 0x37, 0x8b, 0x88, 0x05, 0xdd, 0x2d,
	0x0a, 0x41, 0xd7, 0x68, 0xb5, 0xc6, 0xc6, 0x65, 0xcf, 0xd5, 0xa1, 0xf3, 0x1b, 0xac, 0xc9, 0x1d,
	0x8d, 0xd6, 0x5a, 0x66, 0x4e, 0x85, 0xd8, 0x33, 0x1e, 0xd4, 0x9d, 0x4d, 0xce, 0xa1, 0x17, 0x73,
	0xbc, 0xf7, 0x62, 0x2a, 0x44, 0x26, 0x68, 0x4a, 0x40, 0x16, 0x93, 0x33, 0x30, 0x23, 0xdc, 0xa7,
	0x5c, 0x3b, 0x3d, 0x2c, 0xc2, 0xbd, 0xa4, 0x9c, 0x5b, 0x38, 0xab, 0x38, 0xac, 0x91, 0xf7, 0x5b,
	0x18, 0x1e, 0xc8, 0xdd, 0x6c, 0x69, 0xb8, 0x79, 0xca, 0xb8, 0xf4, 0x86, 0x32, 0x51, 0x1b, 0x8f,
	0x70, 0xaf, 0x0a, 0x9d, 0x19, 0x8c, 0x8e, 0xc9, 0x35, 0xb2, 0xf6, 0xcf, 0x80, 0xc1, 0x81, 0xd8,
	0x34, 0x5a, 0xb1, 0x5a, 0x6b, 0x6f, 0x00, 0x56, 0x21, 0x17, 0x89, 0x17, 0xd1, 0xad, 0x96, 0xeb,
	0x29, 0xe4, 0x07, 0xdd, 0xa2, 0x74, 0xbe, 0xa1, 0x9a, 0x4d, 0xdb, 0x6a, 0x4a, 0x40, 0x93, 0x4a,
	0x54, 0x91, 0xcf, 0x52, 0x52, 0x02, 0x8a, 0xec, 0x43, 0x7b, 0x19, 0x32, 0xeb, 0xb9, 0x82, 0xe5,
	0xa7, 0xfc, 0x9d, 0x35, 0x46, 0x01, 0x72, 0xab, 0xa3, 0xc0, 0x2c, 0x22, 0x36, 0x98, 0xcb, 0x90,
	0x27, 0x77, 0x01, 0x7d, 0xb0, 0xba, 0xa9, 0x8a, 0x8e, 0xc9, 0x3b, 0x78, 0xa1, 0xfe, 0x83, 0x3f,
	0x78, 0x3e, 0x0b, 0xd0, 0x32, 0x15, 0x7f, 0x92, 0x61, 0x13, 0x16, 0xa0, 0x33, 0x2d, 0x5d, 0x47,
	0xfa, 0xcb, 0x8d, 0xda, 0x37, 0x83, 0xc1, 0x37, 0xc6, 0xd7, 0x2c, 0x29, 0x0c, 0x4a, 0x6d, 0xf7,
	0x8a, 0x43, 0xd7, 0x3a, 0x1c, 0xba, 0x29, 0x0c, 0x8f, 0x68, 0x35, 0xb1, 0xf5, 0xf1, 0x6f, 0x1b,
	0x4e, 0xdc, 0xf9, 0x24, 0x13, 0x12, 0x64, 0x06, 0xaf, 0xf4, 0xf2, 0x7a, 0x81, 0x5a, 0x46, 0x62,
	0x5f, 0xe5, 0xeb, 0x5d, 0xde, 0x65, 0xfb, 0xbc, 0x92, 0x13, 0x31, 0x59, 0xc2, 0x69, 0xae, 0xe5,
	0xab, 0x36, 0xea, 0xed, 0x20, 0xce, 0x63, 0x5d, 0xd5, 0xae, 0xda, 0xef, 0x9f, 0xcc, 0x11, 0x31,
	0xf9, 0x05, 0x83, 0xd2, 0x19, 0x6a, 0xc6, 0xc9, 0x45, 0x45, 0xb1, 0x5e, 0x28, 0x7b, 0x5c, 0x9f,
	0xa0, 0xa4, 0x47, 0x65, 0xfb, 0x9c, 0xad, 0xc2, 0x0d, 0x92, 0xb7, 0x15, 0xb5, 0xd9, 0x46, 0xd8,
	0x17, 0xb5, 0xbc, 0x88, 0xc9, 0xa2, 0xd0, 0x99, 0x95, 0xba, 0xc9, 0xbc, 0x33, 0x05, 0xed, 0x63,
	0xf3, 0x52, 0xd4, 0x3e, 0x3a, 0x03, 0xd7, 0x9f, 0xa1, 0x9f, 0x6b, 0x0b, 0xe4, 0xf7, 0xa1, 0x8f,
	0xd7, 0xc3, 0xc5, 0x5c, 0x3e, 0xcc, 0xfa, 0x96, 0x7f, 0xa6, 0xf0, 0xf7, 0xd6, 0xdc, 0x58, 0xe4,
	0x6f, 0xf5, 0xb2, 0xa3, 0x1e, 0xef, 0x4f, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xd7, 0x78,
	0xed, 0xce, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCAccountsClient is the client API for RPCAccounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCAccountsClient interface {
	AccountsDelete(ctx context.Context, in *DeleteAccountReq, opts ...grpc.CallOption) (*DeleteAccountResp, error)
	AccountsChangePassword(ctx context.Context, in *ChangeAccountPasswordReq, opts ...grpc.CallOption) (*ChangeAccountPasswordResp, error)
	AccountsChangeEmail(ctx context.Context, in *ChangeAccountEmailReq, opts ...grpc.CallOption) (*ChangeAccountEmailResp, error)
	AccountsChangeProfile(ctx context.Context, in *ChangeAccountInfoReq, opts ...grpc.CallOption) (*ChangeAccountInfoResp, error)
	AccountsForgotPassword(ctx context.Context, in *ForgotAccountPassReq, opts ...grpc.CallOption) (*ForgotAccountPassResp, error)
}

type rPCAccountsClient struct {
	cc *grpc.ClientConn
}

func NewRPCAccountsClient(cc *grpc.ClientConn) RPCAccountsClient {
	return &rPCAccountsClient{cc}
}

func (c *rPCAccountsClient) AccountsDelete(ctx context.Context, in *DeleteAccountReq, opts ...grpc.CallOption) (*DeleteAccountResp, error) {
	out := new(DeleteAccountResp)
	err := c.cc.Invoke(ctx, "/accounts.RPCAccounts/accounts_delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAccountsClient) AccountsChangePassword(ctx context.Context, in *ChangeAccountPasswordReq, opts ...grpc.CallOption) (*ChangeAccountPasswordResp, error) {
	out := new(ChangeAccountPasswordResp)
	err := c.cc.Invoke(ctx, "/accounts.RPCAccounts/accounts_changePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAccountsClient) AccountsChangeEmail(ctx context.Context, in *ChangeAccountEmailReq, opts ...grpc.CallOption) (*ChangeAccountEmailResp, error) {
	out := new(ChangeAccountEmailResp)
	err := c.cc.Invoke(ctx, "/accounts.RPCAccounts/accounts_changeEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAccountsClient) AccountsChangeProfile(ctx context.Context, in *ChangeAccountInfoReq, opts ...grpc.CallOption) (*ChangeAccountInfoResp, error) {
	out := new(ChangeAccountInfoResp)
	err := c.cc.Invoke(ctx, "/accounts.RPCAccounts/accounts_changeProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAccountsClient) AccountsForgotPassword(ctx context.Context, in *ForgotAccountPassReq, opts ...grpc.CallOption) (*ForgotAccountPassResp, error) {
	out := new(ForgotAccountPassResp)
	err := c.cc.Invoke(ctx, "/accounts.RPCAccounts/accounts_forgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCAccountsServer is the server API for RPCAccounts service.
type RPCAccountsServer interface {
	AccountsDelete(context.Context, *DeleteAccountReq) (*DeleteAccountResp, error)
	AccountsChangePassword(context.Context, *ChangeAccountPasswordReq) (*ChangeAccountPasswordResp, error)
	AccountsChangeEmail(context.Context, *ChangeAccountEmailReq) (*ChangeAccountEmailResp, error)
	AccountsChangeProfile(context.Context, *ChangeAccountInfoReq) (*ChangeAccountInfoResp, error)
	AccountsForgotPassword(context.Context, *ForgotAccountPassReq) (*ForgotAccountPassResp, error)
}

// UnimplementedRPCAccountsServer can be embedded to have forward compatible implementations.
type UnimplementedRPCAccountsServer struct {
}

func (*UnimplementedRPCAccountsServer) AccountsDelete(ctx context.Context, req *DeleteAccountReq) (*DeleteAccountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountsDelete not implemented")
}
func (*UnimplementedRPCAccountsServer) AccountsChangePassword(ctx context.Context, req *ChangeAccountPasswordReq) (*ChangeAccountPasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountsChangePassword not implemented")
}
func (*UnimplementedRPCAccountsServer) AccountsChangeEmail(ctx context.Context, req *ChangeAccountEmailReq) (*ChangeAccountEmailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountsChangeEmail not implemented")
}
func (*UnimplementedRPCAccountsServer) AccountsChangeProfile(ctx context.Context, req *ChangeAccountInfoReq) (*ChangeAccountInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountsChangeProfile not implemented")
}
func (*UnimplementedRPCAccountsServer) AccountsForgotPassword(ctx context.Context, req *ForgotAccountPassReq) (*ForgotAccountPassResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountsForgotPassword not implemented")
}

func RegisterRPCAccountsServer(s *grpc.Server, srv RPCAccountsServer) {
	s.RegisterService(&_RPCAccounts_serviceDesc, srv)
}

func _RPCAccounts_AccountsDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAccountsServer).AccountsDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.RPCAccounts/AccountsDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAccountsServer).AccountsDelete(ctx, req.(*DeleteAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAccounts_AccountsChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeAccountPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAccountsServer).AccountsChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.RPCAccounts/AccountsChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAccountsServer).AccountsChangePassword(ctx, req.(*ChangeAccountPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAccounts_AccountsChangeEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeAccountEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAccountsServer).AccountsChangeEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.RPCAccounts/AccountsChangeEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAccountsServer).AccountsChangeEmail(ctx, req.(*ChangeAccountEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAccounts_AccountsChangeProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeAccountInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAccountsServer).AccountsChangeProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.RPCAccounts/AccountsChangeProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAccountsServer).AccountsChangeProfile(ctx, req.(*ChangeAccountInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAccounts_AccountsForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotAccountPassReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAccountsServer).AccountsForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.RPCAccounts/AccountsForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAccountsServer).AccountsForgotPassword(ctx, req.(*ForgotAccountPassReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCAccounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "accounts.RPCAccounts",
	HandlerType: (*RPCAccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "accounts_delete",
			Handler:    _RPCAccounts_AccountsDelete_Handler,
		},
		{
			MethodName: "accounts_changePassword",
			Handler:    _RPCAccounts_AccountsChangePassword_Handler,
		},
		{
			MethodName: "accounts_changeEmail",
			Handler:    _RPCAccounts_AccountsChangeEmail_Handler,
		},
		{
			MethodName: "accounts_changeProfile",
			Handler:    _RPCAccounts_AccountsChangeProfile_Handler,
		},
		{
			MethodName: "accounts_forgotPassword",
			Handler:    _RPCAccounts_AccountsForgotPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts.proto",
}
