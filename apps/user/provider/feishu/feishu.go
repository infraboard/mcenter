package feishu

import (
	"context"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcube/logger"
	lark "github.com/larksuite/oapi-sdk-go/v3"
)

type Feishu struct {
	conf   *domain.FeishuConfig
	client *lark.Client
	log    logger.Logger
}

func (f *Feishu) ListUser(ctx context.Context) {
	f.client.Contact.User.List(ctx, nil)
}
