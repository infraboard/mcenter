package client_test

import (
	"context"
	"testing"
	"time"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client"
)

var (
	c = client.C()
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

func TestHeartbeat(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.Heartbeat(ctx); err != nil {
		t.Fatal(err)
	}
}

func init() {
	err := client.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
}
