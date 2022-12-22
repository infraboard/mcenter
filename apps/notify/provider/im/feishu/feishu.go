package feishu

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/notify/provider/im"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func NewFeishuNotifyer(conf *domain.FeishuConfig) im.ImNotifyer {
	return &Feishu{
		conf:   conf,
		client: lark.NewClient(conf.AppId, conf.AppSecret),
		log:    zap.L().Named("notifyer.feishu"),
	}
}

type Feishu struct {
	conf   *domain.FeishuConfig
	client *lark.Client
	log    logger.Logger
}

// 发送飞书卡片消息
// 参考: https://github.com/larksuite/oapi-sdk-go/blob/c4a7ad4662a99731b9d2509f26b0d2de26e1676e/sample/api/im.go
func (f *Feishu) SendMessage(ctx context.Context, req *im.SendMessageRequest) error {
	// 设置卡片格式消息
	div := larkcard.NewMessageCardDiv().
		Text(
			larkcard.NewMessageCardLarkMd().Content(req.Content).Build(),
		).
		Build()

	// 卡片备注
	note := larkcard.NewMessageCardNote().
		Elements(
			[]larkcard.MessageCardNoteElement{
				larkcard.NewMessageCardLarkMd().Content("该消息由用户中心(mcenter)提供").Build(),
			}).
		Build()

	// 卡片头
	header := larkcard.NewMessageCardHeader().
		Template(larkcard.TemplateTurquoise).
		Title(larkcard.NewMessageCardPlainText().Content(req.Title).Build()).
		Build()

	// 组装卡片消息
	content, err := larkcard.NewMessageCard().
		Elements([]larkcard.MessageCardElement{div, note}).
		Header(header).String()
	if err != nil {
		return err
	}

	sendReq := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(larkim.ReceiveIdTypeUserId).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			MsgType(larkim.MsgTypeInteractive).
			ReceiveId(req.Uid).
			Content(content).
			Build()).
		Build()

	resp, err := f.client.Im.Message.Create(ctx, sendReq)
	if err != nil {
		return fmt.Errorf("send feishu message to user %s error, %s", req.Uid, err)
	}

	if !resp.Success() {
		return fmt.Errorf("%d, %s [%s]", resp.Code, resp.Msg, resp.RequestId())
	}

	return nil
}
