package gitlab_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/service/provider/gitlab"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	v4  *gitlab.GitlabV4
	ctx = context.Background()

	ProjectID int64 = 29032549
)

func TestListProject(t *testing.T) {
	set, err := v4.Project().ListProjects(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestAddProjectHook(t *testing.T) {
	hook := gitlab.NewGitLabWebHook("9999")
	req := gitlab.NewAddProjectHookRequest(ProjectID, hook)

	resp, err := v4.Project().AddProjectHook(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestDeleteProjectHook(t *testing.T) {
	req := gitlab.NewDeleteProjectHookReqeust(ProjectID, 18877715)
	err := v4.Project().DeleteProjectHook(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func init() {
	zap.DevelopmentSetup()
	conf, err := gitlab.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}
	v4 = gitlab.NewGitlabV4(conf)
}
