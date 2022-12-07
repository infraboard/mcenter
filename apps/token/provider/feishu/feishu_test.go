package feishu_test

import (
	"os"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token/provider/feishu"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	client *feishu.Feishu
)

func TestGetUserInfo(t *testing.T) {
	TestLogin(t)

	u, err := client.GetUserInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestLogin(t *testing.T) {
	// 加载测试配置
	conf := domain.NewDefaultFeishuConfig()
	if err := env.Parse(conf); err != nil {
		t.Fatal(err)
	}

	// 登陆
	client = feishu.NewFeishuClient(conf)
	tk, err := client.Login(ctx, os.Getenv("AUTH_CODE"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func init() {
	zap.DevelopmentSetup()
}
