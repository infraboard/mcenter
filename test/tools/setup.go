package tools

import (
	"os"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger/zap"

	// 注册所有服务
	_ "github.com/infraboard/mcenter/apps"
)

func DevelopmentSetup() {
	// 初始化日志实例
	zap.DevelopmentSetup()

	// 初始化配置, 提前配置好/etc/unit_test.env
	req := ioc.NewLoadConfigRequest()
	err := ioc.ConfigIocObject(req)
	if err != nil {
		panic(err)
	}
}

func AccessToken() string {
	return os.Getenv("MCENTER_ACCESS_TOKEN")
}
