package client_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client"
)

func TestInstanceQuery(t *testing.T) {
	c := client.C()
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
}
