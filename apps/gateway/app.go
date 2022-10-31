package gateway

import context "context"

const (
	AppName = "gateway"
)

type Service interface {
	CreateGateway(context.Context, *CreateGatewayRequest) (*Gateway, error)
	RPCServer
}

func NewDefaultTraefikConfig() *TraefikConfig {
	return &TraefikConfig{
		Endpoints: []string{"127.0.0.1:2379"},
	}
}
