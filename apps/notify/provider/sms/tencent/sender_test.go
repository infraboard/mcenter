package tencent_test

import (
	"context"
	"os"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/notify/provider/sms"
	"github.com/infraboard/mcenter/apps/notify/provider/sms/tencent"
)

var (
	sender sms.SmsNotifyer
	ctx    context.Context
)

func TestSend(t *testing.T) {
	req := &sms.SendSMSRequest{}
	req.TemplateId = os.Getenv("SMS_TENCENT_TEMPLATE_ID")
	req.AddPhone(os.Getenv("TEST_CALL_NUMBER"))
	req.AddParams("600100", "30")

	err := sender.Send(ctx, req)
	if err != nil {
		panic(err)
	}
}

func init() {
	conf := &notify.TencentSmsConfig{}
	if err := env.Parse(conf); err != nil {
		panic(err)
	}
	s, err := tencent.NewSender(conf)
	if err != nil {
		panic(err)
	}
	sender = s
}
