package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mcube/exception"
)

func (i *impl) RegistryInstance(ctx context.Context, req *instance.RegistryRequest) (
	*instance.Instance, error) {
	// 补充实例应用相关信息
	clientId := rpc.GetClientId(ctx)
	app, err := i.svc.DescribeService(ctx, service.NewDescribeServiceRequestByClientId(clientId))
	if err != nil {
		return nil, err
	}

	ins, err := instance.NewInstance(req, app)
	if err != nil {
		return nil, exception.NewBadRequest("validate registry instance error, %s", err)
	}

	if err := i.upsert(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) DescribeInstance(ctx context.Context, req *instance.DescribeInstanceRequest) (
	*instance.Instance, error) {
	return i.get(ctx, req.Id)
}

// 实例注销
func (i *impl) UnRegistryInstance(ctx context.Context, req *instance.UnregistryRequest) (
	*instance.Instance, error) {
	ins, err := i.DescribeInstance(ctx, instance.NewDescribeInstanceRequest(req.InstanceId))
	if err != nil {
		return nil, err
	}

	if err := i.delete(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) Search(ctx context.Context, r *instance.SearchRequest) (
	*instance.InstanceSet, error) {
	// 确定服务实例搜索的范围

	req := newSerchRequest(r)
	set, err := i.search(ctx, req)
	if err != nil {
		return nil, err
	}

	if req.ServiceName != "" {
		i.log.Debug().Msgf("search %s, address: %s", req.ServiceName, set.Items)
	}

	return set, nil
}
