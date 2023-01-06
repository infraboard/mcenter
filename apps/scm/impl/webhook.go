package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/scm"
)

// 应用事件处理
func (s *impl) HandleGitlabEvent(ctx context.Context, in *scm.GitlabWebHookEvent) (
	*scm.GitlabWebHookEvent, error) {
	return nil, nil
}
