package gitlab_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/scm/provider"
	"github.com/infraboard/mcenter/apps/scm/provider/gitlab"
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
	hook := &provider.WebHook{
		PushEvents:          true,
		TagPushEvents:       true,
		MergeRequestsEvents: true,
		Token:               "9999",
		Url:                 "http://www.baidu.com",
	}
	req := provider.NewAddProjectHookRequest(ProjectID, hook)

	resp, err := v4.Project().AddProjectHook(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestDeleteProjectHook(t *testing.T) {
	req := provider.NewDeleteProjectReqeust(ProjectID, 15519532)
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
