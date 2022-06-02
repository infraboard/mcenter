package client_test

import (
	"context"
	"testing"

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

func init() {
	err := client.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
	c = client.C()
}
