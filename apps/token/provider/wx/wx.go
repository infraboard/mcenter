package wx

import (
	"context"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/client/rest"
)

func NewWechatWorkClient(conf *domain.WechatWorkConfig) *WechatWork {
	client := rest.NewRESTClient()
	client.SetBaseURL("https://qyapi.weixin.qq.com/cgi-bin")
	return &WechatWork{
		conf: conf,
		rc:   client,
	}
}

// 钉钉客户端
type WechatWork struct {
	conf *domain.WechatWorkConfig
	rc   *rest.RESTClient
}

// 登陆
func (d *WechatWork) Login(ctx context.Context, code string) (*user.DingDingAccessToken, error) {
	tk, err := d.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	// 设置Token
	return tk, nil
}

// 获取应用access_token https://developer.work.weixin.qq.com/document/path/91039
func (d *WechatWork) GetToken(ctx context.Context) (*user.DingDingAccessToken, error) {
	resp := user.NewDingDingAccessToken()
	err := d.rc.
		Post("/v1.0/oauth2/userAccessToken").
		Header(rest.CONTENT_TYPE_HEADER, restful.MIME_JSON).
		Body(nil).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 获取访问用户敏感信息 https://developer.work.weixin.qq.com/document/path/95833
func (c *WechatWork) GetUserInfo(ctx context.Context) (*User, error) {
	resp := NewUser()
	err := c.rc.
		Get("/v1.0/contact/users/me").
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
