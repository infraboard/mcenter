// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: mcenter/apps/label/pb/rpc.proto

package label

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
	RPC_QueryLabel_FullMethodName    = "/infraboard.mcenter.label.RPC/QueryLabel"
	RPC_DescribeLabel_FullMethodName = "/infraboard.mcenter.label.RPC/DescribeLabel"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 标签Key管理
type RPCClient interface {
	// 查询标签列表
	QueryLabel(ctx context.Context, in *QueryLabelRequest, opts ...grpc.CallOption) (*LabelSet, error)
	// 查询标签列表
	DescribeLabel(ctx context.Context, in *DescribeLabelRequest, opts ...grpc.CallOption) (*Label, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) QueryLabel(ctx context.Context, in *QueryLabelRequest, opts ...grpc.CallOption) (*LabelSet, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LabelSet)
	err := c.cc.Invoke(ctx, RPC_QueryLabel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DescribeLabel(ctx context.Context, in *DescribeLabelRequest, opts ...grpc.CallOption) (*Label, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Label)
	err := c.cc.Invoke(ctx, RPC_DescribeLabel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility.
//
// 标签Key管理
type RPCServer interface {
	// 查询标签列表
	QueryLabel(context.Context, *QueryLabelRequest) (*LabelSet, error)
	// 查询标签列表
	DescribeLabel(context.Context, *DescribeLabelRequest) (*Label, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRPCServer struct{}

func (UnimplementedRPCServer) QueryLabel(context.Context, *QueryLabelRequest) (*LabelSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryLabel not implemented")
}
func (UnimplementedRPCServer) DescribeLabel(context.Context, *DescribeLabelRequest) (*Label, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeLabel not implemented")
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

func _RPC_QueryLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QueryLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryLabel(ctx, req.(*QueryLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DescribeLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DescribeLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DescribeLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DescribeLabel(ctx, req.(*DescribeLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mcenter.label.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryLabel",
			Handler:    _RPC_QueryLabel_Handler,
		},
		{
			MethodName: "DescribeLabel",
			Handler:    _RPC_DescribeLabel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mcenter/apps/label/pb/rpc.proto",
}
