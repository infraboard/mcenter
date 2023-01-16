package deploy

import (
	context "context"
	"fmt"
	"time"

	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	"github.com/rs/xid"
)

const (
	AppName = "deploy"
)

type Service interface {
	CreateDeploy(context.Context, *CreateDeployRequest) (*Deploy, error)
	DeleteDeploy(context.Context, *DeleteDeployRequest) (*Deploy, error)
	RPCServer
}

// New 新建一个domain
func New(req *CreateDeployRequest) (*Deploy, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	d := &Deploy{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
	}

	return d, nil
}

func (req *CreateDeployRequest) Validate() error {
	return validate.Validate(req)
}

func NewQueryDeployRequest() *QueryDeployRequest {
	return &QueryDeployRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewCreateDeployRequest() *CreateDeployRequest {
	return &CreateDeployRequest{}
}

// Validate 校验请求是否合法
func (req *UpdateDeployRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("id required")
	}
	if req.UpdateMode.Equal(pb_request.UpdateMode_PUT) {
		return validate.Validate(req)
	}

	return nil
}

func NewDescribeDeployRequest(id string) *DescribeDeployRequest {
	return &DescribeDeployRequest{
		DescribeValue: id,
	}
}

func (req *DescribeDeployRequest) Validate() error {
	return validate.Validate(req)
}
