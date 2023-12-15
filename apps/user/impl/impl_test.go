package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/v2/ioc"
)

var (
	impl user.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.Controller().Get(user.AppName).(user.Service)
}
