package rest_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/application"
	"github.com/infraboard/mcenter/client/rest"
)

var (
	c *rest.ClientSet
)

func TestCreateApplicaiton(t *testing.T) {
	set, err := c.Application().QueryApplication(context.TODO(), application.NewQueryApplicationRequest())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(set)
}

func init() {
	c = rest.NewClient(rest.NewDefaultConfig())
}
