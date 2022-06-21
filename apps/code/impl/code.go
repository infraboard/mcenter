package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcenter/apps/setting"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/exception"
)

func (s *service) IssueCode(ctx context.Context, req *code.IssueCodeRequest) (
	*code.IssueCodeResponse, error) {
	// 生成验证码
	c, err := code.NewCode(req)
	if err != nil {
		return nil, err
	}

	// 校验凭证合法性
	switch req.IssueBy {
	case code.ISSUE_BY_PASSWORD:
		u, err := s.user.DescribeUser(ctx, user.NewDescriptUserRequestWithName(req.Username))
		if err != nil {
			return nil, err
		}
		if err := u.Password.CheckPassword(req.Password); err != nil {
			return nil, err
		}
	case code.ISSUE_BY_ACCESS_TOKEN:
		tk, err := s.token.ValidateToken(ctx, token.NewValidateTokenRequestByAccessToken(req.AccessToken))
		if err != nil {
			return nil, err
		}
		c.Username = tk.Username
	default:
		return nil, exception.NewBadRequest("uknown issue_by %s", req.IssueBy)
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

	return code.NewIssueCodeResponse(msg), nil
}

func (s *service) send(ctx context.Context, code *code.Code) (string, error) {
	// 查询用户的详情, 获取邮箱, 电话等这些通知信息
	u, err := s.user.DescribeUser(ctx, user.NewDescriptUserRequestWithName(code.Username))
	if err != nil {
		return "", fmt.Errorf("get user error, %s", err)
	}

	// 根据系统配置, 给用户发送通知
	system, err := s.setting.GetSetting()
	if err != nil {
		return "", fmt.Errorf("query system setting error, %s", err)
	}

	switch system.Notify.Type {
	case setting.NOTIFY_TYPE_MAIL:
		s.log.Debugf("mail to user %s", u.Profile.Email)
	case setting.NOTIFY_TYPE_SMS:
		s.log.Debugf("sms to user %s", u.Profile.Phone)
	default:
		return "", fmt.Errorf("unknown notify type %s", system.Notify.Type)
	}

	return "", nil
}

func (s *service) VerifyCode(ctx context.Context, req *code.VerifyCodeRequest) (
	*code.Code, error) {
	return nil, nil
}
