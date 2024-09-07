// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.0
// source: mcenter/apps/policy/pb/rpc.proto

package policy

import (
	context "context"
	namespace "github.com/infraboard/mcenter/apps/namespace"
	role "github.com/infraboard/mcenter/apps/role"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RPC_CreatePolicy_FullMethodName       = "/infraboard.mcenter.policy.RPC/CreatePolicy"
	RPC_QueryPolicy_FullMethodName        = "/infraboard.mcenter.policy.RPC/QueryPolicy"
	RPC_DescribePolicy_FullMethodName     = "/infraboard.mcenter.policy.RPC/DescribePolicy"
	RPC_DeletePolicy_FullMethodName       = "/infraboard.mcenter.policy.RPC/DeletePolicy"
	RPC_CheckPermission_FullMethodName    = "/infraboard.mcenter.policy.RPC/CheckPermission"
	RPC_AvailableNamespace_FullMethodName = "/infraboard.mcenter.policy.RPC/AvailableNamespace"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// RPC 策略服务
type RPCClient interface {
	// 创建策略
	CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// 查询策略列表
	QueryPolicy(ctx context.Context, in *QueryPolicyRequest, opts ...grpc.CallOption) (*PolicySet, error)
	// 查询策略详情
	DescribePolicy(ctx context.Context, in *DescribePolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// 删除策略
	DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// 策略鉴权
	CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*role.Permission, error)
	// 查询用户策略允许的空间
	AvailableNamespace(ctx context.Context, in *AvailableNamespaceRequest, opts ...grpc.CallOption) (*namespace.NamespaceSet, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Policy)
	err := c.cc.Invoke(ctx, RPC_CreatePolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) QueryPolicy(ctx context.Context, in *QueryPolicyRequest, opts ...grpc.CallOption) (*PolicySet, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PolicySet)
	err := c.cc.Invoke(ctx, RPC_QueryPolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DescribePolicy(ctx context.Context, in *DescribePolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Policy)
	err := c.cc.Invoke(ctx, RPC_DescribePolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Policy)
	err := c.cc.Invoke(ctx, RPC_DeletePolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*role.Permission, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(role.Permission)
	err := c.cc.Invoke(ctx, RPC_CheckPermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) AvailableNamespace(ctx context.Context, in *AvailableNamespaceRequest, opts ...grpc.CallOption) (*namespace.NamespaceSet, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(namespace.NamespaceSet)
	err := c.cc.Invoke(ctx, RPC_AvailableNamespace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility.
//
// RPC 策略服务
type RPCServer interface {
	// 创建策略
	CreatePolicy(context.Context, *CreatePolicyRequest) (*Policy, error)
	// 查询策略列表
	QueryPolicy(context.Context, *QueryPolicyRequest) (*PolicySet, error)
	// 查询策略详情
	DescribePolicy(context.Context, *DescribePolicyRequest) (*Policy, error)
	// 删除策略
	DeletePolicy(context.Context, *DeletePolicyRequest) (*Policy, error)
	// 策略鉴权
	CheckPermission(context.Context, *CheckPermissionRequest) (*role.Permission, error)
	// 查询用户策略允许的空间
	AvailableNamespace(context.Context, *AvailableNamespaceRequest) (*namespace.NamespaceSet, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRPCServer struct{}

func (UnimplementedRPCServer) CreatePolicy(context.Context, *CreatePolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicy not implemented")
}
func (UnimplementedRPCServer) QueryPolicy(context.Context, *QueryPolicyRequest) (*PolicySet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPolicy not implemented")
}
func (UnimplementedRPCServer) DescribePolicy(context.Context, *DescribePolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePolicy not implemented")
}
func (UnimplementedRPCServer) DeletePolicy(context.Context, *DeletePolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicy not implemented")
}
func (UnimplementedRPCServer) CheckPermission(context.Context, *CheckPermissionRequest) (*role.Permission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}
func (UnimplementedRPCServer) AvailableNamespace(context.Context, *AvailableNamespaceRequest) (*namespace.NamespaceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvailableNamespace not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}
func (UnimplementedRPCServer) testEmbeddedByValue()             {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	// If the following call pancis, it indicates UnimplementedRPCServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_CreatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CreatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CreatePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CreatePolicy(ctx, req.(*CreatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_QueryPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QueryPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryPolicy(ctx, req.(*QueryPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DescribePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DescribePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DescribePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DescribePolicy(ctx, req.(*DescribePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DeletePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DeletePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DeletePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DeletePolicy(ctx, req.(*DeletePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CheckPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CheckPermission(ctx, req.(*CheckPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_AvailableNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvailableNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).AvailableNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_AvailableNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).AvailableNamespace(ctx, req.(*AvailableNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mcenter.policy.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePolicy",
			Handler:    _RPC_CreatePolicy_Handler,
		},
		{
			MethodName: "QueryPolicy",
			Handler:    _RPC_QueryPolicy_Handler,
		},
		{
			MethodName: "DescribePolicy",
			Handler:    _RPC_DescribePolicy_Handler,
		},
		{
			MethodName: "DeletePolicy",
			Handler:    _RPC_DeletePolicy_Handler,
		},
		{
			MethodName: "CheckPermission",
			Handler:    _RPC_CheckPermission_Handler,
		},
		{
			MethodName: "AvailableNamespace",
			Handler:    _RPC_AvailableNamespace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mcenter/apps/policy/pb/rpc.proto",
}
