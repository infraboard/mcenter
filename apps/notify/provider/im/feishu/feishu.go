package feishu

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
	"github.com/chyroc/lark/card"
	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/notify/provider/im"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewFeishuNotifyer(conf *domain.FeishuConfig) im.ImNotifyer {
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

// 发送飞书卡片消息
func (f *Feishu) SendMessage(ctx context.Context, req *im.SendMessageRequest) error {
	// 设置卡片格式消息
	content := card.Card(
		card.Div().SetText(card.MarkdownText(req.Content)),
		card.Note(card.MarkdownText("该消息由用户中心(mcenter)提供")),
	)
	content.SetHeader(card.Header(req.Title).SetTurquoise())

	_, _, err := f.client.Message.Send().ToUserID(req.Uid).SendCard(ctx, content.String())
	if err != nil {
		return fmt.Errorf("send feishu message to user %s error, %s", req.Uid, err)
	}
	return nil
}
