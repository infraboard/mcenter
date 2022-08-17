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

func TestInstanceSearch(t *testing.T) {
	req := instance.NewSearchRequest()
	resp, err := c.Instance().Search(
		context.Background(),
		req,
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestInstanceRegistry(t *testing.T) {
	req := instance.NewRegistryRequest()
	req.Name = "test_instance_01"
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

func init() {
	err := rpc.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
	c = rpc.C()
}
