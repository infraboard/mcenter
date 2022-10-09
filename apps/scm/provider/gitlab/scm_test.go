package gitlab_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/mcenter/apps/scm/provider"
	"github.com/infraboard/mcenter/apps/scm/provider/gitlab"
	"github.com/stretchr/testify/assert"
)

var (
	GitLabAddr    = "https://gitlab.com"
	PraviateToken = ""

	ProjectID int64 = 29032549
)

func TestListProject(t *testing.T) {
	should := assert.New(t)

	repo := gitlab.NewSCM(GitLabAddr, PraviateToken)
	ps, err := repo.ListProjects()
	should.NoError(err)
	fmt.Println(ps)
}

func TestAddProjectHook(t *testing.T) {
	should := assert.New(t)

	repo := gitlab.NewSCM(GitLabAddr, PraviateToken)

	hook := &provider.WebHook{
		PushEvents:          true,
		TagPushEvents:       true,
		MergeRequestsEvents: true,
		Token:               "9999",
		Url:                 "http://www.baidu.com",
	}
	req := provider.NewAddProjectHookRequest(ProjectID, hook)
	ins, err := repo.AddProjectHook(req)
	should.NoError(err)
	fmt.Println(ins)
}

func TestDeleteProjectHook(t *testing.T) {
	should := assert.New(t)

	repo := gitlab.NewSCM(GitLabAddr, PraviateToken)

	req := provider.NewDeleteProjectReqeust(ProjectID, 8439846)
	err := repo.DeleteProjectHook(req)
	should.NoError(err)
}

func init() {
	PraviateToken = os.Getenv("GITLAB_PRIVATE_TOkEN")
}
