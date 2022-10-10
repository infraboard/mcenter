package endpoint

import context "context"

type Service interface {
	DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*Endpoint, error)
	RPCServer
}
