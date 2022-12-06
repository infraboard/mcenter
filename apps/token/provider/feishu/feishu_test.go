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
	Login()

	u, err := client.GetUserInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func Login() {
	// 加载测试配置
	conf := domain.NewDefaultFeishuConfig()
	if err := env.Parse(conf); err != nil {
		panic(err)
	}

	// 登陆
	client = feishu.NewFeishuClient(conf)
	err := client.Login(ctx, os.Getenv("AUTH_CODE"))
	if err != nil {
		panic(err)
	}
}

func init() {
	zap.DevelopmentSetup()
}
