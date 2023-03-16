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

func TestQueryService(t *testing.T) {
	req := service.NewQueryServiceRequest()
	set, err := impl.QueryService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestQueryGitlabProject(t *testing.T) {
	req := service.NewQueryGitlabProjectRequest()
	req.Address = os.Getenv("GITLAB_ADDRESS")
	req.Token = os.Getenv("GITLAB_PRIVATE_TOkEN")
	set, err := impl.QueryGitlabProject(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}
