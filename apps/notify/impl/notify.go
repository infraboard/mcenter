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
	"github.com/infraboard/mcenter/apps/notify/provider/voice"
	vtencent "github.com/infraboard/mcenter/apps/notify/provider/voice/tencent"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/exception"
)

// 邮件通知
func (s *service) SendNotify(ctx context.Context, req *notify.SendNotifyRequest) (*notify.Record, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate SendNotifyRequest error, %s", err)
	}

	// 查询domain
	dom, err := s.domain.DescribeDomain(ctx, domain.NewDescribeDomainRequestByName(req.Domain))
	if err != nil {
		return nil, fmt.Errorf("get user domain error, %s", err)
	}

	r := notify.NewRecord(req)
	mailSendReq := mail.NewSendMailRequest(req.Title, req.Content)
	for i := range req.Users {
		u, err := s.user.DescribeUser(ctx, user.NewDescriptUserRequestById(req.Users[i]))
		if err != nil {
			return nil, fmt.Errorf("get user error, %s", err)
		}
		switch req.NotifyTye {
		case notify.NOTIFY_TYPE_MAIL:
			// 添加邮件人
			mailSendReq.AddTo(u.Spec.Profile.Email)
		case notify.NOTIFY_TYPE_SMS:
			// 添加电话号码
			sendReq := sms.NewSendSMSRequest()
			sendReq.TemplateId = req.SmsRequest.TemplateId
			sendReq.TemplateParams = req.SmsRequest.TemplateParams
			sendReq.AddPhone(u.Spec.Profile.Phone)
			resp := s.SendSMS(ctx, dom.Spec.NotifyConfig.SmsConfig, sendReq)
			resp.User = u.Spec.Username
			r.AddResponse(resp)
		case notify.NOTIFY_TYPE_VOICE:
			// 添加电话号码
			sendReq := voice.NewSendVoiceRequest(
				u.Spec.Profile.Phone,
				req.VoiceRequest.TemplateId,
				req.VoiceRequest.TemplateParams,
			)
			sendReq.SessionContext = req.SessionContext
			resp := s.SendVoice(ctx, dom.Spec.NotifyConfig.VoiceConfig, sendReq)
			resp.User = u.Spec.Username
			r.AddResponse(resp)
		case notify.NOTIFY_TYPE_IM:
			// 补充IM Id
			sendReq := im.NewSendMessageRequest(req.Title, req.Content, u.Spec.GetFeishuUserId())
			sendReq.ContentType = req.ContentType
			resp := s.SendIM(ctx, dom, sendReq)
			resp.User = u.Spec.Username
			r.AddResponse(resp)
		}
	}

	// 邮件批量通知
	if mailSendReq.HasTo() {
		resp := s.SendMailMail(ctx, dom.Spec.NotifyConfig.MailConfig, mailSendReq)
		r.AddResponse(resp)
	}

	if _, err := s.col.InsertOne(ctx, r); err != nil {
		return nil, exception.NewInternalServerError("inserted a notify record document error, %s", err)
	}
	return r, nil
}

// 邮件通知
func (s *service) SendMailMail(ctx context.Context, conf *notify.MailConfig, req *mail.SendMailRequest) *notify.SendResponse {
	resp := notify.NewSendResponse(req.ToStrings())

	// 发送邮件
	sender := mail.NewSender(conf)
	if err := sender.Send(ctx, req); err != nil {
		resp.SendError(err)
		return resp
	}

	resp.SendSuccess()
	return resp
}

// 短信通知
func (s *service) SendSMS(ctx context.Context, conf *notify.SmsConfig, req *sms.SendSMSRequest) *notify.SendResponse {
	resp := notify.NewSendResponse(req.PhoneNumbersString())

	// 发送短信
	switch conf.Provider {
	case notify.PROVIDER_TENCENT:
		sender, err := tencent.NewSender(conf.Tencent)
		if err != nil {
			resp.SendError(err)
			return resp
		}
		err = sender.Send(ctx, req)
		if err != nil {
			resp.SendError(err)
			return resp
		}
	case notify.PROVIDER_ALI:
		resp.SendError(errors.New("not impl"))
	default:
		resp.SendError(fmt.Errorf("unknow provier: %s", conf.Provider))
	}

	resp.SendSuccess()
	return resp
}

// 语音通知
func (s *service) SendVoice(ctx context.Context, conf *notify.VoiceConfig, req *voice.SendVoiceRequest) *notify.SendResponse {
	resp := notify.NewSendResponse(req.PhoneNumber)

	// 发送短信
	switch conf.Provider {
	case notify.PROVIDER_TENCENT:
		sender, err := vtencent.NewQcloudVoice(conf.Tencent)
		if err != nil {
			resp.SendError(err)
			return resp
		}
		result, err := sender.Call(ctx, req)
		if err != nil {
			resp.SendError(err)
			return resp
		}
		resp.VoiceResponse = result
	case notify.PROVIDER_ALI:
		resp.SendError(errors.New("not impl"))
	default:
		resp.SendError(fmt.Errorf("unknow provier: %s", conf.Provider))
	}

	resp.SendSuccess()
	return resp
}

// 发送IM消息
func (s *service) SendIM(ctx context.Context, dom *domain.Domain, req *im.SendMessageRequest) *notify.SendResponse {
	resp := notify.NewSendResponse(req.Uid)
	if req.Uid == "" {
		resp.SendError(errors.New("feishu uid not found"))
		return resp
	}

	if dom.Spec.FeishuSetting == nil {
		resp.SendError(fmt.Errorf("domain feishu setting is nil"))
		return resp
	}
	notifyer := feishu.NewFeishuNotifyer(dom.Spec.FeishuSetting)

	err := notifyer.SendMessage(ctx, req)
	if err != nil {
		resp.SendError(fmt.Errorf("send msg error, %s", err))
		return resp
	}

	resp.SendSuccess()
	return resp
}

// 查询发送记录
func (s *service) QueryRecord(ctx context.Context, in *notify.QueryRecordRequest) (*notify.RecordSet, error) {
	r := newQueryRequest(in)
	resp, err := s.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find record error, error is %s", err)
	}

	set := notify.NewRecordSet()
	// 循环
	for resp.Next(ctx) {
		ins := notify.NewDefaultRecord()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode user error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get user count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}
