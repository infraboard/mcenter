package grpc

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/infraboard/mcenter/apps/service"
	ioc_grpc "github.com/infraboard/mcube/v2/ioc/config/grpc"
)

func init() {
	ioc.Config().Registry(&grpcAuther{})
}

// internal todo
type grpcAuther struct {
	log *zerolog.Logger
	svc service.MetaService

	ioc.ObjectImpl
}

func (a *grpcAuther) Init() error {
	a.log = log.Sub("Grpc Auther")
	a.svc = ioc.Controller().Get(service.AppName).(service.MetaService)

	// 注册认证中间件
	ioc_grpc.Get().AddInterceptors(a.Auth)
	return nil
}

func (a *grpcAuther) Name() string {
	return "grpc_auth"
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
	svc, err := a.validateServiceCredential(clientId, clientSecret)
	if err != nil {
		return nil, err
	}

	// 注入服务认证后的信息
	svc.InjectGrpcClientMeta(md)
	ctxWrapper := metadata.NewIncomingContext(ctx, md)
	resp, err = handler(ctxWrapper, req)

	// 注入自定义异常
	if err != nil {
		var setErr error
		if e, ok := err.(*exception.ApiException); ok {
			setErr = grpc.SetTrailer(ctx, metadata.Pairs(exception.TRAILER_ERROR_JSON_KEY, e.ToJson()))
			err = status.Errorf(codes.Code(e.ErrorCode()), e.Error())
		} else {
			e := exception.NewApiException(exception.CODE_INTERNAL_SERVER_ERROR, "系统内部错误").WithMessage(err.Error()).WithNamespace(application.Get().AppName)
			setErr = grpc.SetTrailer(ctx, metadata.Pairs(exception.TRAILER_ERROR_JSON_KEY, e.ToJson()))
		}
		if setErr != nil {
			a.log.Error().Msgf("%s", setErr)
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

func (a *grpcAuther) validateServiceCredential(clientId, clientSecret string) (*service.Service, error) {
	if clientId == "" && clientSecret == "" {
		return nil, status.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}

	vsReq := service.NewValidateCredentialRequest(clientId, clientSecret)
	svc, err := a.svc.ValidateCredential(context.Background(), vsReq)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "service auth error, %s", err)
	}

	return svc, nil
}
