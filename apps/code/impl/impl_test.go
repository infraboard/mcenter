package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/ioc"
)

var (
	impl code.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(code.AppName).(code.Service)
}
