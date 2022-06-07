package rest

import (
	"context"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcube/client/rest"
	"github.com/infraboard/mcube/http/response"
)

type InstanceService interface {
	// 实例注册
	RegistryInstance(context.Context, *instance.RegistryRequest) (*instance.Instance, error)
	// 实例注销
	UnRegistry(context.Context, *instance.UnregistryRequest) (*instance.Instance, error)
	// 实例搜索, 用于客户端做服务发现
	Search(context.Context, *instance.SearchRequest) (*instance.InstanceSet, error)
	// 查询实例详情
	DescribeInstance(context.Context, *instance.DescribeInstanceRequest) (*instance.Instance, error)
}

type insImpl struct {
	client *rest.RESTClient
}

func (i *insImpl) RegistryInstance(ctx context.Context, req *instance.RegistryRequest) (
	*instance.Instance, error) {
	ins := instance.NewDefaultInstance()
	resp := response.NewData(ins)

	err := i.client.
		Post("instance").
		Body(req).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, err
	}

	return ins, nil
}

func (i *insImpl) UnRegistry(ctx context.Context, req *instance.UnregistryRequest) (
	*instance.Instance, error) {
	ins := instance.NewDefaultInstance()
	resp := response.NewData(ins)

	err := i.client.
		Delete("instance/" + req.InstanceId).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, err
	}
	return ins, nil
}

func (i *insImpl) Search(ctx context.Context, req *instance.SearchRequest) (
	*instance.InstanceSet, error) {
	set := instance.NewInstanceSet()
	resp := response.NewData(set)

	err := i.client.
		Get("instance/").
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, err
	}

	return set, nil
}

func (i *insImpl) DescribeInstance(ctx context.Context, req *instance.DescribeInstanceRequest) (
	*instance.Instance, error) {
	ins := instance.NewDefaultInstance()
	resp := response.NewData(ins)

	err := i.client.
		Get("instance/" + req.Id).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, err
	}

	return ins, nil
}
