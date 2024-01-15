package middleware

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/clients/rpc"
)

// GrpcAuthUnaryServerInterceptor returns a new unary server interceptor for auth.
func GrpcAuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return NewGrpcAuther(rpc.C().Service()).WithNamespace(application.Get().AppName).Auth
}

func NewGrpcAuther(svr service.RPCClient) *GrpcAuther {
	return &GrpcAuther{
		log:     log.Sub("auther.grpc"),
		service: svr,
	}
}

// internal todo
type GrpcAuther struct {
	namespace string
	log       *zerolog.Logger
	service   service.RPCClient
}

func (a *GrpcAuther) WithNamespace(ns string) *GrpcAuther {
	a.namespace = ns
	return a
}

func (a *GrpcAuther) Auth(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// 重上下文中获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context")
	}

	clientId, clientSecret := a.getClientCredentialsFromMeta(md)

	// 校验调用的客户端凭证是否有效
	svc, err := a.validateServiceCredential(ctx, clientId, clientSecret)
	if err != nil {
		return nil, err
	}

	// 注入服务认证后的信息
	svc.InjectGrpcClientMeta(md)
	newCtx := metadata.NewIncomingContext(ctx, md)
	resp, err = handler(newCtx, req)

	// 注入自定义异常
	if err != nil {
		var setErr error
		if e, ok := err.(*exception.APIException); ok {
			setErr = grpc.SetTrailer(ctx, metadata.Pairs(exception.TRAILER_ERROR_JSON_KEY, e.ToJson()))
			err = status.Errorf(codes.Code(e.ErrorCode()), e.Error())
		} else {
			e := exception.NewAPIException(exception.InternalServerError, "系统内部错误", err.Error()).WithNamespace(a.namespace)
			setErr = grpc.SetTrailer(ctx, metadata.Pairs(exception.TRAILER_ERROR_JSON_KEY, e.ToJson()))
			err = status.Errorf(codes.Code(exception.InternalServerError), e.Error())
		}
		if setErr != nil {
			a.log.Error().Msgf("%s", setErr)
		}
	}

	return resp, err
}

func (a *GrpcAuther) getClientCredentialsFromMeta(md metadata.MD) (
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

func (a *GrpcAuther) validateServiceCredential(ctx context.Context, clientId, clientSecret string) (*service.Service, error) {
	if clientId == "" && clientSecret == "" {
		return nil, status.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}

	vsReq := service.NewValidateCredentialRequest(clientId, clientSecret)
	svc, err := a.service.ValidateCredential(ctx, vsReq)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "grpc service auth error, %s", err)
	}

	return svc, nil
}
