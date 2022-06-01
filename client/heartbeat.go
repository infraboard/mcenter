package client

import (
	"context"
	"fmt"
	"io"

	"github.com/infraboard/mcenter/apps/instance"
)

func (c *ClientSet) Heartbeat(ctx context.Context) error {
	// 1. 注册
	req := instance.NewRegistryRequest()
	ins, err := c.Instance().RegistryInstance(ctx, req)
	if err != nil {
		return err
	}

	// 2. 发送心跳
	stream, err := c.Instance().Heartbeat(ctx)
	if err != nil {
		return err
	}

	// 2.1 发送心跳
	if err := stream.Send(instance.NewHeartbeatRequest(ins.Id)); err != nil {
		c.log.Errorf("send heartbeat error, %s", err)
	}

	// 2.2 循环接收服务端返回的数据
	for {
		select {
		case <-ctx.Done():
			c.log.Infof("heatbeat stoped")
			return nil
		default:
			resp, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					c.log.Infof("receive heartbeat response error, server close")
					return nil
				}
				c.log.Errorf("receive heartbeat receive error, %s", err)
			}
			fmt.Println(resp)
		}
	}
}
