package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/v2/ioc"
)

var (
	impl policy.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.Controller().Get(policy.AppName).(policy.Service)
}
