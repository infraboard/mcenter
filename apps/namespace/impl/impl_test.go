package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl namespace.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(namespace.AppName).(namespace.Service)
}
