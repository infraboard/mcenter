package client_test

import (
	"context"
	"testing"
	"time"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client"
)

var (
	c *client.ClientSet
)

func TestInstanceQuery(t *testing.T) {
	resp, err := c.Instance().Search(
		context.Background(),
		&instance.SearchRequest{},
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestRegistry(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 注册
	hb, err := c.Registry(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	// 上报心跳
	hb.Heartbeat(ctx)
}

func init() {
	err := client.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
	c = client.C()
}
