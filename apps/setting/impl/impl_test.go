package impl_test

import (
	"context"
	"os"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/mcenter/apps/notify"
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
	t.Log(s.JsonFormat())
}

func TestUpdateSetting(t *testing.T) {
	req := setting.NewDefaultSetting()

	// 读取单元测试配置
	req.Notify.Type = notify.NOTIFY_TYPE_SMS
	req.Notify.Code.SmsTemplateID = os.Getenv("SMS_TENCENT_TEMPLATE_ID")
	if err := env.Parse(req.Notify.SMS.TencentConfig); err != nil {
		panic(err)
	}

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
