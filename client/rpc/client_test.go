package rpc_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client/rpc"
)

var (
	c *rpc.ClientSet
)

func TestInstanceRegistry(t *testing.T) {
	req := instance.NewRegistryRequest()
	req.Name = "maudit_test_01"
	req.Address = "127.0.0.1:1100"
	resp, err := c.Instance().RegistryInstance(
		context.Background(),
		req,
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestInstanceSearch(t *testing.T) {
	req := instance.NewSearchRequest()
	req.ServiceName = "maudit"
	resp, err := c.Instance().Search(
		context.Background(),
		req,
	)

	if err != nil {
		t.Fatal(err)
	}

	for i := range resp.Items {
		t.Log(resp.Items[i])
	}
}

func init() {
	err := rpc.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
	c = rpc.C()
}
