package feishu

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewFeishuNotifyer(conf *domain.FeishuConfig) *Feishu {
	return &Feishu{
		conf:   conf,
		client: lark.New(lark.WithAppCredential(conf.AppId, conf.AppSecret)),
		log:    zap.L().Named("notifyer.feishu"),
	}
}

type Feishu struct {
	conf   *domain.FeishuConfig
	client *lark.Lark
	log    logger.Logger
}

func (f *Feishu) Send(ctx context.Context) {
	data, resp, err := f.client.Message.Send().ToUserID("uid").SendCard(ctx, "")
	if err != nil {
		f.log.Errorf("send feishu message %s to user %s error, %s")
	}
	fmt.Println(data, resp)
}
