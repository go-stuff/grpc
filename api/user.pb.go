// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
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

type User struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Groups               []string             `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
	CreatedBy            string               `protobuf:"bytes,4,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ModifiedBy           string               `protobuf:"bytes,6,opt,name=modifiedBy,proto3" json:"modifiedBy,omitempty"`
	ModifiedAt           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=modifiedAt,proto3" json:"modifiedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *User) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetModifiedBy() string {
	if m != nil {
		return m.ModifiedBy
	}
	return ""
}

func (m *User) GetModifiedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ModifiedAt
	}
	return nil
}

type UserSliceRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSliceRequest) Reset()         { *m = UserSliceRequest{} }
func (m *UserSliceRequest) String() string { return proto.CompactTextString(m) }
func (*UserSliceRequest) ProtoMessage()    {}
func (*UserSliceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserSliceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSliceRequest.Unmarshal(m, b)
}
func (m *UserSliceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSliceRequest.Marshal(b, m, deterministic)
}
func (m *UserSliceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSliceRequest.Merge(m, src)
}
func (m *UserSliceRequest) XXX_Size() int {
	return xxx_messageInfo_UserSliceRequest.Size(m)
}
func (m *UserSliceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSliceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserSliceRequest proto.InternalMessageInfo

type UserSliceResponse struct {
	// return a collection of users
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSliceResponse) Reset()         { *m = UserSliceResponse{} }
func (m *UserSliceResponse) String() string { return proto.CompactTextString(m) }
func (*UserSliceResponse) ProtoMessage()    {}
func (*UserSliceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserSliceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSliceResponse.Unmarshal(m, b)
}
func (m *UserSliceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSliceResponse.Marshal(b, m, deterministic)
}
func (m *UserSliceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSliceResponse.Merge(m, src)
}
func (m *UserSliceResponse) XXX_Size() int {
	return xxx_messageInfo_UserSliceResponse.Size(m)
}
func (m *UserSliceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSliceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserSliceResponse proto.InternalMessageInfo

func (m *UserSliceResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UserCreateRequest struct {
	// create a user
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCreateRequest) Reset()         { *m = UserCreateRequest{} }
func (m *UserCreateRequest) String() string { return proto.CompactTextString(m) }
func (*UserCreateRequest) ProtoMessage()    {}
func (*UserCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreateRequest.Unmarshal(m, b)
}
func (m *UserCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreateRequest.Marshal(b, m, deterministic)
}
func (m *UserCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreateRequest.Merge(m, src)
}
func (m *UserCreateRequest) XXX_Size() int {
	return xxx_messageInfo_UserCreateRequest.Size(m)
}
func (m *UserCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreateRequest proto.InternalMessageInfo

func (m *UserCreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserCreateResponse struct {
	// reutrn an id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCreateResponse) Reset()         { *m = UserCreateResponse{} }
func (m *UserCreateResponse) String() string { return proto.CompactTextString(m) }
func (*UserCreateResponse) ProtoMessage()    {}
func (*UserCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreateResponse.Unmarshal(m, b)
}
func (m *UserCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreateResponse.Marshal(b, m, deterministic)
}
func (m *UserCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreateResponse.Merge(m, src)
}
func (m *UserCreateResponse) XXX_Size() int {
	return xxx_messageInfo_UserCreateResponse.Size(m)
}
func (m *UserCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreateResponse proto.InternalMessageInfo

func (m *UserCreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UserReadRequest struct {
	// read a user by id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserReadRequest) Reset()         { *m = UserReadRequest{} }
func (m *UserReadRequest) String() string { return proto.CompactTextString(m) }
func (*UserReadRequest) ProtoMessage()    {}
func (*UserReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *UserReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReadRequest.Unmarshal(m, b)
}
func (m *UserReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReadRequest.Marshal(b, m, deterministic)
}
func (m *UserReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReadRequest.Merge(m, src)
}
func (m *UserReadRequest) XXX_Size() int {
	return xxx_messageInfo_UserReadRequest.Size(m)
}
func (m *UserReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserReadRequest proto.InternalMessageInfo

func (m *UserReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UserReadResponse struct {
	// return a user
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserReadResponse) Reset()         { *m = UserReadResponse{} }
func (m *UserReadResponse) String() string { return proto.CompactTextString(m) }
func (*UserReadResponse) ProtoMessage()    {}
func (*UserReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UserReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReadResponse.Unmarshal(m, b)
}
func (m *UserReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReadResponse.Marshal(b, m, deterministic)
}
func (m *UserReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReadResponse.Merge(m, src)
}
func (m *UserReadResponse) XXX_Size() int {
	return xxx_messageInfo_UserReadResponse.Size(m)
}
func (m *UserReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserReadResponse proto.InternalMessageInfo

func (m *UserReadResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserUpdateRequest struct {
	// update a user by id
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserUpdateRequest) Reset()         { *m = UserUpdateRequest{} }
func (m *UserUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UserUpdateRequest) ProtoMessage()    {}
func (*UserUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UserUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserUpdateRequest.Unmarshal(m, b)
}
func (m *UserUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserUpdateRequest.Marshal(b, m, deterministic)
}
func (m *UserUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdateRequest.Merge(m, src)
}
func (m *UserUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UserUpdateRequest.Size(m)
}
func (m *UserUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdateRequest proto.InternalMessageInfo

func (m *UserUpdateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserUpdateResponse struct {
	// return number of users updated
	// should equal 1 on success
	Updated              int64    `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserUpdateResponse) Reset()         { *m = UserUpdateResponse{} }
func (m *UserUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UserUpdateResponse) ProtoMessage()    {}
func (*UserUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *UserUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserUpdateResponse.Unmarshal(m, b)
}
func (m *UserUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserUpdateResponse.Marshal(b, m, deterministic)
}
func (m *UserUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdateResponse.Merge(m, src)
}
func (m *UserUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UserUpdateResponse.Size(m)
}
func (m *UserUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdateResponse proto.InternalMessageInfo

func (m *UserUpdateResponse) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type UserDeleteRequest struct {
	// delete a user by id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDeleteRequest) Reset()         { *m = UserDeleteRequest{} }
func (m *UserDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*UserDeleteRequest) ProtoMessage()    {}
func (*UserDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *UserDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeleteRequest.Unmarshal(m, b)
}
func (m *UserDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeleteRequest.Marshal(b, m, deterministic)
}
func (m *UserDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeleteRequest.Merge(m, src)
}
func (m *UserDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_UserDeleteRequest.Size(m)
}
func (m *UserDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeleteRequest proto.InternalMessageInfo

func (m *UserDeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UserDeleteResponse struct {
	// return number of users deleted
	// should equal 1 on success
	Deleted              int64    `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDeleteResponse) Reset()         { *m = UserDeleteResponse{} }
func (m *UserDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*UserDeleteResponse) ProtoMessage()    {}
func (*UserDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}

func (m *UserDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeleteResponse.Unmarshal(m, b)
}
func (m *UserDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeleteResponse.Marshal(b, m, deterministic)
}
func (m *UserDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeleteResponse.Merge(m, src)
}
func (m *UserDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_UserDeleteResponse.Size(m)
}
func (m *UserDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeleteResponse proto.InternalMessageInfo

func (m *UserDeleteResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "api.User")
	proto.RegisterType((*UserSliceRequest)(nil), "api.UserSliceRequest")
	proto.RegisterType((*UserSliceResponse)(nil), "api.UserSliceResponse")
	proto.RegisterType((*UserCreateRequest)(nil), "api.UserCreateRequest")
	proto.RegisterType((*UserCreateResponse)(nil), "api.UserCreateResponse")
	proto.RegisterType((*UserReadRequest)(nil), "api.UserReadRequest")
	proto.RegisterType((*UserReadResponse)(nil), "api.UserReadResponse")
	proto.RegisterType((*UserUpdateRequest)(nil), "api.UserUpdateRequest")
	proto.RegisterType((*UserUpdateResponse)(nil), "api.UserUpdateResponse")
	proto.RegisterType((*UserDeleteRequest)(nil), "api.UserDeleteRequest")
	proto.RegisterType((*UserDeleteResponse)(nil), "api.UserDeleteResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdd, 0x6e, 0x94, 0x40,
	0x14, 0x0e, 0x3f, 0x4b, 0xe5, 0x6c, 0xe2, 0xcf, 0x89, 0xad, 0x13, 0xa2, 0x16, 0xd1, 0x0b, 0xae,
	0x68, 0xa4, 0xc6, 0x58, 0xef, 0xaa, 0x3e, 0x01, 0xda, 0x07, 0xa0, 0xcb, 0xe9, 0x66, 0x92, 0xa5,
	0x20, 0x03, 0x26, 0x3e, 0x93, 0xef, 0x68, 0x0c, 0x33, 0x03, 0x33, 0xe0, 0xc5, 0xf6, 0xf2, 0x7c,
	0x73, 0xbe, 0x1f, 0xbe, 0x13, 0x00, 0x06, 0x41, 0x5d, 0xd6, 0x76, 0x4d, 0xdf, 0xa0, 0x57, 0xb6,
	0x3c, 0x3a, 0xdf, 0x37, 0xcd, 0xfe, 0x40, 0x17, 0x12, 0xba, 0x1d, 0xee, 0x2e, 0x7a, 0x5e, 0x93,
	0xe8, 0xcb, 0xba, 0x55, 0x5b, 0xc9, 0x5f, 0x07, 0xfc, 0x1b, 0x41, 0x1d, 0x3e, 0x06, 0x97, 0x57,
	0xcc, 0x89, 0x9d, 0x34, 0x2c, 0x5c, 0x5e, 0x61, 0x04, 0x8f, 0x46, 0xb1, 0xfb, 0xb2, 0x26, 0xe6,
	0x4a, 0x74, 0x9e, 0xf1, 0x0c, 0x82, 0x7d, 0xd7, 0x0c, 0xad, 0x60, 0x5e, 0xec, 0xa5, 0x61, 0xa1,
	0x27, 0x7c, 0x09, 0xe1, 0xae, 0xa3, 0xb2, 0xa7, 0xea, 0xcb, 0x6f, 0xe6, 0x4b, 0x92, 0x01, 0xf0,
	0xd3, 0xfc, 0x7a, 0xdd, 0xb3, 0x4d, 0xec, 0xa4, 0xdb, 0x3c, 0xca, 0x54, 0xbe, 0x6c, 0xca, 0x97,
	0xfd, 0x98, 0xf2, 0x15, 0x66, 0x19, 0x5f, 0x03, 0xd4, 0x4d, 0xc5, 0xef, 0xb8, 0x14, 0x0e, 0xa4,
	0xb0, 0x85, 0xe0, 0x67, 0xf3, 0x7e, 0xdd, 0xb3, 0x93, 0xa3, 0xd2, 0xd6, 0x76, 0x82, 0xf0, 0x74,
	0xfc, 0xfe, 0xef, 0x07, 0xbe, 0xa3, 0x82, 0x7e, 0x0e, 0x24, 0xfa, 0xe4, 0x03, 0x3c, 0xb3, 0x30,
	0xd1, 0x36, 0xf7, 0x82, 0xf0, 0x1c, 0x36, 0x63, 0x01, 0x82, 0x39, 0xb1, 0x97, 0x6e, 0xf3, 0x30,
	0x2b, 0x5b, 0x9e, 0x8d, 0x6b, 0x85, 0xc2, 0x93, 0x5c, 0xb1, 0xbe, 0xca, 0xd8, 0x5a, 0x0a, 0x5f,
	0x81, 0x3f, 0xbe, 0xca, 0x62, 0x17, 0x24, 0x09, 0x27, 0xef, 0x00, 0x6d, 0x8e, 0xb6, 0x5a, 0xdd,
	0x22, 0x79, 0x03, 0x4f, 0x24, 0x87, 0xca, 0x6a, 0xd2, 0x5d, 0xaf, 0xbc, 0x57, 0x9f, 0xa1, 0x56,
	0xb4, 0xcc, 0x11, 0x6f, 0x9d, 0xf7, 0xa6, 0xad, 0x1e, 0x9e, 0x37, 0x53, 0x79, 0x27, 0x8e, 0x36,
	0x62, 0x70, 0x32, 0x48, 0x44, 0x25, 0xf2, 0x8a, 0x69, 0x4c, 0xde, 0x2a, 0x8f, 0x6f, 0x74, 0x20,
	0xe3, 0xb1, 0xce, 0xae, 0x45, 0xa7, 0x25, 0x23, 0x5a, 0x49, 0x64, 0x16, 0xd5, 0x63, 0xfe, 0xc7,
	0x85, 0xad, 0xbc, 0x0f, 0x75, 0xbf, 0xf8, 0x8e, 0xf0, 0x23, 0x6c, 0xe4, 0xa9, 0xf0, 0x74, 0x8e,
	0x6b, 0x9f, 0x33, 0x3a, 0x5b, 0xc3, 0xda, 0xe1, 0x0a, 0x02, 0x55, 0x3c, 0x9a, 0x8d, 0xc5, 0xf5,
	0xa2, 0x17, 0xff, 0xe1, 0x9a, 0x7a, 0x09, 0xfe, 0x58, 0x35, 0x3e, 0x37, 0x05, 0x99, 0xe3, 0x44,
	0xa7, 0x2b, 0xd4, 0xf8, 0xa9, 0xe2, 0x2c, 0xbf, 0x45, 0xfb, 0x96, 0xdf, 0xaa, 0xe1, 0x2b, 0x08,
	0x54, 0x3d, 0x16, 0x75, 0x51, 0xaa, 0x45, 0x5d, 0xf6, 0x78, 0x1b, 0xc8, 0x1f, 0xe0, 0xf2, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xab, 0x06, 0x31, 0x1c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Slice(ctx context.Context, in *UserSliceRequest, opts ...grpc.CallOption) (*UserSliceResponse, error)
	Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	Read(ctx context.Context, in *UserReadRequest, opts ...grpc.CallOption) (*UserReadResponse, error)
	Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserUpdateResponse, error)
	Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Slice(ctx context.Context, in *UserSliceRequest, opts ...grpc.CallOption) (*UserSliceResponse, error) {
	out := new(UserSliceResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/Slice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Read(ctx context.Context, in *UserReadRequest, opts ...grpc.CallOption) (*UserReadResponse, error) {
	out := new(UserReadResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserUpdateResponse, error) {
	out := new(UserUpdateResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error) {
	out := new(UserDeleteResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Slice(context.Context, *UserSliceRequest) (*UserSliceResponse, error)
	Create(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	Read(context.Context, *UserReadRequest) (*UserReadResponse, error)
	Update(context.Context, *UserUpdateRequest) (*UserUpdateResponse, error)
	Delete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Slice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSliceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Slice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Slice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Slice(ctx, req.(*UserSliceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Read(ctx, req.(*UserReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Slice",
			Handler:    _UserService_Slice_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _UserService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
