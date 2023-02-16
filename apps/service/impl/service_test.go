package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcube/app"

	// 注册所有服务
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/test/tools"
)

var (
	impl service.MetaService
	ctx  = context.Background()
)

func TestCreateService(t *testing.T) {
	req := service.NewCreateServiceRequest()
	req.Name = "cmdb"
	req.Description = "资源中心"
	req.Owner = "admin"
	app, err := impl.CreateService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(app)
}

func TestQueryService(t *testing.T) {
	req := service.NewQueryServiceRequest()
	set, err := impl.QueryService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(service.AppName).(service.MetaService)
}
