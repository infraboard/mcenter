package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/deploy"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl deploy.Service
	ctx  = context.Background()
)

func TestXxx(t *testing.T) {
	req := deploy.NewQueryDeployRequest()
	ds, err := impl.QueryDeploy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(deploy.AppName).(deploy.Service)
}
