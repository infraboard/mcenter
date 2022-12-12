package wechatwork_test

import (
	"os"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token/provider/wechatwork"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	client *wechatwork.WechatWork

	code = os.Getenv("WECHAT_WORK_AUTH_CODE")
)

func TestLogin(t *testing.T) {
	// 加载测试配置
	conf := domain.NewDefaultWechatWorkConfig()
	if err := env.Parse(conf); err != nil {
		t.Fatal(err)
	}

	// 登陆
	client = wechatwork.NewWechatWorkClient(conf)
	tk, err := client.Login(ctx, code)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func TestGetUserInfo(t *testing.T) {
	TestLogin(t)

	u, err := client.GetUserInfo(ctx, code)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func init() {
	zap.DevelopmentSetup()
}
