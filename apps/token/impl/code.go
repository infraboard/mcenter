package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) IssueCode(ctx context.Context, req *token.IssueCodeRequest) (
	*token.IssueCodeResponse, error) {
	// 获取验证码颁发器
	issuer := provider.GetCodeIssuer(token.GRANT_TYPE(req.GrantType))

	// 确保有provider
	if issuer == nil {
		return nil, exception.NewBadRequest("grant type %s not support", req.GrantType)
	}

	// 颁发验证码
	c, err := issuer.IssueCode(ctx, req)
	if err != nil {
		return nil, err
	}

	// 保存
	if _, err := s.col.InsertOne(ctx, c); err != nil {
		return nil, exception.NewInternalServerError("inserted verify code(%s) document error, %s",
			c, err)
	}

	// 发送验证码
	msg, err := s.send(ctx, c)
	if err != nil {
		return nil, exception.NewInternalServerError("send verify code error, %s", err)
	}

	return token.NewIssueCodeResponse(msg), nil
}

func (s *service) send(ctx context.Context, code *token.Code) (string, error) {

	// 查询domain, 根据系统配置, 给用户发送通知
	dom, err := s.domain.DescribeDomain(ctx, domain.NewDescribeDomainRequestByName(code.Domain))
	if err != nil {
		return "", fmt.Errorf("get user domain error, %s", err)
	}
	conf := dom.Spec.CodeConfig

	var message string
	switch conf.NotifyType {
	case notify.NOTIFY_TYPE_MAIL:
		content := conf.RenderMailCentent(code.Code, code.ExpiredMinite)
		// 邮件通知
		s.log.Debugf("mail to user %s", code.Username)
		req := notify.NewSendMailRequest(
			"验证码",
			content,
			code.Username,
		)
		record, err := s.notify.SendNotify(ctx, req)
		if err != nil {
			return "", fmt.Errorf("send verify code by mail error, %s", err)
		}
		message = fmt.Sprintf("验证码已通过邮件发送到你的邮箱: %s, 请及时查收", record.Targets())
		s.log.Debugf("send verify code to user: %s by mail ok", code.Username)
	case notify.NOTIFY_TYPE_SMS:
		// 短信通知
		s.log.Debugf("sms to user %s", code.Username)
		req := notify.NewSendSMSRequest(
			conf.SmsTemplateId,
			[]string{code.ExpiredMiniteString()},
			code.Username,
		)
		record, err := s.notify.SendNotify(ctx, req)
		if err != nil {
			return "", fmt.Errorf("send verify code by sms error, %s", err)
		}
		message = fmt.Sprintf("验证码已通过短信发送到你的手机: %s, 请及时查收", record.Targets())
		s.log.Debugf("send verify code to user: %s by sms ok", code.Username)
	default:
		return "", fmt.Errorf("unknown notify type %s", conf.NotifyType)
	}

	return message, nil
}

func (s *service) VerifyCode(ctx context.Context, req *token.VerifyCodeRequest) (
	*token.Code, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate check code request error, %s", err)
	}

	code := token.NewDefaultCode()
	if err := s.col.FindOne(ctx, bson.M{"_id": req.HashID()}).Decode(code); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("verify code: %s  not found", req.Code)
		}

		return nil, exception.NewInternalServerError("find system config %s error, %s", req.Code, err)
	}

	// 校验Token是否过期
	if code.IsExpired() {
		return nil, exception.NewPermissionDeny("verify code is expired")
	}

	// 没过去验证成功, 删除
	if err := s.deleteCode(ctx, code); err != nil {
		s.log.Errorf("delete check ok verify code error, %s", err)
	}

	return code, nil
}
