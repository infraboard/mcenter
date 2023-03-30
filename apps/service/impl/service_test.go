package impl_test

import (
	"os"
	"testing"

	// 注册所有服务
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/test/tools"
)

func TestCreateService(t *testing.T) {
	req := service.NewCreateServiceRequest()
	req.Name = "moperator"
	req.Description = "k8s operator"
	req.Owner = "admin"
	app, err := impl.CreateService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(app)
}

func TestQueryGitlabProject(t *testing.T) {
	req := service.NewQueryGitlabProjectRequest()
	req.Address = os.Getenv("GITLAB_ADDRESS")
	req.Token = os.Getenv("GITLAB_PRIVATE_TOKEN")
	set, err := impl.QueryGitlabProject(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestQueryService(t *testing.T) {
	req := service.NewQueryServiceRequest()
	set, err := impl.QueryService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestCreateServiceFromGitLab(t *testing.T) {
	req := service.NewCreateServiceRequest()
	tools.MustReadJsonFile("./test/create_service.json", req)

	app, err := impl.CreateService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(app)
}

func TestDeleteService(t *testing.T) {
	req := service.NewDeleteServiceRequestWithID("c40e9e7d")
	app, err := impl.DeleteService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(app)
}
