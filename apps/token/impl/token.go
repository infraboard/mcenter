package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcube/exception"
)

func (s *service) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (
	*token.Token, error) {
	// 登陆前安全检查
	if err := s.BeforeLoginSecurityCheck(ctx, req); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	// 颁发令牌
	tk, err := s.IssueTokenNow(ctx, req)
	if err != nil {
		return nil, err
	}

	// 登陆后安全检查
	if err := s.AfterLoginSecurityCheck(ctx, req.VerifyCode, tk); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return tk, nil
}

func (s *service) IssueTokenNow(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	// 获取令牌颁发器
	issuer := provider.Get(token.GRANT_TYPE_PASSWORD)
	tk, err := issuer.IssueToken(ctx, req)
	if err != nil {
		return nil, err
	}

	// 入库保存
	if !req.DryRun {
		if err := s.save(ctx, tk); err != nil {
			return nil, err
		}
	}

	return tk, nil
}

func (s *service) BeforeLoginSecurityCheck(ctx context.Context, req *token.IssueTokenRequest) error {
	// 连续登录失败检测
	if err := s.checker.MaxFailedRetryCheck(ctx, req); err != nil {
		return exception.NewBadRequest("%s", err)
	}

	// IP保护检测
	err := s.checker.IPProtectCheck(ctx, req)
	if err != nil {
		return err
	}

	s.log.Debug("security check complete")
	return nil
}

func (s *service) AfterLoginSecurityCheck(ctx context.Context, verifyCode string, tk *token.Token) error {
	// 如果有校验码, 则直接通过校验码检测用户身份安全
	if verifyCode != "" {
		s.log.Debugf("verify code provided, check code ...")
		_, err := s.code.VerifyCode(ctx, code.NewVerifyCodeRequest(tk.Username, verifyCode))
		if err != nil {
			return exception.NewPermissionDeny("verify code invalidate, error, %s", err)
		}
		s.log.Debugf("verfiy code check passed")
		return nil
	}

	// 异地登录检测
	err := s.checker.OtherPlaceLoggedInChecK(ctx, tk)
	if err != nil {
		return exception.NewVerifyCodeRequiredError("异地登录检测失败: %s", err)
	}

	// 长时间未登录检测
	err = s.checker.NotLoginDaysChecK(ctx, tk)
	if err != nil {
		return exception.NewVerifyCodeRequiredError("长时间未登录检测失败: %s", err)
	}

	return nil
}

// 撤销Token
func (s *service) RevolkToken(ctx context.Context, req *token.RevolkTokenRequest) (
	*token.Token, error) {
	tk, err := s.get(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	if tk.RefreshToken != req.RefreshToken {
		return nil, exception.NewBadRequest("refresh token not connrect")
	}

	if err := s.delete(ctx, tk); err != nil {
		return nil, err
	}
	return tk, nil
}

// 切换Token空间
func (s *service) ChangeNamespace(ctx context.Context, req *token.ChangeNamespaceRequest) (
	*token.Token, error) {
	return nil, nil
}

// 校验Token
func (s *service) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (
	*token.Token, error) {
	return nil, nil
}

// 查询Token, 用于查询Token颁发记录, 也就是登陆日志
func (s *service) QueryToken(ctx context.Context, req *token.QueryTokenRequest) (*token.TokenSet, error) {
	return nil, nil
}
