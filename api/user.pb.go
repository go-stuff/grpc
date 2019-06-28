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
	// @inject_tag: bson:"_id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" bson:"_id"`
	// @inject_tag: bson:"username"
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty" bson:"username"`
	// @inject_tag: bson:"groups"
	Groups []string `protobuf:"bytes,3,rep,name=Groups,proto3" json:"Groups,omitempty" bson:"groups"`
	// @inject_tag: bson:"roleid"
	RoleID string `protobuf:"bytes,4,opt,name=RoleID,proto3" json:"RoleID,omitempty" bson:"roleid"`
	// @inject_tag: bson:"createdby"
	CreatedBy string `protobuf:"bytes,5,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty" bson:"createdby"`
	// @inject_tag: bson:"createdat"
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" bson:"createdat"`
	// @inject_tag: bson:"modifiedby"
	ModifiedBy string `protobuf:"bytes,7,opt,name=ModifiedBy,proto3" json:"ModifiedBy,omitempty" bson:"modifiedby"`
	// @inject_tag: bson:"modifiedat"
	ModifiedAt           *timestamp.Timestamp `protobuf:"bytes,8,opt,name=ModifiedAt,proto3" json:"ModifiedAt,omitempty" bson:"modifiedat"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" bson:"-"`
	XXX_unrecognized     []byte               `json:"-" bson:"-"`
	XXX_sizecache        int32                `json:"-" bson:"-"`
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

func (m *User) GetID() string {
	if m != nil {
		return m.ID
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

func (m *User) GetRoleID() string {
	if m != nil {
		return m.RoleID
	}
	return ""
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

type UserListReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserListReq) Reset()         { *m = UserListReq{} }
func (m *UserListReq) String() string { return proto.CompactTextString(m) }
func (*UserListReq) ProtoMessage()    {}
func (*UserListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListReq.Unmarshal(m, b)
}
func (m *UserListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListReq.Marshal(b, m, deterministic)
}
func (m *UserListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListReq.Merge(m, src)
}
func (m *UserListReq) XXX_Size() int {
	return xxx_messageInfo_UserListReq.Size(m)
}
func (m *UserListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserListReq proto.InternalMessageInfo

type UserListRes struct {
	// return a collection of users
	Users                []*User  `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserListRes) Reset()         { *m = UserListRes{} }
func (m *UserListRes) String() string { return proto.CompactTextString(m) }
func (*UserListRes) ProtoMessage()    {}
func (*UserListRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserListRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListRes.Unmarshal(m, b)
}
func (m *UserListRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListRes.Marshal(b, m, deterministic)
}
func (m *UserListRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListRes.Merge(m, src)
}
func (m *UserListRes) XXX_Size() int {
	return xxx_messageInfo_UserListRes.Size(m)
}
func (m *UserListRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListRes.DiscardUnknown(m)
}

var xxx_messageInfo_UserListRes proto.InternalMessageInfo

func (m *UserListRes) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UserCreateReq struct {
	// create a user
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserCreateReq) Reset()         { *m = UserCreateReq{} }
func (m *UserCreateReq) String() string { return proto.CompactTextString(m) }
func (*UserCreateReq) ProtoMessage()    {}
func (*UserCreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserCreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreateReq.Unmarshal(m, b)
}
func (m *UserCreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreateReq.Marshal(b, m, deterministic)
}
func (m *UserCreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreateReq.Merge(m, src)
}
func (m *UserCreateReq) XXX_Size() int {
	return xxx_messageInfo_UserCreateReq.Size(m)
}
func (m *UserCreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreateReq proto.InternalMessageInfo

func (m *UserCreateReq) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserCreateRes struct {
	// reutrn an id
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserCreateRes) Reset()         { *m = UserCreateRes{} }
func (m *UserCreateRes) String() string { return proto.CompactTextString(m) }
func (*UserCreateRes) ProtoMessage()    {}
func (*UserCreateRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserCreateRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreateRes.Unmarshal(m, b)
}
func (m *UserCreateRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreateRes.Marshal(b, m, deterministic)
}
func (m *UserCreateRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreateRes.Merge(m, src)
}
func (m *UserCreateRes) XXX_Size() int {
	return xxx_messageInfo_UserCreateRes.Size(m)
}
func (m *UserCreateRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreateRes.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreateRes proto.InternalMessageInfo

func (m *UserCreateRes) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type UserReadReq struct {
	// read a user by id
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserReadReq) Reset()         { *m = UserReadReq{} }
func (m *UserReadReq) String() string { return proto.CompactTextString(m) }
func (*UserReadReq) ProtoMessage()    {}
func (*UserReadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *UserReadReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReadReq.Unmarshal(m, b)
}
func (m *UserReadReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReadReq.Marshal(b, m, deterministic)
}
func (m *UserReadReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReadReq.Merge(m, src)
}
func (m *UserReadReq) XXX_Size() int {
	return xxx_messageInfo_UserReadReq.Size(m)
}
func (m *UserReadReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReadReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserReadReq proto.InternalMessageInfo

func (m *UserReadReq) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type UserReadRes struct {
	// return a user
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserReadRes) Reset()         { *m = UserReadRes{} }
func (m *UserReadRes) String() string { return proto.CompactTextString(m) }
func (*UserReadRes) ProtoMessage()    {}
func (*UserReadRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UserReadRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReadRes.Unmarshal(m, b)
}
func (m *UserReadRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReadRes.Marshal(b, m, deterministic)
}
func (m *UserReadRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReadRes.Merge(m, src)
}
func (m *UserReadRes) XXX_Size() int {
	return xxx_messageInfo_UserReadRes.Size(m)
}
func (m *UserReadRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReadRes.DiscardUnknown(m)
}

var xxx_messageInfo_UserReadRes proto.InternalMessageInfo

func (m *UserReadRes) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserReadByUsernameReq struct {
	// read a user by id
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserReadByUsernameReq) Reset()         { *m = UserReadByUsernameReq{} }
func (m *UserReadByUsernameReq) String() string { return proto.CompactTextString(m) }
func (*UserReadByUsernameReq) ProtoMessage()    {}
func (*UserReadByUsernameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UserReadByUsernameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReadByUsernameReq.Unmarshal(m, b)
}
func (m *UserReadByUsernameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReadByUsernameReq.Marshal(b, m, deterministic)
}
func (m *UserReadByUsernameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReadByUsernameReq.Merge(m, src)
}
func (m *UserReadByUsernameReq) XXX_Size() int {
	return xxx_messageInfo_UserReadByUsernameReq.Size(m)
}
func (m *UserReadByUsernameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReadByUsernameReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserReadByUsernameReq proto.InternalMessageInfo

func (m *UserReadByUsernameReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserReadByUsernameRes struct {
	// return a user
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserReadByUsernameRes) Reset()         { *m = UserReadByUsernameRes{} }
func (m *UserReadByUsernameRes) String() string { return proto.CompactTextString(m) }
func (*UserReadByUsernameRes) ProtoMessage()    {}
func (*UserReadByUsernameRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *UserReadByUsernameRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReadByUsernameRes.Unmarshal(m, b)
}
func (m *UserReadByUsernameRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReadByUsernameRes.Marshal(b, m, deterministic)
}
func (m *UserReadByUsernameRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReadByUsernameRes.Merge(m, src)
}
func (m *UserReadByUsernameRes) XXX_Size() int {
	return xxx_messageInfo_UserReadByUsernameRes.Size(m)
}
func (m *UserReadByUsernameRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReadByUsernameRes.DiscardUnknown(m)
}

var xxx_messageInfo_UserReadByUsernameRes proto.InternalMessageInfo

func (m *UserReadByUsernameRes) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserUpdateReq struct {
	// update a user by id
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserUpdateReq) Reset()         { *m = UserUpdateReq{} }
func (m *UserUpdateReq) String() string { return proto.CompactTextString(m) }
func (*UserUpdateReq) ProtoMessage()    {}
func (*UserUpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *UserUpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserUpdateReq.Unmarshal(m, b)
}
func (m *UserUpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserUpdateReq.Marshal(b, m, deterministic)
}
func (m *UserUpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdateReq.Merge(m, src)
}
func (m *UserUpdateReq) XXX_Size() int {
	return xxx_messageInfo_UserUpdateReq.Size(m)
}
func (m *UserUpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdateReq proto.InternalMessageInfo

func (m *UserUpdateReq) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserUpdateRes struct {
	// return number of users updated
	// should equal 1 on success
	Updated              int64    `protobuf:"varint,1,opt,name=Updated,proto3" json:"Updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserUpdateRes) Reset()         { *m = UserUpdateRes{} }
func (m *UserUpdateRes) String() string { return proto.CompactTextString(m) }
func (*UserUpdateRes) ProtoMessage()    {}
func (*UserUpdateRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}

func (m *UserUpdateRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserUpdateRes.Unmarshal(m, b)
}
func (m *UserUpdateRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserUpdateRes.Marshal(b, m, deterministic)
}
func (m *UserUpdateRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdateRes.Merge(m, src)
}
func (m *UserUpdateRes) XXX_Size() int {
	return xxx_messageInfo_UserUpdateRes.Size(m)
}
func (m *UserUpdateRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdateRes.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdateRes proto.InternalMessageInfo

func (m *UserUpdateRes) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type UserDeleteReq struct {
	// delete a user by id
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserDeleteReq) Reset()         { *m = UserDeleteReq{} }
func (m *UserDeleteReq) String() string { return proto.CompactTextString(m) }
func (*UserDeleteReq) ProtoMessage()    {}
func (*UserDeleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{11}
}

func (m *UserDeleteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeleteReq.Unmarshal(m, b)
}
func (m *UserDeleteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeleteReq.Marshal(b, m, deterministic)
}
func (m *UserDeleteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeleteReq.Merge(m, src)
}
func (m *UserDeleteReq) XXX_Size() int {
	return xxx_messageInfo_UserDeleteReq.Size(m)
}
func (m *UserDeleteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeleteReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeleteReq proto.InternalMessageInfo

func (m *UserDeleteReq) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type UserDeleteRes struct {
	// return number of users deleted
	// should equal 1 on success
	Deleted              int64    `protobuf:"varint,1,opt,name=Deleted,proto3" json:"Deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserDeleteRes) Reset()         { *m = UserDeleteRes{} }
func (m *UserDeleteRes) String() string { return proto.CompactTextString(m) }
func (*UserDeleteRes) ProtoMessage()    {}
func (*UserDeleteRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{12}
}

func (m *UserDeleteRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeleteRes.Unmarshal(m, b)
}
func (m *UserDeleteRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeleteRes.Marshal(b, m, deterministic)
}
func (m *UserDeleteRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeleteRes.Merge(m, src)
}
func (m *UserDeleteRes) XXX_Size() int {
	return xxx_messageInfo_UserDeleteRes.Size(m)
}
func (m *UserDeleteRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeleteRes.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeleteRes proto.InternalMessageInfo

func (m *UserDeleteRes) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "api.User")
	proto.RegisterType((*UserListReq)(nil), "api.UserListReq")
	proto.RegisterType((*UserListRes)(nil), "api.UserListRes")
	proto.RegisterType((*UserCreateReq)(nil), "api.UserCreateReq")
	proto.RegisterType((*UserCreateRes)(nil), "api.UserCreateRes")
	proto.RegisterType((*UserReadReq)(nil), "api.UserReadReq")
	proto.RegisterType((*UserReadRes)(nil), "api.UserReadRes")
	proto.RegisterType((*UserReadByUsernameReq)(nil), "api.UserReadByUsernameReq")
	proto.RegisterType((*UserReadByUsernameRes)(nil), "api.UserReadByUsernameRes")
	proto.RegisterType((*UserUpdateReq)(nil), "api.UserUpdateReq")
	proto.RegisterType((*UserUpdateRes)(nil), "api.UserUpdateRes")
	proto.RegisterType((*UserDeleteReq)(nil), "api.UserDeleteReq")
	proto.RegisterType((*UserDeleteRes)(nil), "api.UserDeleteRes")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0xc5, 0x76, 0x9a, 0xd6, 0x37, 0xb4, 0x8c, 0x0b, 0x1b, 0x42, 0xac, 0x8b, 0xf1, 0x93, 0x3b,
	0x86, 0x3a, 0x52, 0x18, 0x63, 0x6f, 0xed, 0x0c, 0x5b, 0x60, 0x7b, 0xd1, 0xd6, 0x0f, 0x70, 0x67,
	0xb5, 0x18, 0x92, 0xd9, 0xb5, 0x9c, 0x41, 0x7e, 0x64, 0x7f, 0xb4, 0xff, 0x1a, 0x92, 0x6c, 0x45,
	0x72, 0x36, 0xd2, 0xb7, 0x9c, 0x73, 0xef, 0x3d, 0x3a, 0x39, 0x27, 0x01, 0xd8, 0x48, 0xd1, 0xb2,
	0xa6, 0xad, 0xbb, 0x1a, 0xa3, 0xa2, 0xa9, 0xe8, 0xfc, 0xa1, 0xae, 0x1f, 0x56, 0xe2, 0x52, 0x53,
	0x77, 0x9b, 0xfb, 0xcb, 0xae, 0x5a, 0x0b, 0xd9, 0x15, 0xeb, 0xc6, 0x6c, 0xa5, 0xbf, 0x43, 0x98,
	0xdc, 0x4a, 0xd1, 0xe2, 0x19, 0x84, 0xcb, 0x9c, 0x04, 0x49, 0x90, 0xc5, 0x3c, 0x5c, 0xe6, 0x48,
	0xe1, 0x44, 0xf1, 0x3f, 0x8b, 0xb5, 0x20, 0xa1, 0x66, 0x2d, 0xc6, 0x17, 0x30, 0xfd, 0xd4, 0xd6,
	0x9b, 0x46, 0x92, 0x28, 0x89, 0xb2, 0x98, 0xf7, 0x48, 0xf1, 0xbc, 0x5e, 0x89, 0x65, 0x4e, 0x26,
	0xfa, 0xa2, 0x47, 0xf8, 0x12, 0xe2, 0x8f, 0xad, 0x28, 0x3a, 0x51, 0xde, 0x6c, 0xc9, 0x91, 0x1e,
	0xed, 0x08, 0x7c, 0x6f, 0xa7, 0xd7, 0x1d, 0x99, 0x26, 0x41, 0x36, 0x5b, 0x50, 0x66, 0x7c, 0xb3,
	0xc1, 0x37, 0xfb, 0x3e, 0xf8, 0xe6, 0xbb, 0x65, 0x7c, 0x05, 0xf0, 0xb5, 0x2e, 0xab, 0xfb, 0x4a,
	0x0b, 0x1f, 0x6b, 0x61, 0x87, 0xc1, 0x0f, 0xbb, 0xf9, 0x75, 0x47, 0x4e, 0x0e, 0x4a, 0x3b, 0xdb,
	0xe9, 0x29, 0xcc, 0xd4, 0xf7, 0xfd, 0x52, 0xc9, 0x8e, 0x8b, 0xc7, 0x94, 0xb9, 0x50, 0xe2, 0x1c,
	0x8e, 0x14, 0x94, 0x24, 0x48, 0xa2, 0x6c, 0xb6, 0x88, 0x59, 0xd1, 0x54, 0x4c, 0x31, 0xdc, 0xf0,
	0x29, 0x83, 0x53, 0xf5, 0xc1, 0x78, 0xe5, 0xe2, 0x11, 0xcf, 0x4d, 0xce, 0x3a, 0x61, 0xef, 0x40,
	0xd3, 0xe9, 0xdc, 0xdf, 0x97, 0xe3, 0x3e, 0xd2, 0x73, 0x63, 0x80, 0x8b, 0xa2, 0x54, 0x72, 0xe3,
	0xf1, 0x1b, 0x77, 0x2c, 0x0f, 0xbd, 0x76, 0x05, 0xcf, 0x87, 0xed, 0x9b, 0xed, 0x50, 0xab, 0x92,
	0x75, 0x5b, 0x0f, 0xfc, 0xd6, 0xd3, 0x77, 0xff, 0x3e, 0x3a, 0xf8, 0x58, 0x1f, 0xc5, 0x6d, 0x53,
	0x3e, 0x2d, 0x8a, 0x0b, 0x7f, 0x5f, 0x22, 0x81, 0x63, 0x03, 0x4a, 0x7d, 0x12, 0xf1, 0x01, 0x0e,
	0xa9, 0xe5, 0x62, 0x25, 0x8c, 0xf4, 0x38, 0x96, 0x0b, 0x7f, 0x41, 0x6b, 0x19, 0x60, 0xb5, 0x7a,
	0xb8, 0xf8, 0x13, 0x9a, 0x08, 0xbf, 0x89, 0xf6, 0x57, 0xf5, 0x43, 0xe0, 0x6b, 0x98, 0xa8, 0xb6,
	0xf1, 0x99, 0xf5, 0xd7, 0xff, 0x16, 0xe8, 0x98, 0x91, 0xf8, 0x16, 0xa6, 0xa6, 0x39, 0x44, 0x3b,
	0xb3, 0xd5, 0xd3, 0x7d, 0x4e, 0x2a, 0x75, 0x15, 0xa4, 0xa3, 0xde, 0x37, 0x4b, 0xc7, 0x8c, 0xc4,
	0xcf, 0x70, 0xe6, 0x87, 0x8e, 0xd4, 0xdb, 0xf1, 0x2a, 0xa4, 0xff, 0x9f, 0x69, 0x9f, 0x26, 0x3a,
	0xc7, 0xa7, 0xed, 0x85, 0xee, 0x73, 0xfa, 0xc2, 0x04, 0xe4, 0x5c, 0xd8, 0xb8, 0xe9, 0x3e, 0x27,
	0xef, 0xa6, 0xfa, 0x8f, 0x75, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xf7, 0x59, 0xda, 0x8c,
	0x04, 0x00, 0x00,
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
	List(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListRes, error)
	Create(ctx context.Context, in *UserCreateReq, opts ...grpc.CallOption) (*UserCreateRes, error)
	Read(ctx context.Context, in *UserReadReq, opts ...grpc.CallOption) (*UserReadRes, error)
	ReadByUsername(ctx context.Context, in *UserReadByUsernameReq, opts ...grpc.CallOption) (*UserReadByUsernameRes, error)
	Update(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateRes, error)
	Delete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteRes, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) List(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListRes, error) {
	out := new(UserListRes)
	err := c.cc.Invoke(ctx, "/api.UserService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *UserCreateReq, opts ...grpc.CallOption) (*UserCreateRes, error) {
	out := new(UserCreateRes)
	err := c.cc.Invoke(ctx, "/api.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Read(ctx context.Context, in *UserReadReq, opts ...grpc.CallOption) (*UserReadRes, error) {
	out := new(UserReadRes)
	err := c.cc.Invoke(ctx, "/api.UserService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ReadByUsername(ctx context.Context, in *UserReadByUsernameReq, opts ...grpc.CallOption) (*UserReadByUsernameRes, error) {
	out := new(UserReadByUsernameRes)
	err := c.cc.Invoke(ctx, "/api.UserService/ReadByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateRes, error) {
	out := new(UserUpdateRes)
	err := c.cc.Invoke(ctx, "/api.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteRes, error) {
	out := new(UserDeleteRes)
	err := c.cc.Invoke(ctx, "/api.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	List(context.Context, *UserListReq) (*UserListRes, error)
	Create(context.Context, *UserCreateReq) (*UserCreateRes, error)
	Read(context.Context, *UserReadReq) (*UserReadRes, error)
	ReadByUsername(context.Context, *UserReadByUsernameReq) (*UserReadByUsernameRes, error)
	Update(context.Context, *UserUpdateReq) (*UserUpdateRes, error)
	Delete(context.Context, *UserDeleteReq) (*UserDeleteRes, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).List(ctx, req.(*UserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateReq)
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
		return srv.(UserServiceServer).Create(ctx, req.(*UserCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReadReq)
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
		return srv.(UserServiceServer).Read(ctx, req.(*UserReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ReadByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReadByUsernameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ReadByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/ReadByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ReadByUsername(ctx, req.(*UserReadByUsernameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateReq)
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
		return srv.(UserServiceServer).Update(ctx, req.(*UserUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteReq)
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
		return srv.(UserServiceServer).Delete(ctx, req.(*UserDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _UserService_List_Handler,
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
			MethodName: "ReadByUsername",
			Handler:    _UserService_ReadByUsername_Handler,
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
