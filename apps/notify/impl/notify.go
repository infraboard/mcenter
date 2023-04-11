package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/notify/provider/im"
	"github.com/infraboard/mcenter/apps/notify/provider/im/feishu"
	"github.com/infraboard/mcenter/apps/notify/provider/mail"
	"github.com/infraboard/mcenter/apps/notify/provider/sms"
	"github.com/infraboard/mcenter/apps/notify/provider/sms/tencent"
	"github.com/infraboard/mcenter/apps/user"
)

// 邮件通知
func (s *service) SendMail(ctx context.Context, req *notify.SendNotifyRequest) (*notify.Record, error) {
	r := notify.NewRecord(req)

	// 查询用户邮箱, 构造邮件发送请求
	for i := range req.Users {
		u, err := s.user.DescribeUser(ctx, user.NewDescriptUserRequestWithName(req.Users[i]))
		if err != nil {
			return nil, fmt.Errorf("get user error, %s", err)
		}
		switch req.NotifyTye {
		case notify.NOTIFY_TYPE_MAIL:
			// 查询用户邮箱, 构造邮件发送请求
			sendReq := mail.NewSendMailRequest(req.Title, req.Content, u.Profile.Email)
			r.AddResponse(s.SendMailMail(ctx, sendReq))
		case notify.NOTIFY_TYPE_SMS:
			sendReq := sms.NewSendSMSRequest()
			sendReq.TemplateId = req.SmsRequest.TemplateId
			sendReq.TemplateParams = req.SmsRequest.TemplateParams
			sendReq.AddPhone(u.Profile.Phone)
		case notify.NOTIFY_TYPE_VOICE:
		case notify.NOTIFY_TYPE_IM:
		}
	}
	return nil, nil
}

// 邮件通知
func (s *service) SendMailMail(ctx context.Context, req *mail.SendMailRequest) *notify.SendResponse {
	resp := notify.NewSendResponse(req.ToStrings())
	// 查询系统邮件设置
	conf, err := s.setting.GetSetting(ctx)
	if err != nil {
		resp.SendError(err)
		return resp
	}

	// 发送邮件
	sender := mail.NewSender(conf.Notify.Email)
	if err := sender.Send(ctx, req); err != nil {
		resp.SendError(err)
		return resp
	}

	resp.SendSuccess()
	return resp
}

// 短信通知
func (s *service) SendSMS(ctx context.Context, req *sms.SendSMSRequest) *notify.SendResponse {
	resp := notify.NewSendResponse(req.PhoneNumbersString())

	// 查询系统短信发送设置
	conf, err := s.setting.GetSetting(ctx)
	if err != nil {
		resp.SendError(err)
		return resp
	}

	// 发送短信
	ss := conf.Notify.SMS
	switch ss.Provider {
	case notify.SMS_PROVIDER_TENCENT:
		sender, err := tencent.NewSender(ss.TencentConfig)
		if err != nil {
			resp.SendError(err)
			return resp
		}
		err = sender.Send(ctx, req)
		if err != nil {
			resp.SendError(err)
			return resp
		}
	case notify.SMS_PROVIDER_ALI:
		resp.SendError(errors.New("not impl"))
	default:
		resp.SendError(fmt.Errorf("unknow provier: %s", ss.Provider))
	}

	return resp
}

// 语音通知
func (s *service) SendVoice(ctx context.Context, req *notify.SendNotifyRequest) (*notify.Record, error) {
	conf, err := s.setting.GetSetting(ctx)
	if err != nil {
		return nil, err
	}

	s.log.Debug(conf)
	return nil, nil
}

// 发送IM消息
func (s *service) SendIM(ctx context.Context, req *notify.SendNotifyRequest) (*notify.Record, error) {
	for i := range req.Users {
		u, err := s.user.DescribeUser(ctx, user.NewDescriptUserRequestWithName(req.Users[i]))
		if err != nil {
			return nil, fmt.Errorf("get user error, %s", err)
		}
		if u.Spec.Feishu.UserId == "" {
			return nil, fmt.Errorf("user feishu id not found")
		}
		d, err := s.domain.DescribeDomain(ctx, domain.NewDescribeDomainRequestById(u.Spec.Domain))
		if err != nil {
			return nil, fmt.Errorf("get user domain error, %s", err)
		}
		notifyer := feishu.NewFeishuNotifyer(d.Spec.FeishuSetting)
		msg := im.NewSendMessageRequest(u.Spec.Feishu.UserId, req.Title, req.Content)
		err = notifyer.SendMessage(ctx, msg)
		if err != nil {
			return nil, fmt.Errorf("send msg error, %s", err)
		}
	}

	return nil, nil
}
