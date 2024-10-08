package grpc

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	ioc_grpc "github.com/infraboard/mcube/v2/ioc/config/grpc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/clients/rpc"
)

func init() {
	ioc.Config().Registry(&GrpcAuther{})
}

// internal todo
type GrpcAuther struct {
	namespace string
	log       *zerolog.Logger
	service   service.RPCClient

	ioc.ObjectImpl
}

func (a *GrpcAuther) Init() error {
	a.namespace = application.Get().AppName
	a.log = log.Sub(AppName)
	a.service = rpc.C().Service()

	// 注册grpc认证中间件
	ioc_grpc.Get().AddInterceptors(a.Auth)
	return nil
}

func (a *GrpcAuther) Name() string {
	return AppName
}

func (m *GrpcAuther) Priority() int {
	return 299
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
		if e, ok := err.(*exception.ApiException); ok {
			setErr = grpc.SetTrailer(ctx, metadata.Pairs(exception.TRAILER_ERROR_JSON_KEY, e.ToJson()))
			err = status.Errorf(codes.Code(e.ErrorCode()), e.Error())
		} else {
			e := exception.NewApiException(exception.CODE_INTERNAL_SERVER_ERROR, "系统内部错误").WithMessagef(err.Error()).WithNamespace(a.namespace)
			setErr = grpc.SetTrailer(ctx, metadata.Pairs(exception.TRAILER_ERROR_JSON_KEY, e.ToJson()))
			err = status.Errorf(codes.Code(exception.CODE_INTERNAL_SERVER_ERROR), e.Error())
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
