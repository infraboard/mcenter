// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.0
// source: mcenter/apps/notify/pb/rpc.proto

package notify

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
	RPC_SendNotify_FullMethodName  = "/infraboard.mcenter.notify.RPC/SendNotify"
	RPC_QueryRecord_FullMethodName = "/infraboard.mcenter.notify.RPC/QueryRecord"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 通知服务
type RPCClient interface {
	// 用户消息通知
	SendNotify(ctx context.Context, in *SendNotifyRequest, opts ...grpc.CallOption) (*Record, error)
	// 查询发送记录
	QueryRecord(ctx context.Context, in *QueryRecordRequest, opts ...grpc.CallOption) (*RecordSet, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) SendNotify(ctx context.Context, in *SendNotifyRequest, opts ...grpc.CallOption) (*Record, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Record)
	err := c.cc.Invoke(ctx, RPC_SendNotify_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) QueryRecord(ctx context.Context, in *QueryRecordRequest, opts ...grpc.CallOption) (*RecordSet, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecordSet)
	err := c.cc.Invoke(ctx, RPC_QueryRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility.
//
// 通知服务
type RPCServer interface {
	// 用户消息通知
	SendNotify(context.Context, *SendNotifyRequest) (*Record, error)
	// 查询发送记录
	QueryRecord(context.Context, *QueryRecordRequest) (*RecordSet, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRPCServer struct{}

func (UnimplementedRPCServer) SendNotify(context.Context, *SendNotifyRequest) (*Record, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotify not implemented")
}
func (UnimplementedRPCServer) QueryRecord(context.Context, *QueryRecordRequest) (*RecordSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRecord not implemented")
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

func _RPC_SendNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).SendNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_SendNotify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).SendNotify(ctx, req.(*SendNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_QueryRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QueryRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryRecord(ctx, req.(*QueryRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mcenter.notify.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendNotify",
			Handler:    _RPC_SendNotify_Handler,
		},
		{
			MethodName: "QueryRecord",
			Handler:    _RPC_QueryRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mcenter/apps/notify/pb/rpc.proto",
}
