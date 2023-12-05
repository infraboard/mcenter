package feishu

import (
	"context"

	"github.com/infraboard/mcenter/apps/domain"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"github.com/rs/zerolog"
)

type Feishu struct {
	conf   *domain.FeishuConfig
	client *lark.Client
	log    *zerolog.Logger
}

func (f *Feishu) ListUser(ctx context.Context) {
	f.client.Contact.User.List(ctx, nil)
}
