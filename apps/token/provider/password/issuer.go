package password

import (
	"context"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	AUTH_FAILED = exception.NewUnauthorized("user or password not connrect")
)

type issuer struct {
	user   user.Service
	domain domain.Service

	log logger.Logger
}

func (i *issuer) Init() error {
	i.user = app.GetInternalApp(user.AppName).(user.Service)
	i.domain = app.GetInternalApp(domain.AppName).(domain.Service)
	i.log = zap.L().Named("issuer.password")
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_PASSWORD
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	if !req.GrantType.Equal(token.GRANT_TYPE_PASSWORD) {
		return nil, exception.NewBadRequest("password issuer is only for %s", token.GRANT_TYPE_PASSWORD)
	}

	if req.Username == "" || req.Password == "" {
		return nil, AUTH_FAILED
	}

	// 检测用户的密码是否正确
	u, err := i.user.DescribeUser(ctx, user.NewDescriptUserRequestWithName(req.Username))
	if err != nil {
		return nil, err
	}
	if err := u.Password.CheckPassword(req.Password); err != nil {
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
		ps := d.Spec.SecuritySetting.PasswordSecurity
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

	// 3. 颁发Token
	tk := token.NewToken(req)
	tk.Domain = u.Spec.Domain
	tk.Username = u.Spec.Username
	tk.UserType = u.Spec.Type
	tk.UserId = u.Id
	return tk, nil
}

func init() {
	provider.Registe(&issuer{})
}
