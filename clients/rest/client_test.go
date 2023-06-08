package rest_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/clients/rest"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	c   *rest.ClientSet
	ctx = context.Background()
)

func TestQueryService(t *testing.T) {
	req := service.NewQueryServiceRequest()
	set, err := c.Service().QueryService(ctx, req)
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
	zap.DevelopmentSetup()
	err := rest.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
	c = rest.C()
}
