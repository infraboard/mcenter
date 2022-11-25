package ldap

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/domain/password"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/exception"
)

type issuer struct {
	domain domain.Service
	user   user.Service
}

func (i *issuer) Init() error {
	i.domain = app.GetInternalApp(domain.AppName).(domain.Service)
	i.user = app.GetInternalApp(user.AppName).(user.Service)
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_LDAP
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	// 从用户名中 获取到DN, 比如oldfish@devcloud.io, 比如username: oldfish domain: devcloud.io
	username, ldapSuffix := user.SpliteUserAndDomain(req.Username)

	// 查询域下 对应的ldap设置
	dom, err := i.domain.DescribeDomain(ctx, domain.NewDescribeDomainRequestByName(ldapSuffix))
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
	u, err := p.CheckUserPassword(username, req.Password)
	if err != nil {
		return nil, err
	}

	// 判断用户是否在数据库存在, 如果不存在需要同步到本地数据库
	lu, err := i.user.DescribeUser(ctx, user.NewDescriptUserRequestWithName(u.Username))
	if err != nil {
		if exception.IsNotFoundError(err) {
			gen := password.New(dom.Spec.SecuritySetting.PasswordSecurity)
			randomPass, err := gen.Generate()
			if err != nil {
				return nil, err
			}
			// 创建本地用户
			newReq := user.NewLDAPCreateUserRequest(dom.Spec.Name, u.Username, *randomPass, "系统自动生成")
			lu, err = i.user.CreateUser(ctx, newReq)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// 颁发Token
	tk := token.NewToken(req)
	tk.Domain = lu.Spec.Domain
	tk.Username = lu.Spec.Username
	tk.UserType = lu.Spec.Type
	tk.UserId = lu.Id

	return nil, nil
}

func init() {
	provider.Registe(&issuer{})
}
