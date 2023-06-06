package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/ioc"
)

var (
	impl instance.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(instance.AppName).(instance.Service)
}
