package feishu

import (
	"context"
	"net/url"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcube/client/rest"
)

func NewFeishuClient(conf *domain.FeishuConfig) *Feishu {
	client := rest.NewRESTClient()
	client.SetBaseURL("https://passport.feishu.cn/suite/passport/oauth")
	return &Feishu{
		conf: conf,
		rc:   client,
	}
}

// 飞书客户端
type Feishu struct {
	conf           *domain.FeishuConfig
	rc             *rest.RESTClient
	isRefreshToken bool
}

func (c *Feishu) IsRefreshToken() bool {
	return c.isRefreshToken
}

// 登陆
func (c *Feishu) Login(ctx context.Context, code string) error {
	if c.conf.Token.IsExpired() {
		tk, err := c.GetToken(ctx, code)
		if err != nil {
			return err
		}
		c.conf.Token = tk
		c.isRefreshToken = true
	}

	// 设置Token
	c.rc.SetBearerTokenAuth(c.conf.Token.AccessToken)
	return nil
}

// 获取token https://open.feishu.cn/document/common-capabilities/sso/api/get-access_token
func (c *Feishu) GetToken(ctx context.Context, code string) (*domain.FeishuAccessToken, error) {
	resp := domain.NewFeishuAccessToken()
	err := c.rc.
		Post("token").
		Header(rest.CONTENT_TYPE_HEADER, "application/x-www-form-urlencoded").
		Body(c.conf.MakeGetTokenFormRequest(code).Encode()).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 获取用户信息 https://open.feishu.cn/document/common-capabilities/sso/api/get-user-info
func (c *Feishu) GetUserInfo(ctx context.Context) (*User, error) {
	resp := NewUser()
	err := c.rc.
		Get("userinfo").
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 刷新已过期的 access_token https://open.feishu.cn/document/common-capabilities/sso/api/refresh-access_token
func (c *Feishu) RefreshToken(ctx context.Context, refreshToken string) (*domain.FeishuAccessToken, error) {
	form := make(url.Values)
	form.Add("grant_type", "refresh_token")
	form.Add("refresh_token", refreshToken)

	resp := domain.NewFeishuAccessToken()
	err := c.rc.
		Post("token").
		Header(rest.CONTENT_TYPE_HEADER, "application/x-www-form-urlencoded").
		Body(form.Encode()).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
