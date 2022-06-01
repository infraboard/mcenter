package client

import (
	"context"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client/heartbeat"
)

func (c *ClientSet) Registry(ctx context.Context, req *instance.RegistryRequest) (heartbeat.Heartbeater, error) {
	ins, err := c.Instance().RegistryInstance(ctx, req)
	if err != nil {
		return nil, err
	}

	hb := heartbeat.NewHeartbeater(c.Instance(), ins)
	return hb, nil

}
