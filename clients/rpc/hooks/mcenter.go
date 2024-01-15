package hooks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/instance"
	mcenter "github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mcenter/clients/rpc/middleware"
	"github.com/infraboard/mcenter/clients/rpc/tools"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/infraboard/mcube/v2/ioc/config/grpc"
	ioc_http "github.com/infraboard/mcube/v2/ioc/config/http"
	"github.com/infraboard/mcube/v2/ioc/config/log"
)

func NewMcenterAppHook() *McenterAppHook {
	return &McenterAppHook{
		log:                log.Sub("mcenter.app.hooks"),
		EnableApiAuth:      true,
		EnableGrpcAuth:     true,
		EnableGrpcRegistry: true,
	}
}

type McenterAppHook struct {
	log *zerolog.Logger
	ins *instance.Instance

	EnableApiAuth      bool
	EnableGrpcAuth     bool
	EnableGrpcRegistry bool
}

func (s *McenterAppHook) SetupAppHook() {
	// 开启GRPC服务注册与注销
	if s.EnableGrpcRegistry {
		grpc.Get().PostStart = s.grpcPostStart
		grpc.Get().PreStop = s.grpcPreStop
	}

	// 开启鉴权时, 需要注册功能列表与添加鉴权中间件
	if s.EnableApiAuth {
		rb := ioc_http.Get().GetRouterBuilder()
		rb.BeforeLoadHooks(s.httpBeforeLoad)
		rb.AfterLoadHooks(s.httpAfterLoad)
	}

	// 补充Grpc认证
	if s.EnableGrpcAuth {
		grpc.Get().AddInterceptors(middleware.GrpcAuthUnaryServerInterceptor())
	}
}

func (s *McenterAppHook) httpBeforeLoad(r http.Handler) {
	if router, ok := r.(*restful.Container); ok {
		// 认证中间件
		router.Filter(middleware.RestfulServerInterceptor())
	}
}

func (s *McenterAppHook) httpAfterLoad(r http.Handler) {
	if router, ok := r.(*restful.Container); ok {
		// 注册服务权限条目
		s.log.Info().Msg("start registry endpoints ...")

		register := tools.NewEndpointRegister()
		err := register.Registry(context.Background(), router, application.Short())
		if err != nil {
			s.log.Warn().Msgf("registry endpoints error, %s", err)
		} else {
			s.log.Debug().Msg("service endpoints registry success")
		}
	}
}

func (s *McenterAppHook) grpcPostStart(ctx context.Context) error {
	mcenter := mcenter.C()

	req := instance.NewRegistryRequest()
	req.Address = grpc.Get().Addr()
	ins, err := mcenter.Instance().RegistryInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("registry to mcenter error, %s", err)
	}
	s.ins = ins

	s.log.Info().Msgf("registry instance to mcenter success")
	return nil
}

func (s *McenterAppHook) grpcPreStop(ctx context.Context) error {
	mcenter := mcenter.C()

	// 提前 剔除注册中心的地址
	if s.ins != nil {
		req := instance.NewUnregistryRequest(s.ins.Id)
		if _, err := mcenter.Instance().UnRegistryInstance(ctx, req); err != nil {
			s.log.Error().Msgf("unregistry error, %s", err)
		} else {
			s.log.Info().Msg("unregistry success")
		}
	}
	return nil
}
