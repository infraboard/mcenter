package feishu

import (
	"context"
	"net/url"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/v2/client/rest"
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
	conf *domain.FeishuConfig
	rc   *rest.RESTClient
}

// 登陆
func (c *Feishu) Login(ctx context.Context, code string) (*user.FeishuAccessToken, error) {
	tk, err := c.GetToken(ctx, code)
	if err != nil {
		return nil, err
	}

	// 设置Token
	c.rc.SetBearerTokenAuth(tk.AccessToken)
	return tk, nil
}

// 获取token https://open.feishu.cn/document/common-capabilities/sso/api/get-access_token
func (c *Feishu) GetToken(ctx context.Context, code string) (*user.FeishuAccessToken, error) {
	resp := user.NewFeishuAccessToken()

	err := c.rc.
		Post("token").
		Header(rest.CONTENT_TYPE_HEADER, "application/x-www-form-urlencoded").
		Body(c.conf.MakeGetTokenFormRequest(code)).
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
func (c *Feishu) RefreshToken(ctx context.Context, refreshToken string) (*user.FeishuAccessToken, error) {
	form := make(url.Values)
	form.Add("grant_type", "refresh_token")
	form.Add("refresh_token", refreshToken)

	resp := user.NewFeishuAccessToken()
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
