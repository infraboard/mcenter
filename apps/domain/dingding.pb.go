// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: mcenter/apps/domain/pb/dingding.proto

package domain

import (
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

type DingDingConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 开启钉钉认证
	// @gotags: bson:"enabled" json:"enabled"
	Enabled bool `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled" bson:"enabled"`
	// 飞书应用凭证, Oauth2.0时 也叫client_id
	// @gotags: bson:"client_id" json:"client_id" env:"DINGDING_CLIENT_ID"
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id" bson:"client_id" env:"DINGDING_CLIENT_ID"`
	// 飞书应用凭证, Oauth2.0时 也叫client_secret
	// @gotags: bson:"client_secret" json:"client_secret" env:"DINGDING_CLIENT_SECRET"
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret" bson:"client_secret" env:"DINGDING_CLIENT_SECRET"`
	// Oauth2.0时, 应用服务地址页面
	// @gotags: bson:"redirect_uri" json:"redirect_uri" env:"DINGDING_REDIRECT_URI"
	RedirectUri   string `protobuf:"bytes,3,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri" bson:"redirect_uri" env:"DINGDING_REDIRECT_URI"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DingDingConfig) Reset() {
	*x = DingDingConfig{}
	mi := &file_mcenter_apps_domain_pb_dingding_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DingDingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DingDingConfig) ProtoMessage() {}

func (x *DingDingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_domain_pb_dingding_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DingDingConfig.ProtoReflect.Descriptor instead.
func (*DingDingConfig) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_domain_pb_dingding_proto_rawDescGZIP(), []int{0}
}

func (x *DingDingConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *DingDingConfig) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *DingDingConfig) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *DingDingConfig) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

var File_mcenter_apps_domain_pb_dingding_proto protoreflect.FileDescriptor

var file_mcenter_apps_domain_pb_dingding_proto_rawDesc = string([]byte{
	0x0a, 0x25, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x69, 0x6e, 0x67, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x22, 0x8f, 0x01, 0x0a, 0x0e, 0x44, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x55, 0x72, 0x69, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_mcenter_apps_domain_pb_dingding_proto_rawDescOnce sync.Once
	file_mcenter_apps_domain_pb_dingding_proto_rawDescData []byte
)

func file_mcenter_apps_domain_pb_dingding_proto_rawDescGZIP() []byte {
	file_mcenter_apps_domain_pb_dingding_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_domain_pb_dingding_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_mcenter_apps_domain_pb_dingding_proto_rawDesc), len(file_mcenter_apps_domain_pb_dingding_proto_rawDesc)))
	})
	return file_mcenter_apps_domain_pb_dingding_proto_rawDescData
}

var file_mcenter_apps_domain_pb_dingding_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mcenter_apps_domain_pb_dingding_proto_goTypes = []any{
	(*DingDingConfig)(nil), // 0: infraboard.mcenter.domain.DingDingConfig
}
var file_mcenter_apps_domain_pb_dingding_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mcenter_apps_domain_pb_dingding_proto_init() }
func file_mcenter_apps_domain_pb_dingding_proto_init() {
	if File_mcenter_apps_domain_pb_dingding_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_mcenter_apps_domain_pb_dingding_proto_rawDesc), len(file_mcenter_apps_domain_pb_dingding_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcenter_apps_domain_pb_dingding_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_domain_pb_dingding_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_domain_pb_dingding_proto_msgTypes,
	}.Build()
	File_mcenter_apps_domain_pb_dingding_proto = out.File
	file_mcenter_apps_domain_pb_dingding_proto_goTypes = nil
	file_mcenter_apps_domain_pb_dingding_proto_depIdxs = nil
}
