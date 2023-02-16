package middleware

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/client/rpc"
)

// GrpcAuthUnaryServerInterceptor returns a new unary server interceptor for auth.
func GrpcAuthUnaryServerInterceptor(namespace string) grpc.UnaryServerInterceptor {
	return newGrpcAuther(rpc.C().Service()).WithNamespace(namespace).Auth
}

func newGrpcAuther(svr service.RPCClient) *grpcAuther {
	return &grpcAuther{
		log:     zap.L().Named("auther.grpc"),
		service: svr,
	}
}

// internal todo
type grpcAuther struct {
	namespace string
	log       logger.Logger
	service   service.RPCClient
}

func (a *grpcAuther) WithNamespace(ns string) *grpcAuther {
	a.namespace = ns
	return a
}

func (a *grpcAuther) Auth(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// 重上下文中获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context")
	}

	clientId, clientSecret := a.GetClientCredentialsFromMeta(md)

	// 校验调用的客户端凭证是否有效
	if err := a.validateServiceCredential(clientId, clientSecret); err != nil {
		return nil, err
	}

	resp, err = handler(ctx, req)

	// 注入自定义异常
	if err != nil {
		var setErr error
		if e, ok := err.(exception.APIException); ok {
			setErr = grpc.SetTrailer(ctx, metadata.Pairs(rpc.TRAILER_ERROR_JSON_KEY, e.ToJson()))
			err = status.Errorf(codes.Code(e.ErrorCode()), e.Error())
		} else {
			e := exception.NewAPIException(a.namespace, exception.InternalServerError, "系统内部错误", err.Error())
			setErr = grpc.SetTrailer(ctx, metadata.Pairs(rpc.TRAILER_ERROR_JSON_KEY, e.ToJson()))
			err = status.Errorf(codes.Code(exception.InternalServerError), e.Error())
		}
		if setErr != nil {
			a.log.Error(setErr)
		}
	}

	return resp, err
}

func (a *grpcAuther) GetClientCredentialsFromMeta(md metadata.MD) (
	clientId, clientSecret string) {
	cids := md.Get(service.ClientHeaderKey)
	sids := md.Get(service.ClientSecretKey)
	if len(cids) > 0 {
		clientId = cids[0]
	}
	if len(sids) > 0 {
		clientSecret = sids[0]
	}
	return
}

func (a *grpcAuther) validateServiceCredential(clientId, clientSecret string) error {
	if clientId == "" && clientSecret == "" {
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}

	vsReq := service.NewValidateCredentialRequest(clientId, clientSecret)
	_, err := a.service.ValidateCredential(context.Background(), vsReq)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "service auth error, %s", err)
	}

	return nil
}
