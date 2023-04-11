package impl_test

import (
	"context"

	"github.com/infraboard/mcube/app"

	// 注册所有服务
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/test/tools"
)

var (
	impl notify.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(notify.AppName).(notify.Service)
}
