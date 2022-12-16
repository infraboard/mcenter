package rpc

import (
	"context"
	"strings"

	"github.com/infraboard/mcube/exception"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	TRAILER_ERROR_JSON_KEY = "err_json"
)

func NewExceptionUnaryClientInterceptor(ns string) *ExceptionUnaryClientInterceptor {
	return &ExceptionUnaryClientInterceptor{
		namespace: ns,
	}
}

type ExceptionUnaryClientInterceptor struct {
	namespace string
}

func (e *ExceptionUnaryClientInterceptor) WithNamespace(ns string) {
	e.namespace = ns
}

func (e *ExceptionUnaryClientInterceptor) UnaryClientInterceptor(
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

	if err != nil {
		errMsg := err.Error()
		if !strings.HasPrefix(errMsg, "rpc error:") {
			err = exception.NewAPIException(
				e.namespace,
				exception.InternalServerError,
				"服务内部异常",
				errMsg,
			)
		}
	}

	return err
}
