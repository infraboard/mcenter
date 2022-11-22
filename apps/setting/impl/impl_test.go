package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/setting"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl setting.Service
	ctx  = context.Background()
)

func TestGetSetting(t *testing.T) {
	s, err := impl.GetSetting(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}

func TestUpdateSetting(t *testing.T) {
	req := setting.NewDefaultSetting()
	s, err := impl.UpdateSetting(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(setting.AppName).(setting.Service)
}
