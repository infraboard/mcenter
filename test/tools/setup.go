package tools

import (
	"os"

	"github.com/infraboard/mcube/v2/ioc"

	// 注册所有服务
	_ "github.com/infraboard/mcenter/apps"
)

func DevelopmentSetup() {
	os.Setenv("IP2REGION_ENABLE", "false")
	ioc.DevelopmentSetup()
}

func AccessToken() string {
	return os.Getenv("MCENTER_ACCESS_TOKEN")
}
