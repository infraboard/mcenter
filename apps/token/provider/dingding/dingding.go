package dingding

import (
	"context"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/client/rest"
)

func NewDingDingClient(conf *domain.DingDingConfig) *DingDing {
	client := rest.NewRESTClient()
	client.SetBaseURL("https://api.dingtalk.com")
	return &DingDing{
		conf: conf,
		rc:   client,
	}
}

// 钉钉客户端
type DingDing struct {
	conf *domain.DingDingConfig
	rc   *rest.RESTClient
}

// 登陆
func (d *DingDing) Login(ctx context.Context, code string) (*user.DingDingAccessToken, error) {
	tk, err := d.GetToken(ctx, code)
	if err != nil {
		return nil, err
	}

	// 设置Token
	d.rc.SetHeader("x-acs-dingtalk-access-token", tk.AccessToken)
	return tk, nil
}

func (d *DingDing) NewGetUserTokenRequest(code string) *GetUserTokenRequest {
	return &GetUserTokenRequest{
		ClientId:     d.conf.ClientId,
		ClientSecret: d.conf.ClientSecret,
		Code:         code,
		GrantType:    "authorization_code",
	}
}

// 获取用户个人token https://open.dingtalk.com/document/orgapp-server/obtain-user-token?spm=ding_open_doc.document.0.0.5cf7722fu9sqAm#doc-api-dingtalk-GetUserToken
func (d *DingDing) GetToken(ctx context.Context, code string) (*user.DingDingAccessToken, error) {
	resp := user.NewDingDingAccessToken()
	err := d.rc.
		Post("/v1.0/oauth2/userAccessToken").
		Header(rest.CONTENT_TYPE_HEADER, restful.MIME_JSON).
		Body(d.NewGetUserTokenRequest(code)).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 获取用户通讯录个人信息 https://open.dingtalk.com/document/orgapp-server/dingtalk-retrieve-user-information?spm=ding_open_doc.document.0.0.5cf7722fu9sqAm#doc-api-dingtalk-GetUser
func (c *DingDing) GetUserInfo(ctx context.Context) (*User, error) {
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
