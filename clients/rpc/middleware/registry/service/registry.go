package service

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/instance"
	mcenter "github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/grpc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
)

func init() {
	ioc.Config().Registry(&ServiceRegistor{})
}

type ServiceRegistor struct {
	ioc.ObjectImpl

	log *zerolog.Logger
	ins *instance.Instance
}

func (a *ServiceRegistor) Name() string {
	return AppName
}

func (r *ServiceRegistor) Priority() int {
	return -199
}

func (a *ServiceRegistor) Init() error {
	a.log = log.Sub(AppName)

	grpc.Get().PostStart = a.grpcPostStart
	grpc.Get().PreStop = a.grpcPreStop
	return nil
}

func (s *ServiceRegistor) grpcPostStart(ctx context.Context) error {
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

func (s *ServiceRegistor) grpcPreStop(ctx context.Context) error {
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
