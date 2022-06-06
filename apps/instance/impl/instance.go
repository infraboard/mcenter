package impl

import (
	"context"
	"io"

	"github.com/infraboard/mcenter/apps/application"
	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client/rpc/auth"
	"github.com/infraboard/mcube/exception"
)

func (i *impl) RegistryInstance(ctx context.Context, req *instance.RegistryRequest) (
	*instance.Instance, error) {
	// 补充实例应用相关信息
	clientId := auth.GetClientId(ctx)
	app, err := i.app.DescribeApplication(ctx, application.NewDescribeApplicationRequestByClientId(clientId))
	if err != nil {
		return nil, err
	}

	ins, err := instance.NewInstance(req, app)
	if err != nil {
		return nil, exception.NewBadRequest("validate create instance error, %s", err)
	}

	if err := i.upsert(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) Heartbeat(stream instance.Service_HeartbeatServer) error {
	for {
		// 处理请求
		req, err := stream.Recv()
		if err != nil {
			// 如果遇到io.EOF表示客户端流被关闭
			if err == io.EOF {
				return nil
			}

			i.log.Warnf("receive heartbeat error, %s", err)
			return nil
		}

		i.log.Debugf("instance %s", req.InstanceId)

		// 发送响应
		resp := instance.NewHeartbeatResponse()
		err = stream.Send(resp)
		if err != nil {
			// 服务端发送异常, 函数退出, 服务端流关闭
			return err
		}
	}
}

func (i *impl) DescribeInstance(ctx context.Context, req *instance.DescribeInstanceRequest) (
	*instance.Instance, error) {
	return i.get(ctx, req.Id)
}

// 实例注销
func (i *impl) UnRegistry(ctx context.Context, req *instance.UnregistryRequest) (
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
	req := newSerchRequest(r)
	set, err := i.search(ctx, req)
	if err != nil {
		return nil, err
	}

	if req.ApplicationName != "" {
		i.log.Debugf("search %s, address: %s", req.ApplicationName, set.Items)
	}

	return set, nil
}
