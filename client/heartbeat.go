package client

import (
	"context"
	"fmt"
	"io"
	"log"

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
	go func() {
		for {
			select {
			case <-ctx.Done():
				c.log.Debugf("heatbeat stoped")
				return
			default:
				if err := stream.Send(instance.NewHeartbeatRequest(ins.Id)); err != nil {
					c.log.Errorf("send heartbeat error, %s", err)
				}
			}
		}
	}()

	// 2.2 然后在循环中接收服务端返回的数据
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply)
	}

	return nil
}
