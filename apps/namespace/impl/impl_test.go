package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/v2/ioc"
)

var (
	impl namespace.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.Controller().Get(namespace.AppName).(namespace.Service)
}
