// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/role/pb/rpc.proto

package role

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_apps_role_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_role_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x1a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x97, 0x03, 0x0a, 0x03, 0x52, 0x50,
	0x43, 0x12, 0x58, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x29,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x12, 0x5b, 0x0a, 0x0c, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2c, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x6a, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x12, 0x6d, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_apps_role_pb_rpc_proto_goTypes = []interface{}{
	(*QueryRoleRequest)(nil),          // 0: infraboard.mcenter.role.QueryRoleRequest
	(*DescribeRoleRequest)(nil),       // 1: infraboard.mcenter.role.DescribeRoleRequest
	(*QueryPermissionRequest)(nil),    // 2: infraboard.mcenter.role.QueryPermissionRequest
	(*DescribePermissionRequest)(nil), // 3: infraboard.mcenter.role.DescribePermissionRequest
	(*RoleSet)(nil),                   // 4: infraboard.mcenter.role.RoleSet
	(*Role)(nil),                      // 5: infraboard.mcenter.role.Role
	(*PermissionSet)(nil),             // 6: infraboard.mcenter.role.PermissionSet
	(*Permission)(nil),                // 7: infraboard.mcenter.role.Permission
}
var file_apps_role_pb_rpc_proto_depIdxs = []int32{
	0, // 0: infraboard.mcenter.role.RPC.QueryRole:input_type -> infraboard.mcenter.role.QueryRoleRequest
	1, // 1: infraboard.mcenter.role.RPC.DescribeRole:input_type -> infraboard.mcenter.role.DescribeRoleRequest
	2, // 2: infraboard.mcenter.role.RPC.QueryPermission:input_type -> infraboard.mcenter.role.QueryPermissionRequest
	3, // 3: infraboard.mcenter.role.RPC.DescribePermission:input_type -> infraboard.mcenter.role.DescribePermissionRequest
	4, // 4: infraboard.mcenter.role.RPC.QueryRole:output_type -> infraboard.mcenter.role.RoleSet
	5, // 5: infraboard.mcenter.role.RPC.DescribeRole:output_type -> infraboard.mcenter.role.Role
	6, // 6: infraboard.mcenter.role.RPC.QueryPermission:output_type -> infraboard.mcenter.role.PermissionSet
	7, // 7: infraboard.mcenter.role.RPC.DescribePermission:output_type -> infraboard.mcenter.role.Permission
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apps_role_pb_rpc_proto_init() }
func file_apps_role_pb_rpc_proto_init() {
	if File_apps_role_pb_rpc_proto != nil {
		return
	}
	file_apps_role_pb_role_proto_init()
	file_apps_role_pb_permission_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_role_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_role_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_role_pb_rpc_proto_depIdxs,
	}.Build()
	File_apps_role_pb_rpc_proto = out.File
	file_apps_role_pb_rpc_proto_rawDesc = nil
	file_apps_role_pb_rpc_proto_goTypes = nil
	file_apps_role_pb_rpc_proto_depIdxs = nil
}