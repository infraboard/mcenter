package feishu

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
	log    logger.Logger
}

func (i *issuer) Init() error {
	i.domain = app.GetInternalApp(domain.AppName).(domain.Service)
	i.log = zap.L().Named("issuer.feishu")
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_FEISHU
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	// 从用户名中 获取到DN, 比如oldfish@default, 比如username: oldfish domain: default
	_, domainName := user.SpliteUserAndDomain(req.Username)

	// 查询域下 对应的飞书设置
	dom, err := i.domain.DescribeDomain(ctx, domain.NewDescribeDomainRequestByName(domainName))
	if err != nil {
		return nil, err
	}

	if dom.Spec.FeishuSetting == nil {
		return nil, fmt.Errorf("domain feishu not settting")
	}

	// 获取Token
	client := NewFeishuClient(dom.Spec.FeishuSetting)
	if err := client.Login(ctx, req.AuthCode); err != nil {
		return nil, err
	}

	// 获取用户信息
	fu, err := client.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	// 同步飞书用户
	fmt.Println(fu)

	return nil, nil
}

func init() {
	provider.Registe(&issuer{})
}
