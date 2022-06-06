package rest

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/infraboard/mcenter/apps/instance"
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
	client *resty.Client
}

func (i *insImpl) RegistryInstance(ctx context.Context, req *instance.RegistryRequest) (
	*instance.Instance, error) {
	return nil, nil
}

func (i *insImpl) UnRegistry(ctx context.Context, req *instance.UnregistryRequest) (
	*instance.Instance, error) {
	return nil, nil
}

func (i *insImpl) Search(context.Context, *instance.SearchRequest) (
	*instance.InstanceSet, error) {
	return nil, nil
}

func (i *insImpl) DescribeInstance(ctx context.Context, req *instance.DescribeInstanceRequest) (
	*instance.Instance, error) {
	return nil, nil
}
