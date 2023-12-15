package impl_test

import (
	"context"

	// 注册所有服务
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/v2/ioc"
)

var (
	impl notify.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.Controller().Get(notify.AppName).(notify.Service)
}
