package impl_test

import (
	"testing"

	// 注册所有服务
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/test/tools"
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
