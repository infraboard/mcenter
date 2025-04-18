// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: mcenter/apps/notify/pb/rpc.proto

package notify

import (
	request "github.com/infraboard/mcube/v2/http/request"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 查询发送记录
type QueryRecordRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 通知类型
	// @gotags: json:"notify_tye"
	NotifyTye *NOTIFY_TYPE `protobuf:"varint,2,opt,name=notify_tye,json=notifyTye,proto3,enum=infraboard.mcenter.notify.NOTIFY_TYPE,oneof" json:"notify_tye"`
	// 域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain"`
	// 空间
	// @gotags: json:"namespace"
	Namespace     string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryRecordRequest) Reset() {
	*x = QueryRecordRequest{}
	mi := &file_mcenter_apps_notify_pb_rpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRecordRequest) ProtoMessage() {}

func (x *QueryRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_notify_pb_rpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRecordRequest.ProtoReflect.Descriptor instead.
func (*QueryRecordRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_notify_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRecordRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryRecordRequest) GetNotifyTye() NOTIFY_TYPE {
	if x != nil && x.NotifyTye != nil {
		return *x.NotifyTye
	}
	return NOTIFY_TYPE_MAIL
}

func (x *QueryRecordRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *QueryRecordRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_mcenter_apps_notify_pb_rpc_proto protoreflect.FileDescriptor

var file_mcenter_apps_notify_pb_rpc_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x1a, 0x18, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x62, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01, 0x0a,
	0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x4a, 0x0a, 0x0a, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x79, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x4e, 0x4f, 0x54, 0x49,
	0x46, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x54, 0x79, 0x65, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x79, 0x65, 0x32, 0xc8, 0x01, 0x0a,
	0x03, 0x52, 0x50, 0x43, 0x12, 0x5d, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x2c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x62, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x2d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_mcenter_apps_notify_pb_rpc_proto_rawDescOnce sync.Once
	file_mcenter_apps_notify_pb_rpc_proto_rawDescData []byte
)

func file_mcenter_apps_notify_pb_rpc_proto_rawDescGZIP() []byte {
	file_mcenter_apps_notify_pb_rpc_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_notify_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_mcenter_apps_notify_pb_rpc_proto_rawDesc), len(file_mcenter_apps_notify_pb_rpc_proto_rawDesc)))
	})
	return file_mcenter_apps_notify_pb_rpc_proto_rawDescData
}

var file_mcenter_apps_notify_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mcenter_apps_notify_pb_rpc_proto_goTypes = []any{
	(*QueryRecordRequest)(nil),  // 0: infraboard.mcenter.notify.QueryRecordRequest
	(*request.PageRequest)(nil), // 1: infraboard.mcube.page.PageRequest
	(NOTIFY_TYPE)(0),            // 2: infraboard.mcenter.notify.NOTIFY_TYPE
	(*SendNotifyRequest)(nil),   // 3: infraboard.mcenter.notify.SendNotifyRequest
	(*Record)(nil),              // 4: infraboard.mcenter.notify.Record
	(*RecordSet)(nil),           // 5: infraboard.mcenter.notify.RecordSet
}
var file_mcenter_apps_notify_pb_rpc_proto_depIdxs = []int32{
	1, // 0: infraboard.mcenter.notify.QueryRecordRequest.page:type_name -> infraboard.mcube.page.PageRequest
	2, // 1: infraboard.mcenter.notify.QueryRecordRequest.notify_tye:type_name -> infraboard.mcenter.notify.NOTIFY_TYPE
	3, // 2: infraboard.mcenter.notify.RPC.SendNotify:input_type -> infraboard.mcenter.notify.SendNotifyRequest
	0, // 3: infraboard.mcenter.notify.RPC.QueryRecord:input_type -> infraboard.mcenter.notify.QueryRecordRequest
	4, // 4: infraboard.mcenter.notify.RPC.SendNotify:output_type -> infraboard.mcenter.notify.Record
	5, // 5: infraboard.mcenter.notify.RPC.QueryRecord:output_type -> infraboard.mcenter.notify.RecordSet
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mcenter_apps_notify_pb_rpc_proto_init() }
func file_mcenter_apps_notify_pb_rpc_proto_init() {
	if File_mcenter_apps_notify_pb_rpc_proto != nil {
		return
	}
	file_mcenter_apps_notify_pb_notify_proto_init()
	file_mcenter_apps_notify_pb_rpc_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_mcenter_apps_notify_pb_rpc_proto_rawDesc), len(file_mcenter_apps_notify_pb_rpc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mcenter_apps_notify_pb_rpc_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_notify_pb_rpc_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_notify_pb_rpc_proto_msgTypes,
	}.Build()
	File_mcenter_apps_notify_pb_rpc_proto = out.File
	file_mcenter_apps_notify_pb_rpc_proto_goTypes = nil
	file_mcenter_apps_notify_pb_rpc_proto_depIdxs = nil
}
