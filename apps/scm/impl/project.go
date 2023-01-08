package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/scm"
	"github.com/infraboard/mcenter/apps/scm/provider/gitlab"
)

// 查询项目列表
func (s *impl) QueryProject(ctx context.Context, in *scm.QueryProjectRequest) (
	ps *scm.ProjectSet, err error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	switch in.Provider {
	case scm.PROVIDER_GITLAB:
		conf := gitlab.NewDefaultConfig()
		conf.SetAddress(in.Address)
		conf.SetPrivateToken(in.Token)
		v4 := gitlab.NewGitlabV4(conf)
		ps, err = v4.Project().ListProjects(ctx)
	case scm.PROVIDER_GITHUB:
	}

	if err != nil {
		return nil, err
	}
	return ps, nil
}
