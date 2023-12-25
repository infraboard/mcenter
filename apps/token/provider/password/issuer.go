package password

import (
	"context"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
)

var (
	AUTH_FAILED = exception.NewUnauthorized("user or password not connrect")
)

type issuer struct {
	user   user.Service
	domain domain.Service

	log *zerolog.Logger
}

func (i *issuer) Init() error {
	i.user = ioc.Controller().Get(user.AppName).(user.Service)
	i.domain = ioc.Controller().Get(domain.AppName).(domain.Service)
	i.log = log.Sub("issuer.password")
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_PASSWORD
}

func (i *issuer) validate(ctx context.Context, username, pass string) (*user.User, error) {
	if username == "" || pass == "" {
		return nil, AUTH_FAILED
	}

	// 检测用户的密码是否正确
	u, err := i.user.DescribeUser(ctx, user.NewDescriptUserRequestByName(username))
	if err != nil {
		return nil, err
	}
	if err := u.Password.CheckPassword(pass); err != nil {
		return nil, AUTH_FAILED
	}

	// 检测密码是否过期
	var expiredRemain, expiredDays uint
	switch u.Spec.Type {
	case user.TYPE_SUB:
		// 子账号过期策略
		d, err := i.domain.DescribeDomain(ctx, domain.NewDescribeDomainRequestByName(u.Spec.Domain))
		if err != nil {
			return nil, err
		}
		ps := d.Spec.PasswordConfig
		expiredRemain, expiredDays = uint(ps.BeforeExpiredRemindDays), uint(ps.PasswordExpiredDays)
	default:
		// 主账号和管理员密码过期策略
		expiredRemain, expiredDays = uint(u.Password.ExpiredRemind), uint(u.Password.ExpiredDays)
	}

	// 检查密码是否过期
	err = u.Password.CheckPasswordExpired(expiredRemain, expiredDays)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	u, err := i.validate(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// 3. 颁发Token
	tk := token.NewToken(req)
	tk.Domain = u.Spec.Domain
	tk.Username = u.Spec.Username
	tk.UserType = u.Spec.Type
	tk.UserId = u.Meta.Id
	return tk, nil
}

func (i *issuer) IssueCode(ctx context.Context, req *token.IssueCodeRequest) (*token.Code, error) {
	_, err := i.validate(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// 颁发Token
	c, err := token.NewCode(req)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func init() {
	provider.Registe(&issuer{})
}
