package endpoint

import context "context"

const (
	AppName = "endpoint"
)

type Service interface {
	DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*Endpoint, error)
	RPCServer
}
