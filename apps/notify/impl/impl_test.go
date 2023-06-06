package impl_test

import (
	"context"

	// 注册所有服务
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/ioc"
)

var (
	impl notify.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(notify.AppName).(notify.Service)
}
