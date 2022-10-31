package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/gateway"
)

func (s *impl) CreateGateway(ctx context.Context, req *gateway.CreateGatewayRequest) (
	*gateway.Gateway, error) {
	return nil, nil
}

func (s *impl) QueryGateway(ctx context.Context, req *gateway.QueryGatewayRequest) (
	*gateway.GatewaySet, error) {
	return nil, nil
}

func (s *impl) DescribeGateway(ctx context.Context, req *gateway.DescribeGatewayRequest) (
	*gateway.Gateway, error) {
	return nil, nil
}
