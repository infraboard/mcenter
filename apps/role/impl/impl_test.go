package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/v2/ioc"
)

var (
	impl role.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(role.AppName).(role.Service)
}
