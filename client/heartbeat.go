package client

import (
	"context"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client/lifecycle"
)

func (c *ClientSet) Registry(ctx context.Context, req *instance.RegistryRequest) (lifecycle.Lifecycler, error) {
	ins, err := c.Instance().RegistryInstance(ctx, req)
	if err != nil {
		return nil, err
	}

	lc := lifecycle.NewLifecycler(c.Instance(), ins)
	return lc, nil
}
