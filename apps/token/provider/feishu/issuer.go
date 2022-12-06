package feishu

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

	// 如果刷新了Token配置，需要更新域相关配置
	if client.IsRefreshToken() {
		dom.Spec.FeishuSetting.Token = client.conf.Token
		req := domain.NewPatchPomainRequest(dom.Id, dom.Spec)
		_, err := i.domain.UpdateDomain(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	// 获取用户信息
	fu, err := client.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	// 同步飞书用户
	// 判断用户是否在数据库存在, 如果不存在需要同步到本地数据库
	lu, err := i.user.DescribeUser(ctx, user.NewDescriptUserRequestWithName(fu.Name))
	if err != nil {
		if exception.IsNotFoundError(err) {
			i.log.Debugf("sync user: %s(%s) to db", fu.Name, dom.Spec.Name)
			gen := password.New(dom.Spec.SecuritySetting.PasswordSecurity)
			randomPass, err := gen.Generate()
			if err != nil {
				return nil, err
			}
			// 创建本地用户
			newReq := user.NewFeishuCreateUserRequest(dom.Spec.Name, fu.Name, *randomPass, "系统自动生成")
			lu, err = i.user.CreateUser(ctx, newReq)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// 更新用户Profile
	updateReq := user.NewPatchUserRequest(lu.Id)
	updateReq.Profile = fu.ToProfile()
	_, err = i.user.UpdateUser(ctx, updateReq)
	if err != nil {
		return nil, err
	}

	// 颁发Token
	tk := token.NewToken(req)
	tk.Domain = lu.Spec.Domain
	tk.Username = lu.Spec.Username
	tk.UserType = lu.Spec.Type
	tk.UserId = lu.Id
	return tk, nil
}

func init() {
	provider.Registe(&issuer{})
}
