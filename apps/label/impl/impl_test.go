package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/label"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/ioc"
)

var (
	impl label.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(label.AppName).(label.Service)
}
