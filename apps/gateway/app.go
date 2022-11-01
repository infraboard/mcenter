package gateway

import (
	context "context"
	"time"

	"github.com/infraboard/mcenter/common/validate"
	"github.com/rs/xid"
)

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

func (req *CreateGatewayRequest) Validate() error {
	return validate.Validate(req)
}

func NewDefaultGateway() *Gateway {
	return &Gateway{
		Spec: NewCreateGatewayRequest(),
	}
}

func NewCreateGatewayRequest() *CreateGatewayRequest {
	return &CreateGatewayRequest{}
}

func New(req *CreateGatewayRequest) (*Gateway, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	now := time.Now().UnixMilli()
	return &Gateway{
		Id:       xid.New().String(),
		CreateAt: now,
		UpdateAt: now,
		Spec:     req,
	}, nil
}

func NewGatewaySet() *GatewaySet {
	return &GatewaySet{
		Items: []*Gateway{},
	}
}

func (s *GatewaySet) Add(item *Gateway) {
	s.Items = append(s.Items, item)
}

func (req *QueryGatewayRequest) Validate() error {
	return validate.Validate(req)
}

func (req *DescribeGatewayRequest) Validate() error {
	return validate.Validate(req)
}
