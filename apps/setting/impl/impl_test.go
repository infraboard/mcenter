package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/setting"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl setting.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(setting.AppName).(setting.Service)
}
