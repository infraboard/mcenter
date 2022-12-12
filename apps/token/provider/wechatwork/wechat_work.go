package wechatwork

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/domain"
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
func (w *WechatWork) Login(ctx context.Context) (*domain.WechatWorkAccessToken, error) {
	err := w.getToken(ctx)
	if err != nil {
		return nil, err
	}

	return w.conf.AccessToken, nil
}

func (w *WechatWork) GetUserInfo(ctx context.Context, code string) (*User, error) {
	ui, err := w.getUserInfo(ctx, code)
	if err != nil {
		return nil, err
	}

	u, err := w.getUserBase(ctx, ui.UserId)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// 获取应用接口凭证 https://developer.work.weixin.qq.com/tutorial/detail/46
func (w *WechatWork) getToken(ctx context.Context) error {
	resp := NewGetTokenResponse()
	err := w.rc.
		Post("gettoken").
		Param("corpid", w.conf.CorpId).
		Param("corpsecret", w.conf.SuitSecret).
		Do(ctx).
		Into(resp)
	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf(resp.Message)
	}

	w.conf.AccessToken = &resp.WechatWorkAccessToken
	return nil
}

// 获取访问用户身份 https://developer.work.weixin.qq.com/document/path/91023
func (w *WechatWork) getUserInfo(ctx context.Context, code string) (*UserInfo, error) {
	resp := NewUserInfoResponse()
	err := w.rc.
		Post("auth/getuserinfo").
		Param("access_token", w.conf.AccessToken.AccessToken).
		Param("code", code).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Message)
	}

	return &resp.UserInfo, nil
}

// 读取成员: https://developer.work.weixin.qq.com/document/path/90196
func (w *WechatWork) getUserBase(ctx context.Context, userId string) (*User, error) {
	resp := NewUserDetailResponse()
	err := w.rc.
		Post("user/get").
		Param("access_token", w.conf.AccessToken.AccessToken).
		Param("userid", userId).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Message)
	}

	return &resp.User, nil
}

// 获取访问用户敏感信息: https://developer.work.weixin.qq.com/document/path/95833
func (w *WechatWork) getUserDetail(ctx context.Context, userTicket string) (*UserDetail, error) {
	resp := NewUserDetailResponse()
	err := w.rc.
		Post("user/getuserdetail").
		Param("access_token", w.conf.AccessToken.AccessToken).
		Param("user_ticket", userTicket).
		Do(ctx).
		Into(&resp.UserDetail)
	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Message)
	}

	return &resp.UserDetail, nil
}
