package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/exception"
)

var (
	AUTH_FAILED = exception.NewUnauthorized("user or password not connrect")
)

func (s *service) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (
	*token.Token, error) {
	switch req.GrantType {
	case token.GRANT_TYPE_PASSWORD:
		if req.Username == "" || req.Password == "" {
			return nil, AUTH_FAILED
		}

		// 1. 检测用户的密码是否正确
		u, err := s.user.DescribeUser(ctx, user.NewDescriptUserRequestWithName(req.Username))
		if err != nil {
			return nil, err
		}
		if err := u.Password.CheckPassword(req.Password); err != nil {
			return nil, AUTH_FAILED
		}

		// 2. 检测密码是否过期
		var expiredRemain, expiredDays uint
		switch u.Spec.Type {
		case user.TYPE_SUB:
			// 2.1 子账号过期策略
			d, err := s.domain.DescribeDomain(ctx, domain.NewDescribeDomainRequest(u.Spec.Domain))
			if err != nil {
				return nil, err
			}
			ps := d.Spec.SecuritySetting.PasswordSecurity
			expiredRemain, expiredDays = uint(ps.BeforeExpiredRemindDays), uint(ps.PasswordExpiredDays)
		default:
			// 2.2 主账号和管理员密码过期策略
			expiredRemain, expiredDays = uint(u.Password.ExpiredRemind), uint(u.Password.ExpiredDays)
		}

		err = u.Password.CheckPasswordExpired(expiredRemain, expiredDays)
		if err != nil {
			return nil, err
		}

		// 3. 颁发Token
		tk := token.NewToken(req)
		tk.Domain = u.Spec.Domain
		tk.Username = u.Spec.Username
		tk.UserId = u.Id
		return tk, nil
	default:
		return nil, exception.NewBadRequest("grant type %s not implemented", req.GrantType)
	}
}

// 撤销Token
func (s *service) RevolkToken(ctx context.Context, req *token.RevolkTokenRequest) (
	*token.Token, error) {
	return nil, nil
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
