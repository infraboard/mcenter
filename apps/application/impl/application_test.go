package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"

	// 注册所有服务
	_ "github.com/infraboard/mcenter/apps/all"
	"github.com/infraboard/mcenter/apps/application"
)

var (
	impl application.ServiceServer
)

func TestCreateApplication(t *testing.T) {
	req := application.NewCreateApplicationRequest()
	req.Name = "cmdb"
	req.Description = "资源中心"
	req.Owner = "admin"
	app, err := impl.CreateApplication(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(app)
}

func TestQueryApplication(t *testing.T) {
	req := application.NewQueryApplicationRequest()
	set, err := impl.QueryApplication(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func init() {
	zap.DevelopmentSetup()

	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	if err := app.InitAllApp(); err != nil {
		panic(err)
	}

	impl = app.GetGrpcApp(application.AppName).(application.ServiceServer)
}
