package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl policy.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(policy.AppName).(policy.Service)
}
