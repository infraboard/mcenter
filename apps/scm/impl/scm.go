package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/scm"
)

// 查询项目列表
func (s *impl) QueryProject(ctx context.Context, in *scm.QueryProjectRequest) (
	*scm.ProjectSet, error) {
	return nil, nil
}

// 应用事件处理
func (s *impl) HandleEvent(ctx context.Context, in *scm.WebHookEvent) (
	*scm.WebHookEvent, error) {
	return nil, nil
}
