package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/notify/provider/mail"
	"github.com/infraboard/mcenter/apps/notify/provider/sms"
	"github.com/infraboard/mcenter/apps/notify/provider/sms/tencent"
	"github.com/infraboard/mcenter/apps/user"
)

// 发送短信
func (s *service) SendMail(ctx context.Context, req *notify.SendMailRequest) (*notify.SendMailResponse, error) {
	conf, err := s.setting.GetSetting(ctx)
	if err != nil {
		return nil, err
	}

	// 查询用户的详情, 获取邮箱, 电话等这些通知信息
	mails := []string{}
	sender := mail.NewSender(conf.Notify.Email)
	for i := range req.Users {
		u, err := s.user.DescribeUser(ctx, user.NewDescriptUserRequestWithName(req.Users[i]))
		if err != nil {
			return nil, fmt.Errorf("get user error, %s", err)
		}
		if u.Profile.Email != "" {
			mails = append(mails, u.Profile.Email)
		}
	}

	sendReq := mail.NewSendMailRequest(mails, req.Title, req.Content)
	if err := sender.Send(ctx, sendReq); err != nil {
		return nil, err
	}

	return &notify.SendMailResponse{SuccessedMails: mails}, nil
}

// 短信通知
func (s *service) SendSMS(ctx context.Context, req *notify.SendSMSRequest) (*notify.SendSmsResponse, error) {

	conf, err := s.setting.GetSetting(ctx)
	if err != nil {
		return nil, err
	}

	ss := conf.Notify.SMS

	sendReq := sms.NewSendSMSRequest()
	sendReq.TemplateId = req.TemplateId
	sendReq.TemplateParams = req.TemplateParams
	for i := range req.Users {
		// 查询用户的详情, 获取邮箱, 电话等这些通知信息
		u, err := s.user.DescribeUser(ctx, user.NewDescriptUserRequestWithName(req.Users[i]))
		if err != nil {
			return nil, fmt.Errorf("get user error, %s", err)
		}
		if u.Profile.Phone != "" {
			sendReq.AddPhone(u.Profile.Phone)
		}
	}

	switch ss.Provider {
	case notify.SMS_PROVIDER_TENCENT:
		sender, err := tencent.NewSender(ss.TencentConfig)
		if err != nil {
			return nil, err
		}
		if err := sender.Send(ctx, sendReq); err != nil {
			return nil, err
		}
	case notify.SMS_PROVIDER_ALI:
		return nil, fmt.Errorf("not impl")
	default:
		return nil, fmt.Errorf("unknow provier: %s", ss.Provider)
	}

	return nil, nil
}

// 语音通知
func (s *service) SendVoice(ctx context.Context, req *notify.SendVoiceRequest) (*notify.SendVoiceResponse, error) {
	conf, err := s.setting.GetSetting(ctx)
	if err != nil {
		return nil, err
	}

	s.log.Debug(conf)
	return &notify.SendVoiceResponse{}, nil
}

// 发送IM消息
func (s *service) SendIM(ctx context.Context, req *notify.SendIMRequest) (*notify.SendImResponse, error) {
	conf, err := s.setting.GetSetting(ctx)
	if err != nil {
		return nil, err
	}
	s.log.Debug(conf)
	return nil, nil
}
