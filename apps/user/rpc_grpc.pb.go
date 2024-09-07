// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.0
// source: mcenter/apps/user/pb/rpc.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RPC_QueryUser_FullMethodName    = "/infraboard.mcenter.user.RPC/QueryUser"
	RPC_DescribeUser_FullMethodName = "/infraboard.mcenter.user.RPC/DescribeUser"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service 用户服务
type RPCClient interface {
	// 查询用户列表
	QueryUser(ctx context.Context, in *QueryUserRequest, opts ...grpc.CallOption) (*UserSet, error)
	// 查询用户详情
	DescribeUser(ctx context.Context, in *DescribeUserRequest, opts ...grpc.CallOption) (*User, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) QueryUser(ctx context.Context, in *QueryUserRequest, opts ...grpc.CallOption) (*UserSet, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserSet)
	err := c.cc.Invoke(ctx, RPC_QueryUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DescribeUser(ctx context.Context, in *DescribeUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, RPC_DescribeUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility.
//
// Service 用户服务
type RPCServer interface {
	// 查询用户列表
	QueryUser(context.Context, *QueryUserRequest) (*UserSet, error)
	// 查询用户详情
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRPCServer struct{}

func (UnimplementedRPCServer) QueryUser(context.Context, *QueryUserRequest) (*UserSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUser not implemented")
}
func (UnimplementedRPCServer) DescribeUser(context.Context, *DescribeUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeUser not implemented")
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

func _RPC_QueryUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QueryUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryUser(ctx, req.(*QueryUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DescribeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DescribeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DescribeUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DescribeUser(ctx, req.(*DescribeUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mcenter.user.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryUser",
			Handler:    _RPC_QueryUser_Handler,
		},
		{
			MethodName: "DescribeUser",
			Handler:    _RPC_DescribeUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mcenter/apps/user/pb/rpc.proto",
}
