package wx_test

import (
	"os"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token/provider/wx"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	client *wx.WechatWork
)

func TestLogin(t *testing.T) {
	// 加载测试配置
	conf := domain.NewDefaultWechatWorkConfig()
	if err := env.Parse(conf); err != nil {
		t.Fatal(err)
	}

	// 登陆
	client = wx.NewWechatWorkClient(conf)
	tk, err := client.Login(ctx, os.Getenv("WECHAT_WORK_AUTH_CODE"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func TestGetUserInfo(t *testing.T) {
	TestLogin(t)

	u, err := client.GetUserInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func init() {
	zap.DevelopmentSetup()
}
