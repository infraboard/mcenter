// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: mcenter/apps/policy/pb/rpc.proto

package policy

import (
	role "github.com/infraboard/mcenter/apps/role"
	request "github.com/infraboard/mcube/v2/http/request"
	resource "github.com/infraboard/mcube/v2/pb/resource"
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

// QueryPolicyRequest 获取子账号列表
type QueryPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: bson:"page" json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page" bson:"page"`
	// 策略所属域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain"`
	// 用户空间
	// @gotags: json:"namespace"
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace"`
	// 用户名称
	// @gotags: json:"user_id"
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id"`
	// 用户角色
	// @gotags: json:"role_id"
	RoleId string `protobuf:"bytes,5,opt,name=role_id,json=roleId,proto3" json:"role_id"`
	// 是否查询角色相关信息
	// @gotags: json:"with_role"
	WithRole bool `protobuf:"varint,6,opt,name=with_role,json=withRole,proto3" json:"with_role"`
	// 是否查询空间相关信息
	// @gotags: json:"with_namespace"
	WithNamespace bool `protobuf:"varint,7,opt,name=with_namespace,json=withNamespace,proto3" json:"with_namespace"`
}

func (x *QueryPolicyRequest) Reset() {
	*x = QueryPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_policy_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPolicyRequest) ProtoMessage() {}

func (x *QueryPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_policy_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPolicyRequest.ProtoReflect.Descriptor instead.
func (*QueryPolicyRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_policy_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryPolicyRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryPolicyRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *QueryPolicyRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryPolicyRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *QueryPolicyRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *QueryPolicyRequest) GetWithRole() bool {
	if x != nil {
		return x.WithRole
	}
	return false
}

func (x *QueryPolicyRequest) GetWithNamespace() bool {
	if x != nil {
		return x.WithNamespace
	}
	return false
}

// DescribePolicyRequest todo
type DescribePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribePolicyRequest) Reset() {
	*x = DescribePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_policy_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePolicyRequest) ProtoMessage() {}

func (x *DescribePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_policy_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePolicyRequest.ProtoReflect.Descriptor instead.
func (*DescribePolicyRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_policy_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DescribePolicyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeletePolicyRequest todo
type DeletePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源范围
	// @gotags: json:"scope"
	Scope *resource.Scope `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope"`
	// @gotags: json:"id"
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
}

func (x *DeletePolicyRequest) Reset() {
	*x = DeletePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_policy_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyRequest) ProtoMessage() {}

func (x *DeletePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_policy_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyRequest.ProtoReflect.Descriptor instead.
func (*DeletePolicyRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_policy_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *DeletePolicyRequest) GetScope() *resource.Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *DeletePolicyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_mcenter_apps_policy_pb_rpc_proto protoreflect.FileDescriptor

var file_mcenter_apps_policy_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x18, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x62, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x63,
	0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x12, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75,
	0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x77, 0x69, 0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x77, 0x69, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x81, 0x04,
	0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x61, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x62, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x12, 0x65, 0x0a, 0x0e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x30,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x61, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x2e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x69, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_policy_pb_rpc_proto_rawDescOnce sync.Once
	file_mcenter_apps_policy_pb_rpc_proto_rawDescData = file_mcenter_apps_policy_pb_rpc_proto_rawDesc
)

func file_mcenter_apps_policy_pb_rpc_proto_rawDescGZIP() []byte {
	file_mcenter_apps_policy_pb_rpc_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_policy_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_policy_pb_rpc_proto_rawDescData)
	})
	return file_mcenter_apps_policy_pb_rpc_proto_rawDescData
}

var file_mcenter_apps_policy_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mcenter_apps_policy_pb_rpc_proto_goTypes = []interface{}{
	(*QueryPolicyRequest)(nil),     // 0: infraboard.mcenter.policy.QueryPolicyRequest
	(*DescribePolicyRequest)(nil),  // 1: infraboard.mcenter.policy.DescribePolicyRequest
	(*DeletePolicyRequest)(nil),    // 2: infraboard.mcenter.policy.DeletePolicyRequest
	(*request.PageRequest)(nil),    // 3: infraboard.mcube.page.PageRequest
	(*resource.Scope)(nil),         // 4: infraboard.mcube.resource.Scope
	(*CreatePolicyRequest)(nil),    // 5: infraboard.mcenter.policy.CreatePolicyRequest
	(*CheckPermissionRequest)(nil), // 6: infraboard.mcenter.policy.CheckPermissionRequest
	(*Policy)(nil),                 // 7: infraboard.mcenter.policy.Policy
	(*PolicySet)(nil),              // 8: infraboard.mcenter.policy.PolicySet
	(*role.Permission)(nil),        // 9: infraboard.mcenter.role.Permission
}
var file_mcenter_apps_policy_pb_rpc_proto_depIdxs = []int32{
	3, // 0: infraboard.mcenter.policy.QueryPolicyRequest.page:type_name -> infraboard.mcube.page.PageRequest
	4, // 1: infraboard.mcenter.policy.DeletePolicyRequest.scope:type_name -> infraboard.mcube.resource.Scope
	5, // 2: infraboard.mcenter.policy.RPC.CreatePolicy:input_type -> infraboard.mcenter.policy.CreatePolicyRequest
	0, // 3: infraboard.mcenter.policy.RPC.QueryPolicy:input_type -> infraboard.mcenter.policy.QueryPolicyRequest
	1, // 4: infraboard.mcenter.policy.RPC.DescribePolicy:input_type -> infraboard.mcenter.policy.DescribePolicyRequest
	2, // 5: infraboard.mcenter.policy.RPC.DeletePolicy:input_type -> infraboard.mcenter.policy.DeletePolicyRequest
	6, // 6: infraboard.mcenter.policy.RPC.CheckPermission:input_type -> infraboard.mcenter.policy.CheckPermissionRequest
	7, // 7: infraboard.mcenter.policy.RPC.CreatePolicy:output_type -> infraboard.mcenter.policy.Policy
	8, // 8: infraboard.mcenter.policy.RPC.QueryPolicy:output_type -> infraboard.mcenter.policy.PolicySet
	7, // 9: infraboard.mcenter.policy.RPC.DescribePolicy:output_type -> infraboard.mcenter.policy.Policy
	7, // 10: infraboard.mcenter.policy.RPC.DeletePolicy:output_type -> infraboard.mcenter.policy.Policy
	9, // 11: infraboard.mcenter.policy.RPC.CheckPermission:output_type -> infraboard.mcenter.role.Permission
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mcenter_apps_policy_pb_rpc_proto_init() }
func file_mcenter_apps_policy_pb_rpc_proto_init() {
	if File_mcenter_apps_policy_pb_rpc_proto != nil {
		return
	}
	file_mcenter_apps_policy_pb_permission_proto_init()
	file_mcenter_apps_policy_pb_policy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_policy_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPolicyRequest); i {
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
		file_mcenter_apps_policy_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePolicyRequest); i {
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
		file_mcenter_apps_policy_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePolicyRequest); i {
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
			RawDescriptor: file_mcenter_apps_policy_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mcenter_apps_policy_pb_rpc_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_policy_pb_rpc_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_policy_pb_rpc_proto_msgTypes,
	}.Build()
	File_mcenter_apps_policy_pb_rpc_proto = out.File
	file_mcenter_apps_policy_pb_rpc_proto_rawDesc = nil
	file_mcenter_apps_policy_pb_rpc_proto_goTypes = nil
	file_mcenter_apps_policy_pb_rpc_proto_depIdxs = nil
}
