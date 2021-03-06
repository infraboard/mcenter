// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: apps/namespace/pb/namespace.proto

package namespace

import (
	request "github.com/infraboard/mcube/http/request"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Visible int32

const (
	// 默认空间是私有的
	Visible_PRIVATE Visible = 0
	// PUBLIC  公开的空间, 对所有人可见
	Visible_PUBLIC Visible = 1
)

// Enum value maps for Visible.
var (
	Visible_name = map[int32]string{
		0: "PRIVATE",
		1: "PUBLIC",
	}
	Visible_value = map[string]int32{
		"PRIVATE": 0,
		"PUBLIC":  1,
	}
)

func (x Visible) Enum() *Visible {
	p := new(Visible)
	*p = x
	return p
}

func (x Visible) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Visible) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_namespace_pb_namespace_proto_enumTypes[0].Descriptor()
}

func (Visible) Type() protoreflect.EnumType {
	return &file_apps_namespace_pb_namespace_proto_enumTypes[0]
}

func (x Visible) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Visible.Descriptor instead.
func (Visible) EnumDescriptor() ([]byte, []int) {
	return file_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{0}
}

// Namespace tenant resource container
type Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,1,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 项目修改时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,2,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 空间定义
	// @gotags: bson:"spec" json:"spec"
	Spec *CreateNamespaceRequest `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec" bson:"spec"`
}

func (x *Namespace) Reset() {
	*x = Namespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_namespace_pb_namespace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespace) ProtoMessage() {}

func (x *Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_apps_namespace_pb_namespace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespace.ProtoReflect.Descriptor instead.
func (*Namespace) Descriptor() ([]byte, []int) {
	return file_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{0}
}

func (x *Namespace) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Namespace) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Namespace) GetSpec() *CreateNamespaceRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type CreateNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 所属域ID
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 空间名称, 不允许修改
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" bson:"name"`
	// 空间负责人
	// @gotags: bson:"owner" json:"owner"
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner" bson:"owner"`
	// 禁用项目, 该项目所有人暂时都无法访问
	// @gotags: bson:"enabled" json:"enabled"
	Enabled bool `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled" bson:"enabled"`
	// 项目描述图片
	// @gotags: bson:"picture" json:"picture"
	Picture string `protobuf:"bytes,5,opt,name=picture,proto3" json:"picture" bson:"picture"`
	// 项目描述
	// @gotags: bson:"description" json:"description"
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description" bson:"description"`
	// 空间可见性, 默认是私有空间
	// @gotags: bson:"visible" json:"visible"
	Visible Visible `protobuf:"varint,7,opt,name=visible,proto3,enum=infraboard.mcenter.namespace.Visible" json:"visible" bson:"visible"`
	// 扩展信息
	// @gotags: bson:"meta" json:"meta"
	Meta map[string]string `protobuf:"bytes,8,rep,name=meta,proto3" json:"meta" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"meta"`
}

func (x *CreateNamespaceRequest) Reset() {
	*x = CreateNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_namespace_pb_namespace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceRequest) ProtoMessage() {}

func (x *CreateNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_namespace_pb_namespace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceRequest.ProtoReflect.Descriptor instead.
func (*CreateNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNamespaceRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNamespaceRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *CreateNamespaceRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *CreateNamespaceRequest) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *CreateNamespaceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateNamespaceRequest) GetVisible() Visible {
	if x != nil {
		return x.Visible
	}
	return Visible_PRIVATE
}

func (x *CreateNamespaceRequest) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type NamespaceSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 列表
	// @gotags: json:"items"
	Items []*Namespace `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *NamespaceSet) Reset() {
	*x = NamespaceSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_namespace_pb_namespace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceSet) ProtoMessage() {}

func (x *NamespaceSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_namespace_pb_namespace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceSet.ProtoReflect.Descriptor instead.
func (*NamespaceSet) Descriptor() ([]byte, []int) {
	return file_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{2}
}

func (x *NamespaceSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *NamespaceSet) GetItems() []*Namespace {
	if x != nil {
		return x.Items
	}
	return nil
}

// QueryNamespaceRequest 查询应用列表
type QueryNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页请求
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain"`
	// 空间的名称
	// @gotags: json:"name"
	Name []string `protobuf:"bytes,5,rep,name=name,proto3" json:"name"`
}

func (x *QueryNamespaceRequest) Reset() {
	*x = QueryNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_namespace_pb_namespace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryNamespaceRequest) ProtoMessage() {}

func (x *QueryNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_namespace_pb_namespace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryNamespaceRequest.ProtoReflect.Descriptor instead.
func (*QueryNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{3}
}

func (x *QueryNamespaceRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryNamespaceRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *QueryNamespaceRequest) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

// DescriptNamespaceRequest 查询应用详情
type DescriptNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain"`
	// 名称
	// @gotags: json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
}

func (x *DescriptNamespaceRequest) Reset() {
	*x = DescriptNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_namespace_pb_namespace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescriptNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescriptNamespaceRequest) ProtoMessage() {}

func (x *DescriptNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_namespace_pb_namespace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescriptNamespaceRequest.ProtoReflect.Descriptor instead.
func (*DescriptNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{4}
}

func (x *DescriptNamespaceRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *DescriptNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// DeleteNamespaceRequest todo
type DeleteNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain"`
	// 名称
	// @gotags: json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
}

func (x *DeleteNamespaceRequest) Reset() {
	*x = DeleteNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_namespace_pb_namespace_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNamespaceRequest) ProtoMessage() {}

func (x *DeleteNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_namespace_pb_namespace_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNamespaceRequest.ProtoReflect.Descriptor instead.
func (*DeleteNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteNamespaceRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *DeleteNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_apps_namespace_pb_namespace_proto protoreflect.FileDescriptor

var file_apps_namespace_pb_namespace_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x48, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x22, 0xfe, 0x02, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x52, 0x07,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x52, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d,
	0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x63, 0x0a, 0x0c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x7b, 0x0a, 0x15, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x44,
	0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x22, 0x0a, 0x07, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x32, 0xee, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43,
	0x12, 0x71, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x33, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x53, 0x65, 0x74, 0x12, 0x74, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x36, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_apps_namespace_pb_namespace_proto_rawDescOnce sync.Once
	file_apps_namespace_pb_namespace_proto_rawDescData = file_apps_namespace_pb_namespace_proto_rawDesc
)

func file_apps_namespace_pb_namespace_proto_rawDescGZIP() []byte {
	file_apps_namespace_pb_namespace_proto_rawDescOnce.Do(func() {
		file_apps_namespace_pb_namespace_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_namespace_pb_namespace_proto_rawDescData)
	})
	return file_apps_namespace_pb_namespace_proto_rawDescData
}

var file_apps_namespace_pb_namespace_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_namespace_pb_namespace_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_apps_namespace_pb_namespace_proto_goTypes = []interface{}{
	(Visible)(0),                     // 0: infraboard.mcenter.namespace.Visible
	(*Namespace)(nil),                // 1: infraboard.mcenter.namespace.Namespace
	(*CreateNamespaceRequest)(nil),   // 2: infraboard.mcenter.namespace.CreateNamespaceRequest
	(*NamespaceSet)(nil),             // 3: infraboard.mcenter.namespace.NamespaceSet
	(*QueryNamespaceRequest)(nil),    // 4: infraboard.mcenter.namespace.QueryNamespaceRequest
	(*DescriptNamespaceRequest)(nil), // 5: infraboard.mcenter.namespace.DescriptNamespaceRequest
	(*DeleteNamespaceRequest)(nil),   // 6: infraboard.mcenter.namespace.DeleteNamespaceRequest
	nil,                              // 7: infraboard.mcenter.namespace.CreateNamespaceRequest.MetaEntry
	(*request.PageRequest)(nil),      // 8: infraboard.mcube.page.PageRequest
}
var file_apps_namespace_pb_namespace_proto_depIdxs = []int32{
	2, // 0: infraboard.mcenter.namespace.Namespace.spec:type_name -> infraboard.mcenter.namespace.CreateNamespaceRequest
	0, // 1: infraboard.mcenter.namespace.CreateNamespaceRequest.visible:type_name -> infraboard.mcenter.namespace.Visible
	7, // 2: infraboard.mcenter.namespace.CreateNamespaceRequest.meta:type_name -> infraboard.mcenter.namespace.CreateNamespaceRequest.MetaEntry
	1, // 3: infraboard.mcenter.namespace.NamespaceSet.items:type_name -> infraboard.mcenter.namespace.Namespace
	8, // 4: infraboard.mcenter.namespace.QueryNamespaceRequest.page:type_name -> infraboard.mcube.page.PageRequest
	4, // 5: infraboard.mcenter.namespace.RPC.QueryNamespace:input_type -> infraboard.mcenter.namespace.QueryNamespaceRequest
	5, // 6: infraboard.mcenter.namespace.RPC.DescribeNamespace:input_type -> infraboard.mcenter.namespace.DescriptNamespaceRequest
	3, // 7: infraboard.mcenter.namespace.RPC.QueryNamespace:output_type -> infraboard.mcenter.namespace.NamespaceSet
	1, // 8: infraboard.mcenter.namespace.RPC.DescribeNamespace:output_type -> infraboard.mcenter.namespace.Namespace
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_apps_namespace_pb_namespace_proto_init() }
func file_apps_namespace_pb_namespace_proto_init() {
	if File_apps_namespace_pb_namespace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_namespace_pb_namespace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namespace); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_namespace_pb_namespace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNamespaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_namespace_pb_namespace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_namespace_pb_namespace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryNamespaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_namespace_pb_namespace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescriptNamespaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_namespace_pb_namespace_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNamespaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_namespace_pb_namespace_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_namespace_pb_namespace_proto_goTypes,
		DependencyIndexes: file_apps_namespace_pb_namespace_proto_depIdxs,
		EnumInfos:         file_apps_namespace_pb_namespace_proto_enumTypes,
		MessageInfos:      file_apps_namespace_pb_namespace_proto_msgTypes,
	}.Build()
	File_apps_namespace_pb_namespace_proto = out.File
	file_apps_namespace_pb_namespace_proto_rawDesc = nil
	file_apps_namespace_pb_namespace_proto_goTypes = nil
	file_apps_namespace_pb_namespace_proto_depIdxs = nil
}
