package wechatwork

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

type issuer struct {
	domain domain.Service
	user   user.Service
	log    logger.Logger
}

func (i *issuer) Init() error {
	i.domain = app.GetInternalApp(domain.AppName).(domain.Service)
	i.user = app.GetInternalApp(user.AppName).(user.Service)
	i.log = zap.L().Named("issuer.wechat")
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_WECHAT_WORK
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	// 从用户名中 获取到DN, 比如oldfish@default, 比如username: oldfish domain: default
	_, domainName := user.SpliteUserAndDomain(req.Username)

	// 查询域下 对应的飞书设置
	dom, err := i.domain.DescribeDomain(ctx, domain.NewDescribeDomainRequestByName(domainName))
	if err != nil {
		return nil, err
	}

	if dom.Spec.DingdingSetting == nil {
		return nil, fmt.Errorf("domain dingding not settting")
	}

	return nil, nil
}

func init() {
	provider.Registe(&issuer{})
}
