package endpoint

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mcube/v2/ioc/config/log"
)

const (
	AppName = "mcenter_endpoint_registry_middleware"
)

func init() {
	ioc.Api().Registry(&EndpointRegistor{})
}

type EndpointRegistor struct {
	ioc.ObjectImpl

	log      *zerolog.Logger
	endpoint endpoint.Service
}

func (r *EndpointRegistor) Name() string {
	return AppName
}

func (r *EndpointRegistor) Init() error {
	r.endpoint = ioc.Controller().Get(endpoint.AppName).(endpoint.Service)
	r.log = log.Sub(AppName)
	r.Registry()
	return nil
}

func (r *EndpointRegistor) Priority() int {
	return -199
}

func (r *EndpointRegistor) Registry() {
	// 注册服务权限条目
	r.log.Info().Msg("start registry endpoints ...")

	router := gorestful.RootRouter()

	entries := []*endpoint.Entry{}
	wss := router.RegisteredWebServices()
	for i := range wss {
		es := endpoint.TransferRoutesToEntry(wss[i].Routes())
		entries = append(entries, endpoint.GetPRBACEntry(es)...)
	}

	req := endpoint.NewRegistryRequest(application.Short(), entries)
	req.ServiceId = application.Get().AppName
	_, err := r.endpoint.RegistryEndpoint(context.Background(), req)
	if err != nil {
		r.log.Warn().Msgf("registry endpoints error, %s", err)
	} else {
		r.log.Debug().Msg("service endpoints registry success")
	}
}
