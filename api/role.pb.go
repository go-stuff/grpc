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
	XXX_NoUnkeyedLiteral struct{}             `json:"-" bson:"-"`
	XXX_unrecognized     []byte               `json:"-" bson:"-"`
	XXX_sizecache        int32                `json:"-" bson:"-"`
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

type RoleListReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *RoleListReq) Reset()         { *m = RoleListReq{} }
func (m *RoleListReq) String() string { return proto.CompactTextString(m) }
func (*RoleListReq) ProtoMessage()    {}
func (*RoleListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{1}
}

func (m *RoleListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleListReq.Unmarshal(m, b)
}
func (m *RoleListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleListReq.Marshal(b, m, deterministic)
}
func (m *RoleListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleListReq.Merge(m, src)
}
func (m *RoleListReq) XXX_Size() int {
	return xxx_messageInfo_RoleListReq.Size(m)
}
func (m *RoleListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleListReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoleListReq proto.InternalMessageInfo

type RoleListRes struct {
	// return a collection of Roles
	Roles                []*Role  `protobuf:"bytes,1,rep,name=Roles,proto3" json:"Roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *RoleListRes) Reset()         { *m = RoleListRes{} }
func (m *RoleListRes) String() string { return proto.CompactTextString(m) }
func (*RoleListRes) ProtoMessage()    {}
func (*RoleListRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{2}
}

func (m *RoleListRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleListRes.Unmarshal(m, b)
}
func (m *RoleListRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleListRes.Marshal(b, m, deterministic)
}
func (m *RoleListRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleListRes.Merge(m, src)
}
func (m *RoleListRes) XXX_Size() int {
	return xxx_messageInfo_RoleListRes.Size(m)
}
func (m *RoleListRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleListRes.DiscardUnknown(m)
}

var xxx_messageInfo_RoleListRes proto.InternalMessageInfo

func (m *RoleListRes) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type RoleCreateReq struct {
	// create a Role
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	CreatedBy            string   `protobuf:"bytes,3,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty"`
	ModifiedBy           string   `protobuf:"bytes,4,opt,name=ModifiedBy,proto3" json:"ModifiedBy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
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

func (m *RoleCreateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoleCreateReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RoleCreateReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *RoleCreateReq) GetModifiedBy() string {
	if m != nil {
		return m.ModifiedBy
	}
	return ""
}

type RoleCreateRes struct {
	// reutrn an id
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
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
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
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
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
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

type RoleReadByNameReq struct {
	// read a Role by id
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *RoleReadByNameReq) Reset()         { *m = RoleReadByNameReq{} }
func (m *RoleReadByNameReq) String() string { return proto.CompactTextString(m) }
func (*RoleReadByNameReq) ProtoMessage()    {}
func (*RoleReadByNameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{7}
}

func (m *RoleReadByNameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleReadByNameReq.Unmarshal(m, b)
}
func (m *RoleReadByNameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleReadByNameReq.Marshal(b, m, deterministic)
}
func (m *RoleReadByNameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleReadByNameReq.Merge(m, src)
}
func (m *RoleReadByNameReq) XXX_Size() int {
	return xxx_messageInfo_RoleReadByNameReq.Size(m)
}
func (m *RoleReadByNameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleReadByNameReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoleReadByNameReq proto.InternalMessageInfo

func (m *RoleReadByNameReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RoleReadByNameRes struct {
	// return a Role
	Role                 *Role    `protobuf:"bytes,1,opt,name=Role,proto3" json:"Role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *RoleReadByNameRes) Reset()         { *m = RoleReadByNameRes{} }
func (m *RoleReadByNameRes) String() string { return proto.CompactTextString(m) }
func (*RoleReadByNameRes) ProtoMessage()    {}
func (*RoleReadByNameRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{8}
}

func (m *RoleReadByNameRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleReadByNameRes.Unmarshal(m, b)
}
func (m *RoleReadByNameRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleReadByNameRes.Marshal(b, m, deterministic)
}
func (m *RoleReadByNameRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleReadByNameRes.Merge(m, src)
}
func (m *RoleReadByNameRes) XXX_Size() int {
	return xxx_messageInfo_RoleReadByNameRes.Size(m)
}
func (m *RoleReadByNameRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleReadByNameRes.DiscardUnknown(m)
}

var xxx_messageInfo_RoleReadByNameRes proto.InternalMessageInfo

func (m *RoleReadByNameRes) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type RoleUpdateReq struct {
	// update a Role by id
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	ModifiedBy           string   `protobuf:"bytes,4,opt,name=ModifiedBy,proto3" json:"ModifiedBy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *RoleUpdateReq) Reset()         { *m = RoleUpdateReq{} }
func (m *RoleUpdateReq) String() string { return proto.CompactTextString(m) }
func (*RoleUpdateReq) ProtoMessage()    {}
func (*RoleUpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{9}
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

func (m *RoleUpdateReq) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *RoleUpdateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoleUpdateReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RoleUpdateReq) GetModifiedBy() string {
	if m != nil {
		return m.ModifiedBy
	}
	return ""
}

type RoleUpdateRes struct {
	// return number of Roles updated
	// should equal 1 on success
	Updated              int64    `protobuf:"varint,1,opt,name=Updated,proto3" json:"Updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *RoleUpdateRes) Reset()         { *m = RoleUpdateRes{} }
func (m *RoleUpdateRes) String() string { return proto.CompactTextString(m) }
func (*RoleUpdateRes) ProtoMessage()    {}
func (*RoleUpdateRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{10}
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
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *RoleDeleteReq) Reset()         { *m = RoleDeleteReq{} }
func (m *RoleDeleteReq) String() string { return proto.CompactTextString(m) }
func (*RoleDeleteReq) ProtoMessage()    {}
func (*RoleDeleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{11}
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
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *RoleDeleteRes) Reset()         { *m = RoleDeleteRes{} }
func (m *RoleDeleteRes) String() string { return proto.CompactTextString(m) }
func (*RoleDeleteRes) ProtoMessage()    {}
func (*RoleDeleteRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{12}
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
	proto.RegisterType((*RoleListReq)(nil), "api.RoleListReq")
	proto.RegisterType((*RoleListRes)(nil), "api.RoleListRes")
	proto.RegisterType((*RoleCreateReq)(nil), "api.RoleCreateReq")
	proto.RegisterType((*RoleCreateRes)(nil), "api.RoleCreateRes")
	proto.RegisterType((*RoleReadReq)(nil), "api.RoleReadReq")
	proto.RegisterType((*RoleReadRes)(nil), "api.RoleReadRes")
	proto.RegisterType((*RoleReadByNameReq)(nil), "api.RoleReadByNameReq")
	proto.RegisterType((*RoleReadByNameRes)(nil), "api.RoleReadByNameRes")
	proto.RegisterType((*RoleUpdateReq)(nil), "api.RoleUpdateReq")
	proto.RegisterType((*RoleUpdateRes)(nil), "api.RoleUpdateRes")
	proto.RegisterType((*RoleDeleteReq)(nil), "api.RoleDeleteReq")
	proto.RegisterType((*RoleDeleteRes)(nil), "api.RoleDeleteRes")
}

func init() { proto.RegisterFile("role.proto", fileDescriptor_48a3ff9f7c9032f8) }

var fileDescriptor_48a3ff9f7c9032f8 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x55, 0x3e, 0xba, 0xd5, 0xce, 0xaa, 0x08, 0xe6, 0x80, 0x2c, 0x8b, 0xb2, 0x2b, 0x5f, 0x58,
	0x10, 0x4a, 0xd1, 0x72, 0x41, 0x88, 0x4b, 0xcb, 0x5e, 0x2a, 0x01, 0x07, 0x03, 0x3f, 0x20, 0x6d,
	0xdc, 0xca, 0xd2, 0x2e, 0x4e, 0x63, 0x17, 0xa9, 0x37, 0x24, 0x7e, 0x15, 0xff, 0x0e, 0xd9, 0xde,
	0x78, 0x9d, 0xa4, 0x55, 0x2e, 0xbd, 0x65, 0xde, 0x9b, 0xbc, 0xbc, 0x79, 0x33, 0x01, 0x68, 0xd4,
	0x46, 0x14, 0x75, 0xa3, 0x8c, 0xc2, 0xac, 0xac, 0x25, 0x9d, 0x5f, 0x2b, 0x75, 0xbd, 0x11, 0x27,
	0x0e, 0xba, 0xb8, 0xbd, 0x3a, 0x31, 0x72, 0x2b, 0xb4, 0x29, 0xb7, 0xb5, 0xef, 0x62, 0x7f, 0x52,
	0xc8, 0xb9, 0xda, 0x08, 0x7c, 0x02, 0xe9, 0xf9, 0x9a, 0x24, 0x8b, 0x64, 0x39, 0xe5, 0xe9, 0xf9,
	0x1a, 0x11, 0xf2, 0x6f, 0xe5, 0x56, 0x90, 0xd4, 0x21, 0xee, 0x19, 0x17, 0x30, 0x5b, 0x0b, 0x7d,
	0xd9, 0xc8, 0xda, 0x48, 0xf5, 0x8b, 0x64, 0x8e, 0x8a, 0x21, 0x7c, 0x01, 0xd3, 0xcf, 0x8d, 0x28,
	0x8d, 0xa8, 0xce, 0xee, 0x48, 0xee, 0xf8, 0x3d, 0x80, 0x1f, 0x02, 0x7b, 0x6a, 0xc8, 0xc1, 0x22,
	0x59, 0xce, 0x56, 0xb4, 0xf0, 0x0e, 0x8b, 0xd6, 0x61, 0xf1, 0xa3, 0x75, 0xc8, 0xf7, 0xcd, 0xf8,
	0x12, 0xe0, 0xab, 0xaa, 0xe4, 0x95, 0x74, 0xc2, 0x13, 0x27, 0x1c, 0x21, 0xf8, 0x71, 0xcf, 0x9f,
	0x1a, 0x72, 0x38, 0x2a, 0x1d, 0x75, 0xb3, 0x23, 0x98, 0xd9, 0x04, 0xbe, 0x48, 0x6d, 0xb8, 0xb8,
	0x61, 0x45, 0x5c, 0x6a, 0x9c, 0xc3, 0x81, 0x2d, 0x35, 0x49, 0x16, 0xd9, 0x72, 0xb6, 0x9a, 0x16,
	0x65, 0x2d, 0x0b, 0x8b, 0x70, 0x8f, 0xb3, 0xbf, 0x09, 0x1c, 0xd9, 0x27, 0x6f, 0x96, 0x8b, 0x9b,
	0x10, 0x5d, 0xf2, 0x70, 0x74, 0xe9, 0x48, 0x74, 0x59, 0x3f, 0xba, 0x6e, 0x00, 0x79, 0x3f, 0x00,
	0x36, 0xef, 0x9a, 0xd0, 0xfd, 0x7d, 0xb2, 0x63, 0x3f, 0x16, 0x17, 0x65, 0x65, 0x3d, 0xf6, 0xe9,
	0xb7, 0x31, 0xad, 0xf1, 0xd8, 0x5f, 0x85, 0x6b, 0xe8, 0x0c, 0xed, 0x60, 0xf6, 0x0a, 0x9e, 0xb5,
	0xdd, 0x67, 0x77, 0x76, 0xbe, 0x07, 0xc6, 0x66, 0xab, 0x61, 0xe3, 0xa8, 0xf8, 0xad, 0x1f, 0xe5,
	0x67, 0x5d, 0xed, 0xf2, 0x7c, 0x9c, 0xd3, 0x1c, 0x4b, 0xf0, 0x75, 0xf7, 0xb3, 0x1a, 0x09, 0x1c,
	0xfa, 0xa2, 0x72, 0xdf, 0xce, 0x78, 0x5b, 0xb6, 0x61, 0xaf, 0xc5, 0x46, 0xdc, 0xeb, 0xb0, 0xd5,
	0x6a, 0x1b, 0x9c, 0x96, 0x2f, 0x82, 0xd6, 0xae, 0x5c, 0xfd, 0x4b, 0x7d, 0xf2, 0xdf, 0x45, 0xf3,
	0x5b, 0x5e, 0x0a, 0x7c, 0x03, 0xb9, 0x3d, 0x3d, 0x7c, 0x1a, 0x62, 0xd9, 0x1d, 0x26, 0xed, 0x23,
	0x1a, 0xdf, 0xc1, 0xc4, 0x2f, 0x1c, 0x31, 0x70, 0xe1, 0x0c, 0xe9, 0x10, 0xd3, 0x56, 0xdd, 0xee,
	0x22, 0x52, 0xdf, 0x1d, 0x04, 0xed, 0x23, 0x1a, 0x3f, 0x01, 0xec, 0xf7, 0x86, 0xcf, 0x3b, 0x7c,
	0xd8, 0x3a, 0xbd, 0x1f, 0x77, 0xde, 0x7c, 0x5c, 0x91, 0xb7, 0xb0, 0x52, 0x3a, 0xc4, 0xdc, 0x1b,
	0x3e, 0x94, 0xe8, 0x8d, 0x10, 0x31, 0x1d, 0x62, 0xfa, 0x62, 0xe2, 0xfe, 0xec, 0xf7, 0xff, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x69, 0xc2, 0xd3, 0x12, 0xf7, 0x04, 0x00, 0x00,
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
	List(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListRes, error)
	Create(ctx context.Context, in *RoleCreateReq, opts ...grpc.CallOption) (*RoleCreateRes, error)
	Read(ctx context.Context, in *RoleReadReq, opts ...grpc.CallOption) (*RoleReadRes, error)
	ReadByName(ctx context.Context, in *RoleReadByNameReq, opts ...grpc.CallOption) (*RoleReadByNameRes, error)
	Update(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateRes, error)
	Delete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteRes, error)
}

type roleServiceClient struct {
	cc *grpc.ClientConn
}

func NewRoleServiceClient(cc *grpc.ClientConn) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) List(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListRes, error) {
	out := new(RoleListRes)
	err := c.cc.Invoke(ctx, "/api.RoleService/List", in, out, opts...)
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

func (c *roleServiceClient) ReadByName(ctx context.Context, in *RoleReadByNameReq, opts ...grpc.CallOption) (*RoleReadByNameRes, error) {
	out := new(RoleReadByNameRes)
	err := c.cc.Invoke(ctx, "/api.RoleService/ReadByName", in, out, opts...)
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
	List(context.Context, *RoleListReq) (*RoleListRes, error)
	Create(context.Context, *RoleCreateReq) (*RoleCreateRes, error)
	Read(context.Context, *RoleReadReq) (*RoleReadRes, error)
	ReadByName(context.Context, *RoleReadByNameReq) (*RoleReadByNameRes, error)
	Update(context.Context, *RoleUpdateReq) (*RoleUpdateRes, error)
	Delete(context.Context, *RoleDeleteReq) (*RoleDeleteRes, error)
}

func RegisterRoleServiceServer(s *grpc.Server, srv RoleServiceServer) {
	s.RegisterService(&_RoleService_serviceDesc, srv)
}

func _RoleService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RoleService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).List(ctx, req.(*RoleListReq))
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

func _RoleService_ReadByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleReadByNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).ReadByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RoleService/ReadByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).ReadByName(ctx, req.(*RoleReadByNameReq))
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
			MethodName: "List",
			Handler:    _RoleService_List_Handler,
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
			MethodName: "ReadByName",
			Handler:    _RoleService_ReadByName_Handler,
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
