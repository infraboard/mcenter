package feishu_test

import (
	"context"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/notify/provider/im"
	"github.com/infraboard/mcenter/apps/notify/provider/im/feishu"
)

var (
	notifyer im.ImNotifyer
	ctx      = context.Background()
)

func TestSendMessage(t *testing.T) {
	req := im.NewSendMessageRequest("2fbc2b39", "验证码", "验证码测试")
	if err := notifyer.SendMessage(ctx, req); err != nil {
		t.Fatal(err)
	}
}

func init() {
	conf := domain.NewDefaultFeishuConfig()
	if err := env.Parse(conf); err != nil {
		panic(err)
	}
	notifyer = feishu.NewFeishuNotifyer(conf)
}
