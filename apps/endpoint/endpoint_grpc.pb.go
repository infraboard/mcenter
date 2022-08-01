// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: apps/endpoint/pb/endpoint.proto

package endpoint

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	DescribeEndpoint(ctx context.Context, in *DescribeEndpointRequest, opts ...grpc.CallOption) (*Endpoint, error)
	QueryEndpoints(ctx context.Context, in *QueryEndpointRequest, opts ...grpc.CallOption) (*EndpointSet, error)
	RegistryEndpoint(ctx context.Context, in *RegistryRequest, opts ...grpc.CallOption) (*RegistryResponse, error)
	DeleteEndpoint(ctx context.Context, in *DeleteEndpointRequest, opts ...grpc.CallOption) (*Endpoint, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) DescribeEndpoint(ctx context.Context, in *DescribeEndpointRequest, opts ...grpc.CallOption) (*Endpoint, error) {
	out := new(Endpoint)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.endpoint.RPC/DescribeEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) QueryEndpoints(ctx context.Context, in *QueryEndpointRequest, opts ...grpc.CallOption) (*EndpointSet, error) {
	out := new(EndpointSet)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.endpoint.RPC/QueryEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) RegistryEndpoint(ctx context.Context, in *RegistryRequest, opts ...grpc.CallOption) (*RegistryResponse, error) {
	out := new(RegistryResponse)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.endpoint.RPC/RegistryEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeleteEndpoint(ctx context.Context, in *DeleteEndpointRequest, opts ...grpc.CallOption) (*Endpoint, error) {
	out := new(Endpoint)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.endpoint.RPC/DeleteEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	DescribeEndpoint(context.Context, *DescribeEndpointRequest) (*Endpoint, error)
	QueryEndpoints(context.Context, *QueryEndpointRequest) (*EndpointSet, error)
	RegistryEndpoint(context.Context, *RegistryRequest) (*RegistryResponse, error)
	DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*Endpoint, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) DescribeEndpoint(context.Context, *DescribeEndpointRequest) (*Endpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeEndpoint not implemented")
}
func (UnimplementedRPCServer) QueryEndpoints(context.Context, *QueryEndpointRequest) (*EndpointSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryEndpoints not implemented")
}
func (UnimplementedRPCServer) RegistryEndpoint(context.Context, *RegistryRequest) (*RegistryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistryEndpoint not implemented")
}
func (UnimplementedRPCServer) DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*Endpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEndpoint not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_DescribeEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DescribeEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.endpoint.RPC/DescribeEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DescribeEndpoint(ctx, req.(*DescribeEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_QueryEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.endpoint.RPC/QueryEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryEndpoints(ctx, req.(*QueryEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_RegistryEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).RegistryEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.endpoint.RPC/RegistryEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).RegistryEndpoint(ctx, req.(*RegistryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DeleteEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DeleteEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.endpoint.RPC/DeleteEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DeleteEndpoint(ctx, req.(*DeleteEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mcenter.endpoint.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeEndpoint",
			Handler:    _RPC_DescribeEndpoint_Handler,
		},
		{
			MethodName: "QueryEndpoints",
			Handler:    _RPC_QueryEndpoints_Handler,
		},
		{
			MethodName: "RegistryEndpoint",
			Handler:    _RPC_RegistryEndpoint_Handler,
		},
		{
			MethodName: "DeleteEndpoint",
			Handler:    _RPC_DeleteEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/endpoint/pb/endpoint.proto",
}
