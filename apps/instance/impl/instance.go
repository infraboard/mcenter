package impl

import (
	"context"
	"io"

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
	for {
		// 处理请求
		req, err := stream.Recv()
		if err != nil {
			// 如果遇到io.EOF表示客户端流被关闭
			if err == io.EOF {
				return nil
			}
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

func (i *impl) UnRegistry(context.Context, *instance.UnregistryRequest) (
	*instance.UnregistryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegistry not implemented")
}

func (i *impl) Search(context.Context, *instance.SearchRequest) (
	*instance.InstanceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
