package impl_test

import (
	"os"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/setting"
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
	req.Code.NotifyType = notify.NOTIFY_TYPE_SMS
	req.Code.SmsTemplateID = os.Getenv("SMS_TENCENT_TEMPLATE_ID")
	if err := env.Parse(req.Notify.SMS.TencentConfig); err != nil {
		panic(err)
	}

	s, err := impl.UpdateSetting(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}
