package ldap

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/logger"
	"github.com/rs/zerolog"
)

type issuer struct {
	domain domain.Service
	user   user.Service

	log *zerolog.Logger
}

func (i *issuer) Init() error {
	i.domain = ioc.GetController(domain.AppName).(domain.Service)
	i.user = ioc.GetController(user.AppName).(user.Service)
	i.log = logger.Sub("issuer.ldap")
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_LDAP
}

func (i *issuer) validate(ctx context.Context, username, pass string) (*user.User, error) {
	// 从用户名中 获取到DN, 比如oldfish@default, 比如username: oldfish domain: default
	username, domainName := user.SpliteUserAndDomain(username)

	// 查询域下 对应的ldap设置
	dom, err := i.domain.DescribeDomain(ctx, domain.NewDescribeDomainRequestByName(domainName))
	if err != nil {
		return nil, err
	}

	if dom.Spec.LdapSetting == nil {
		return nil, fmt.Errorf("domain ldap not settting")
	}

	// 连接Ldap Server
	p := NewProvider(dom.Spec.LdapSetting)
	err = p.CheckConnect()
	if err != nil {
		return nil, err
	}

	// 检查用户密码是否正确
	u, err := p.CheckUserPassword(username, pass)
	if err != nil {
		return nil, err
	}

	// 判断用户是否在数据库存在, 如果不存在需要同步到本地数据库
	lu, err := i.user.DescribeUser(ctx, user.NewDescriptUserRequestByName(u.Username))
	if err != nil {
		if exception.IsNotFoundError(err) {
			i.log.Debug().Msgf("sync user: %s(%s) to db", u.Username, dom.Spec.Name)
			// 创建本地用户, 密码等同于LDAP密码
			newReq := user.NewLDAPCreateUserRequest(dom.Spec.Name, u.Username, pass, "系统自动生成")
			lu, err = i.user.CreateUser(ctx, newReq)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return lu, nil
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	u, err := i.validate(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// 颁发Token
	tk := token.NewToken(req)
	tk.Domain = u.Spec.Domain
	tk.Username = u.Spec.Username
	tk.UserType = u.Spec.Type
	tk.UserId = u.Meta.Id
	return nil, nil
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
