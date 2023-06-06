package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/ioc"
)

var (
	impl endpoint.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(endpoint.AppName).(endpoint.Service)
}
