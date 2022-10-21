package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl policy.Service
	ctx  = context.Background()
)

func TestCreatePolicy(t *testing.T) {
	req := policy.NewCreatePolicyRequest()
	req.Username = "test"
	req.RoleId = "xxx"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.CreateBy = "test"
	r, err := impl.CreatePolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestQueryPolicy(t *testing.T) {
	req := policy.NewQueryPolicyRequest()
	r, err := impl.QueryPolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(policy.AppName).(policy.Service)
}
