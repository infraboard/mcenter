package heartbeat

import (
	"context"
	"fmt"
	"io"
	"time"

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
	stream instance.Service_HeartbeatClient
}

func (h *heatbeat) Heartbeat(ctx context.Context) error {
	stream, err := h.client.Heartbeat(ctx)
	if err != nil {
		return err
	}
	h.stream = stream

	go h.sender(ctx)
	go h.receiver(ctx)
	return nil
}

func (h *heatbeat) Interval() time.Duration {
	if h.ins.Config.Heartbeat.Interval == 0 {
		h.ins.Config.Heartbeat.Interval = 5
	}

	return time.Duration(h.ins.Config.Heartbeat.Interval) * time.Second
}

// 发送心跳
func (h *heatbeat) sender(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			h.log.Errorf("sender panic: %s", err)
		}
	}()

	tk := time.NewTicker(h.Interval())
	for {
		select {
		case <-ctx.Done():
			h.log.Infof("heatbeat sender stoped")
			return
		case <-tk.C:
			if err := h.stream.Send(instance.NewHeartbeatRequest(h.ins.Id)); err != nil {
				h.log.Errorf("send heartbeat error, %s", err)
			}
		}
	}
}

// 循环接收服务端返回的数据
func (h *heatbeat) receiver(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			h.log.Error("receiver panic: %s", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			h.log.Infof("receiver stoped")
			return
		default:
			resp, err := h.stream.Recv()
			if err != nil {
				if err == io.EOF {
					h.log.Infof("receive heartbeat response error, server close")
					return
				}
				h.log.Errorf("receive heartbeat receive error, %s", err)
			}
			fmt.Println(resp)
		}
	}
}
