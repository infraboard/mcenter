package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"

	// 注册所有服务
	_ "github.com/infraboard/mcenter/apps/all"
	"github.com/infraboard/mcenter/apps/service"
)

var (
	impl service.MetaService
)

func TestCreateService(t *testing.T) {
	req := service.NewCreateServiceRequest()
	req.Name = "cmdb"
	req.Description = "资源中心"
	req.Owner = "admin"
	app, err := impl.CreateService(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(app)
}

func TestQueryService(t *testing.T) {
	req := service.NewQueryServiceRequest()
	set, err := impl.QueryService(context.Background(), req)
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

	impl = app.GetInternalApp(service.AppName).(service.MetaService)
}
