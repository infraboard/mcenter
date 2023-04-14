package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl endpoint.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(endpoint.AppName).(endpoint.Service)
}
