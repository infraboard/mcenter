package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/notify/provider/mail"
	"github.com/infraboard/mcenter/apps/notify/provider/sms/tencent"
)

// 发送短信
func (s *service) SendMail(ctx context.Context, req *notify.SendMailRequest) (*notify.SendResponse, error) {
	conf, err := s.setting.GetSetting(ctx)
	if err != nil {
		return nil, err
	}
	sender := mail.NewSender(conf.Notify.Email)
	if err := sender.Send(req); err != nil {
		return nil, err
	}
	return notify.NewSendResponse(), nil
}

// 邮件通知
func (s *service) SendSMS(ctx context.Context, req *notify.SendSMSRequest) (*notify.SendResponse, error) {
	conf, err := s.setting.GetSetting(ctx)
	if err != nil {
		return nil, err
	}

	sms := conf.Notify.SMS

	switch sms.Provider {
	case notify.SMS_PROVIDER_TENCENT:
		sender, err := tencent.NewSender(sms.TencentConfig)
		if err != nil {
			return nil, err
		}
		if err := sender.Send(ctx, req); err != nil {
			return nil, err
		}
	case notify.SMS_PROVIDER_ALI:
		return nil, fmt.Errorf("not impl")
	default:
		return nil, fmt.Errorf("unknow provier: %s", sms.Provider)
	}

	return notify.NewSendResponse(), nil
}
