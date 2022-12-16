package rpc

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	TRAILER_ERROR_JSON_KEY = "err_json"
)

func ExceptionUnaryClientInterceptor(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {

	var trailer metadata.MD
	opts = append(opts, grpc.Trailer(&trailer))
	err := invoker(ctx, method, req, reply, cc, opts...)
	t := trailer.Get(TRAILER_ERROR_JSON_KEY)
	if len(t) > 0 {
		err = exception.NewAPIExceptionFromString(t[0])
	}

	return err
}
