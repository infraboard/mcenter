package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/exception"
)

func (s *service) IssueCode(ctx context.Context, req *code.IssueCodeRequest) (
	*code.IssueCodeResponse, error) {
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
		fmt.Println(tk)
	default:
		return nil, exception.NewBadRequest("uknown issue_by %s", req.IssueBy)
	}

	// 生成验证码

	// 发送验证码

	return nil, nil
}

func (s *service) VerifyCode(ctx context.Context, req *code.VerifyCodeRequest) (
	*code.Code, error) {
	return nil, nil
}
