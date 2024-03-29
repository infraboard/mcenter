package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/label"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/v2/ioc"
)

var (
	impl label.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.Controller().Get(label.AppName).(label.Service)
}
