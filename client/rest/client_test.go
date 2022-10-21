package rest_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/client/rest"
	"github.com/infraboard/mcenter/test/tools"
)

var (
	c   *rest.ClientSet
	ctx = context.Background()
)

func TestQueryService(t *testing.T) {
	set, err := c.Service().QueryService(ctx, service.NewQueryServiceRequest())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(set)
}

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest(tools.AccessToken())
	tk, err := c.Token().ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tk)
}

func init() {
	c = rest.NewClient(rest.NewDefaultConfig())
}
