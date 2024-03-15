package endpoint

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/clients/rpc/tools"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mcube/v2/ioc/config/log"
)

func init() {
	ioc.Config().Registry(&EndpointRegistor{})
}

type EndpointRegistor struct {
	ioc.ObjectImpl

	log *zerolog.Logger
}

func (r *EndpointRegistor) Name() string {
	return AppName
}

func (r *EndpointRegistor) Init() error {
	r.log = log.Sub(AppName)

	return nil
}

func (r *EndpointRegistor) Registry() {
	// 注册服务权限条目
	r.log.Info().Msg("start registry endpoints ...")

	router := gorestful.RootRouter()
	register := tools.NewEndpointRegister()
	err := register.Registry(context.Background(), router, application.Short())
	if err != nil {
		r.log.Warn().Msgf("registry endpoints error, %s", err)
	} else {
		r.log.Debug().Msg("service endpoints registry success")
	}
}
