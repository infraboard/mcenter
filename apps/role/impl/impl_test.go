package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl role.Service
	ctx  = context.Background()
)

func TestCreateRole(t *testing.T) {
	req := role.NewCreateRoleRequest()
	r, err := impl.CreateRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(role.AppName).(role.Service)
}
