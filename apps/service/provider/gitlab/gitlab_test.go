package gitlab_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/service/provider/gitlab"
	"github.com/infraboard/mcenter/test/tools"
)

var (
	v4  *gitlab.GitlabV4
	ctx = context.Background()

	ProjectID string = "29032549"
)

func init() {
	tools.DevelopmentSetup()
	conf, err := gitlab.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	v4 = gitlab.NewGitlabV4(conf)
}
