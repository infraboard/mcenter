package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/permission"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl permission.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(permission.AppName).(permission.Service)
}
