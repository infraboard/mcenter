package ldap

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/app"
)

type issuer struct {
	domain domain.Service
}

func (i *issuer) Init() error {
	i.domain = app.GetInternalApp(domain.AppName).(domain.Service)
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

	p := NewProvider(dom.Spec.LdapSetting)
	err = p.CheckConnect()
	if err != nil {
		return nil, err
	}

	ok, err := p.CheckUserPassword(username, req.Password)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf("auth failed")
	}

	return nil, nil
}

func init() {
	provider.Registe(&issuer{})
}
