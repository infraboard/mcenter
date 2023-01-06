package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/scm"
)

// 应用事件处理
func (s *impl) HandleEvent(ctx context.Context, in *scm.WebHookEvent) (
	*scm.WebHookEvent, error) {
	return nil, nil
}
