// Code generated by protoc-gen-go. DO NOT EDIT.
// source: role.proto

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

type Role struct {
	// @inject_tag: bson:"_id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" bson:"_id"`
	// @inject_tag: bson:"name"
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty" bson:"name"`
	// @inject_tag: bson:"description"
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty" bson:"description"`
	// @inject_tag: bson:"createdby"
	CreatedBy string `protobuf:"bytes,4,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty" bson:"createdby"`
	// @inject_tag: bson:"createdat"
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" bson:"createdat"`
	// @inject_tag: bson:"modifiedby"
	ModifiedBy string `protobuf:"bytes,6,opt,name=ModifiedBy,proto3" json:"ModifiedBy,omitempty" bson:"modifiedby"`
	// @inject_tag: bson:"modifiedat"
	ModifiedAt           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=ModifiedAt,proto3" json:"ModifiedAt,omitempty" bson:"modifiedat"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{0}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Role) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Role) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Role) GetModifiedBy() string {
	if m != nil {
		return m.ModifiedBy
	}
	return ""
}

func (m *Role) GetModifiedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ModifiedAt
	}
	return nil
}

type RoleSliceReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleSliceReq) Reset()         { *m = RoleSliceReq{} }
func (m *RoleSliceReq) String() string { return proto.CompactTextString(m) }
func (*RoleSliceReq) ProtoMessage()    {}
func (*RoleSliceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{1}
}

func (m *RoleSliceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleSliceReq.Unmarshal(m, b)
}
func (m *RoleSliceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleSliceReq.Marshal(b, m, deterministic)
}
func (m *RoleSliceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleSliceReq.Merge(m, src)
}
func (m *RoleSliceReq) XXX_Size() int {
	return xxx_messageInfo_RoleSliceReq.Size(m)
}
func (m *RoleSliceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleSliceReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoleSliceReq proto.InternalMessageInfo

type RoleSliceRes struct {
	// return a collection of Roles
	Roles                []*Role  `protobuf:"bytes,1,rep,name=Roles,proto3" json:"Roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleSliceRes) Reset()         { *m = RoleSliceRes{} }
func (m *RoleSliceRes) String() string { return proto.CompactTextString(m) }
func (*RoleSliceRes) ProtoMessage()    {}
func (*RoleSliceRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{2}
}

func (m *RoleSliceRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleSliceRes.Unmarshal(m, b)
}
func (m *RoleSliceRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleSliceRes.Marshal(b, m, deterministic)
}
func (m *RoleSliceRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleSliceRes.Merge(m, src)
}
func (m *RoleSliceRes) XXX_Size() int {
	return xxx_messageInfo_RoleSliceRes.Size(m)
}
func (m *RoleSliceRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleSliceRes.DiscardUnknown(m)
}

var xxx_messageInfo_RoleSliceRes proto.InternalMessageInfo

func (m *RoleSliceRes) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type RoleCreateReq struct {
	// create a Role
	Role                 *Role    `protobuf:"bytes,1,opt,name=Role,proto3" json:"Role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleCreateReq) Reset()         { *m = RoleCreateReq{} }
func (m *RoleCreateReq) String() string { return proto.CompactTextString(m) }
func (*RoleCreateReq) ProtoMessage()    {}
func (*RoleCreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{3}
}

func (m *RoleCreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleCreateReq.Unmarshal(m, b)
}
func (m *RoleCreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleCreateReq.Marshal(b, m, deterministic)
}
func (m *RoleCreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleCreateReq.Merge(m, src)
}
func (m *RoleCreateReq) XXX_Size() int {
	return xxx_messageInfo_RoleCreateReq.Size(m)
}
func (m *RoleCreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleCreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoleCreateReq proto.InternalMessageInfo

func (m *RoleCreateReq) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type RoleCreateRes struct {
	// reutrn an id
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleCreateRes) Reset()         { *m = RoleCreateRes{} }
func (m *RoleCreateRes) String() string { return proto.CompactTextString(m) }
func (*RoleCreateRes) ProtoMessage()    {}
func (*RoleCreateRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{4}
}

func (m *RoleCreateRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleCreateRes.Unmarshal(m, b)
}
func (m *RoleCreateRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleCreateRes.Marshal(b, m, deterministic)
}
func (m *RoleCreateRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleCreateRes.Merge(m, src)
}
func (m *RoleCreateRes) XXX_Size() int {
	return xxx_messageInfo_RoleCreateRes.Size(m)
}
func (m *RoleCreateRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleCreateRes.DiscardUnknown(m)
}

var xxx_messageInfo_RoleCreateRes proto.InternalMessageInfo

func (m *RoleCreateRes) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type RoleReadReq struct {
	// read a Role by id
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleReadReq) Reset()         { *m = RoleReadReq{} }
func (m *RoleReadReq) String() string { return proto.CompactTextString(m) }
func (*RoleReadReq) ProtoMessage()    {}
func (*RoleReadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{5}
}

func (m *RoleReadReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleReadReq.Unmarshal(m, b)
}
func (m *RoleReadReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleReadReq.Marshal(b, m, deterministic)
}
func (m *RoleReadReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleReadReq.Merge(m, src)
}
func (m *RoleReadReq) XXX_Size() int {
	return xxx_messageInfo_RoleReadReq.Size(m)
}
func (m *RoleReadReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleReadReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoleReadReq proto.InternalMessageInfo

func (m *RoleReadReq) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type RoleReadRes struct {
	// return a Role
	Role                 *Role    `protobuf:"bytes,1,opt,name=Role,proto3" json:"Role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleReadRes) Reset()         { *m = RoleReadRes{} }
func (m *RoleReadRes) String() string { return proto.CompactTextString(m) }
func (*RoleReadRes) ProtoMessage()    {}
func (*RoleReadRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{6}
}

func (m *RoleReadRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleReadRes.Unmarshal(m, b)
}
func (m *RoleReadRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleReadRes.Marshal(b, m, deterministic)
}
func (m *RoleReadRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleReadRes.Merge(m, src)
}
func (m *RoleReadRes) XXX_Size() int {
	return xxx_messageInfo_RoleReadRes.Size(m)
}
func (m *RoleReadRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleReadRes.DiscardUnknown(m)
}

var xxx_messageInfo_RoleReadRes proto.InternalMessageInfo

func (m *RoleReadRes) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type RoleUpdateReq struct {
	// update a Role by id
	Role                 *Role    `protobuf:"bytes,1,opt,name=Role,proto3" json:"Role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleUpdateReq) Reset()         { *m = RoleUpdateReq{} }
func (m *RoleUpdateReq) String() string { return proto.CompactTextString(m) }
func (*RoleUpdateReq) ProtoMessage()    {}
func (*RoleUpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{7}
}

func (m *RoleUpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleUpdateReq.Unmarshal(m, b)
}
func (m *RoleUpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleUpdateReq.Marshal(b, m, deterministic)
}
func (m *RoleUpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleUpdateReq.Merge(m, src)
}
func (m *RoleUpdateReq) XXX_Size() int {
	return xxx_messageInfo_RoleUpdateReq.Size(m)
}
func (m *RoleUpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleUpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoleUpdateReq proto.InternalMessageInfo

func (m *RoleUpdateReq) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type RoleUpdateRes struct {
	// return number of Roles updated
	// should equal 1 on success
	Updated              int64    `protobuf:"varint,1,opt,name=Updated,proto3" json:"Updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleUpdateRes) Reset()         { *m = RoleUpdateRes{} }
func (m *RoleUpdateRes) String() string { return proto.CompactTextString(m) }
func (*RoleUpdateRes) ProtoMessage()    {}
func (*RoleUpdateRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{8}
}

func (m *RoleUpdateRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleUpdateRes.Unmarshal(m, b)
}
func (m *RoleUpdateRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleUpdateRes.Marshal(b, m, deterministic)
}
func (m *RoleUpdateRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleUpdateRes.Merge(m, src)
}
func (m *RoleUpdateRes) XXX_Size() int {
	return xxx_messageInfo_RoleUpdateRes.Size(m)
}
func (m *RoleUpdateRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleUpdateRes.DiscardUnknown(m)
}

var xxx_messageInfo_RoleUpdateRes proto.InternalMessageInfo

func (m *RoleUpdateRes) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type RoleDeleteReq struct {
	// delete a Role by id
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleDeleteReq) Reset()         { *m = RoleDeleteReq{} }
func (m *RoleDeleteReq) String() string { return proto.CompactTextString(m) }
func (*RoleDeleteReq) ProtoMessage()    {}
func (*RoleDeleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{9}
}

func (m *RoleDeleteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleDeleteReq.Unmarshal(m, b)
}
func (m *RoleDeleteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleDeleteReq.Marshal(b, m, deterministic)
}
func (m *RoleDeleteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleDeleteReq.Merge(m, src)
}
func (m *RoleDeleteReq) XXX_Size() int {
	return xxx_messageInfo_RoleDeleteReq.Size(m)
}
func (m *RoleDeleteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleDeleteReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoleDeleteReq proto.InternalMessageInfo

func (m *RoleDeleteReq) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type RoleDeleteRes struct {
	// return number of Roles deleted
	// should equal 1 on success
	Deleted              int64    `protobuf:"varint,1,opt,name=Deleted,proto3" json:"Deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleDeleteRes) Reset()         { *m = RoleDeleteRes{} }
func (m *RoleDeleteRes) String() string { return proto.CompactTextString(m) }
func (*RoleDeleteRes) ProtoMessage()    {}
func (*RoleDeleteRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{10}
}

func (m *RoleDeleteRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleDeleteRes.Unmarshal(m, b)
}
func (m *RoleDeleteRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleDeleteRes.Marshal(b, m, deterministic)
}
func (m *RoleDeleteRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleDeleteRes.Merge(m, src)
}
func (m *RoleDeleteRes) XXX_Size() int {
	return xxx_messageInfo_RoleDeleteRes.Size(m)
}
func (m *RoleDeleteRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleDeleteRes.DiscardUnknown(m)
}

var xxx_messageInfo_RoleDeleteRes proto.InternalMessageInfo

func (m *RoleDeleteRes) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

func init() {
	proto.RegisterType((*Role)(nil), "api.Role")
	proto.RegisterType((*RoleSliceReq)(nil), "api.RoleSliceReq")
	proto.RegisterType((*RoleSliceRes)(nil), "api.RoleSliceRes")
	proto.RegisterType((*RoleCreateReq)(nil), "api.RoleCreateReq")
	proto.RegisterType((*RoleCreateRes)(nil), "api.RoleCreateRes")
	proto.RegisterType((*RoleReadReq)(nil), "api.RoleReadReq")
	proto.RegisterType((*RoleReadRes)(nil), "api.RoleReadRes")
	proto.RegisterType((*RoleUpdateReq)(nil), "api.RoleUpdateReq")
	proto.RegisterType((*RoleUpdateRes)(nil), "api.RoleUpdateRes")
	proto.RegisterType((*RoleDeleteReq)(nil), "api.RoleDeleteReq")
	proto.RegisterType((*RoleDeleteRes)(nil), "api.RoleDeleteRes")
}

func init() { proto.RegisterFile("role.proto", fileDescriptor_48a3ff9f7c9032f8) }

var fileDescriptor_48a3ff9f7c9032f8 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3d, 0x8f, 0xd3, 0x40,
	0x10, 0x95, 0x3f, 0x92, 0x53, 0xc6, 0x70, 0x82, 0xa9, 0x56, 0x16, 0x47, 0x2c, 0x57, 0x39, 0x04,
	0x0e, 0x0a, 0x0d, 0xa2, 0x3b, 0x70, 0x73, 0x05, 0x14, 0x0e, 0xfc, 0x00, 0x27, 0x9e, 0x44, 0x2b,
	0x39, 0xac, 0xe3, 0x35, 0x48, 0x74, 0xfc, 0x72, 0x84, 0x3c, 0x1b, 0x7f, 0x17, 0xb9, 0xce, 0xf3,
	0xde, 0x9b, 0x97, 0x37, 0x6f, 0x03, 0x50, 0xaa, 0x9c, 0xa2, 0xa2, 0x54, 0x95, 0x42, 0x27, 0x2d,
	0xa4, 0xbf, 0x3c, 0x2a, 0x75, 0xcc, 0x69, 0xcd, 0xd0, 0xee, 0xd7, 0x61, 0x5d, 0xc9, 0x13, 0xe9,
	0x2a, 0x3d, 0x15, 0x46, 0x15, 0xfe, 0xb5, 0xc1, 0x4d, 0x54, 0x4e, 0x78, 0x0b, 0xf6, 0x63, 0x2c,
	0xac, 0xc0, 0x5a, 0x2d, 0x12, 0xfb, 0x31, 0x46, 0x04, 0xf7, 0x5b, 0x7a, 0x22, 0x61, 0x33, 0xc2,
	0xdf, 0x18, 0x80, 0x17, 0x93, 0xde, 0x97, 0xb2, 0xa8, 0xa4, 0xfa, 0x29, 0x1c, 0xa6, 0xfa, 0x10,
	0xbe, 0x82, 0xc5, 0x97, 0x92, 0xd2, 0x8a, 0xb2, 0xcf, 0x7f, 0x84, 0xcb, 0x7c, 0x07, 0xe0, 0xc7,
	0x96, 0x7d, 0xa8, 0xc4, 0x2c, 0xb0, 0x56, 0xde, 0xc6, 0x8f, 0x4c, 0xc2, 0xa8, 0x49, 0x18, 0x7d,
	0x6f, 0x12, 0x26, 0x9d, 0x18, 0x5f, 0x03, 0x7c, 0x55, 0x99, 0x3c, 0x48, 0x36, 0x9e, 0xb3, 0x71,
	0x0f, 0xc1, 0x4f, 0x1d, 0xff, 0x50, 0x89, 0x9b, 0xab, 0xd6, 0x3d, 0x75, 0x78, 0x0b, 0xcf, 0xea,
	0x06, 0xb6, 0xb9, 0xdc, 0x53, 0x42, 0xe7, 0x70, 0x3d, 0x98, 0x35, 0x2e, 0x61, 0x56, 0xcf, 0x5a,
	0x58, 0x81, 0xb3, 0xf2, 0x36, 0x8b, 0x28, 0x2d, 0x64, 0x54, 0x23, 0x89, 0xc1, 0xc3, 0x08, 0x9e,
	0xd7, 0x1f, 0x26, 0x6d, 0x42, 0x67, 0xbc, 0x33, 0x9d, 0x72, 0x9b, 0x83, 0x05, 0x86, 0xc3, 0xe5,
	0x50, 0xaf, 0xc7, 0xdd, 0x87, 0x77, 0xe0, 0xb1, 0x9c, 0xd2, 0xac, 0xb6, 0x1b, 0xd3, 0x6f, 0xfb,
	0xb4, 0xbe, 0xf6, 0x6b, 0x97, 0x74, 0x3f, 0x8a, 0xec, 0x69, 0xe9, 0xee, 0x87, 0x7a, 0x8d, 0x02,
	0x6e, 0xcc, 0x90, 0xf1, 0x8a, 0x93, 0x34, 0x63, 0x73, 0x48, 0x4c, 0x39, 0x19, 0xeb, 0x71, 0xd2,
	0xfb, 0xa1, 0x80, 0xbd, 0xcc, 0xd0, 0x7a, 0x5d, 0xc6, 0xcd, 0x3f, 0xcb, 0x5c, 0xb5, 0xa5, 0xf2,
	0xb7, 0xdc, 0x13, 0xbe, 0x83, 0x19, 0xbf, 0x00, 0xbe, 0x6c, 0x03, 0x36, 0x2f, 0xe4, 0x4f, 0x20,
	0x8d, 0xef, 0x61, 0x6e, 0xfa, 0x44, 0x6c, 0xc9, 0xf6, 0x41, 0xfc, 0x29, 0xa6, 0xf1, 0x0d, 0xb8,
	0x75, 0x83, 0xf8, 0xa2, 0x2b, 0xc0, 0xf4, 0xed, 0x8f, 0x11, 0x76, 0x37, 0x37, 0xf7, 0xdc, 0xdb,
	0x42, 0xfd, 0x29, 0xc6, 0x1b, 0xe6, 0xb2, 0xde, 0x46, 0xdb, 0x93, 0x3f, 0xc5, 0xf4, 0x6e, 0xce,
	0x7f, 0xd3, 0x0f, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xf6, 0xf0, 0x4d, 0xc4, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoleServiceClient interface {
	Slice(ctx context.Context, in *RoleSliceReq, opts ...grpc.CallOption) (*RoleSliceRes, error)
	Create(ctx context.Context, in *RoleCreateReq, opts ...grpc.CallOption) (*RoleCreateRes, error)
	Read(ctx context.Context, in *RoleReadReq, opts ...grpc.CallOption) (*RoleReadRes, error)
	Update(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateRes, error)
	Delete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteRes, error)
}

type roleServiceClient struct {
	cc *grpc.ClientConn
}

func NewRoleServiceClient(cc *grpc.ClientConn) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) Slice(ctx context.Context, in *RoleSliceReq, opts ...grpc.CallOption) (*RoleSliceRes, error) {
	out := new(RoleSliceRes)
	err := c.cc.Invoke(ctx, "/api.RoleService/Slice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Create(ctx context.Context, in *RoleCreateReq, opts ...grpc.CallOption) (*RoleCreateRes, error) {
	out := new(RoleCreateRes)
	err := c.cc.Invoke(ctx, "/api.RoleService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Read(ctx context.Context, in *RoleReadReq, opts ...grpc.CallOption) (*RoleReadRes, error) {
	out := new(RoleReadRes)
	err := c.cc.Invoke(ctx, "/api.RoleService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Update(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateRes, error) {
	out := new(RoleUpdateRes)
	err := c.cc.Invoke(ctx, "/api.RoleService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Delete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteRes, error) {
	out := new(RoleDeleteRes)
	err := c.cc.Invoke(ctx, "/api.RoleService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
type RoleServiceServer interface {
	Slice(context.Context, *RoleSliceReq) (*RoleSliceRes, error)
	Create(context.Context, *RoleCreateReq) (*RoleCreateRes, error)
	Read(context.Context, *RoleReadReq) (*RoleReadRes, error)
	Update(context.Context, *RoleUpdateReq) (*RoleUpdateRes, error)
	Delete(context.Context, *RoleDeleteReq) (*RoleDeleteRes, error)
}

func RegisterRoleServiceServer(s *grpc.Server, srv RoleServiceServer) {
	s.RegisterService(&_RoleService_serviceDesc, srv)
}

func _RoleService_Slice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleSliceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Slice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RoleService/Slice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Slice(ctx, req.(*RoleSliceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RoleService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Create(ctx, req.(*RoleCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RoleService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Read(ctx, req.(*RoleReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RoleService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Update(ctx, req.(*RoleUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RoleService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Delete(ctx, req.(*RoleDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Slice",
			Handler:    _RoleService_Slice_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _RoleService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _RoleService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RoleService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RoleService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role.proto",
}
