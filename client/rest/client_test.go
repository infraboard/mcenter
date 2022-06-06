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
	app, err := c.Application().CreateApplication(context.TODO(), application.NewCreateApplicationRequest())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(app)
}

func init() {

}
