package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/permission"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl permission.Service
	ctx  = context.Background()
)

func TestCreateRole(t *testing.T) {
	req := permission.NewCheckPermissionRequest()
	r, err := impl.CheckPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(permission.AppName).(permission.Service)
}
