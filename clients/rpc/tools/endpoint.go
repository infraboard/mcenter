package tools

import (
	"context"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/clients/rpc"
)

func NewEndpointRegister() *EndpointRegister {
	return &EndpointRegister{
		c: rpc.C(),
	}
}

// 业务
// 1. 先获取当前路由的所有条目
// 2. 转换成需要注册的 数据结构
// 3. 调用rpc 进行注册
type EndpointRegister struct {
	c *rpc.ClientSet
}

// gorestful 的跟理由传递给过滤
func (r *EndpointRegister) Registry(ctx context.Context, c *restful.Container, version string) error {
	entries := []*endpoint.Entry{}
	wss := c.RegisteredWebServices()
	for i := range wss {
		es := endpoint.TransferRoutesToEntry(wss[i].Routes())
		entries = append(entries, endpoint.GetPRBACEntry(es)...)
	}

	req := endpoint.NewRegistryRequest(version, entries)
	_, err := r.c.Endpoint().RegistryEndpoint(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}
