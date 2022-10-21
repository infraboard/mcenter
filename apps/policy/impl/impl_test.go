package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl policy.Service
	ctx  = context.Background()
)

func TestCreateRole(t *testing.T) {
	req := policy.NewCreatePolicyRequest()
	r, err := impl.CreatePolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(policy.AppName).(policy.Service)
}
