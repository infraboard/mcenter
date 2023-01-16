package deploy

import (
	context "context"

	"github.com/infraboard/mcube/http/request"
)

const (
	AppName = "deploy"
)

type Service interface {
	CreateDeploy(context.Context, *CreateDeployRequest) (*Deploy, error)
	DeleteDeploy(context.Context, *DeleteDeployRequest) (*Deploy, error)
	RPCServer
}

func NewQueryDeployRequest() *QueryDeployRequest {
	return &QueryDeployRequest{
		Page: request.NewDefaultPageRequest(),
	}
}
