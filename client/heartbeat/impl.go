package heartbeat

import (
	"context"
	"fmt"
	"io"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewHeartbeater(
	client instance.ServiceClient,
	ins *instance.Instance,
) Heartbeater {
	return &heatbeat{
		client: client,
		ins:    ins,
		log:    zap.L().Named("heartbeat"),
	}
}

type heatbeat struct {
	client instance.ServiceClient
	log    logger.Logger
	ins    *instance.Instance
}

func (h *heatbeat) Heartbeat(ctx context.Context) error {
	stream, err := h.client.Heartbeat(ctx)
	if err != nil {
		return err
	}

	// 1 发送心跳
	if err := stream.Send(instance.NewHeartbeatRequest(h.ins.Id)); err != nil {
		h.log.Errorf("send heartbeat error, %s", err)
	}

	// 2 循环接收服务端返回的数据
	for {
		select {
		case <-ctx.Done():
			h.log.Infof("heatbeat stoped")
			return nil
		default:
			resp, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					h.log.Infof("receive heartbeat response error, server close")
					return nil
				}
				h.log.Errorf("receive heartbeat receive error, %s", err)
			}
			fmt.Println(resp)
		}
	}
}
