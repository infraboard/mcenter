package impl

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcube/exception"
)

func (i *impl) RegistryInstance(ctx context.Context, req *instance.RegistryRequest) (
	*instance.Instance, error) {
	ins, err := instance.NewInstance(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create instance error, %s", err)
	}

	if err := i.save(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) Heartbeat(stream instance.Service_HeartbeatServer) error {
	return nil
}

func (i *impl) UnRegistry(context.Context, *instance.UnregistryRequest) (
	*instance.UnregistryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegistry not implemented")
}

func (i *impl) Search(context.Context, *instance.SearchRequest) (
	*instance.InstanceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
