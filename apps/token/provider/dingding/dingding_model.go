package dingding

import (
	"strings"

	"github.com/infraboard/mcenter/apps/user"
)

type GetUserTokenRequest struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Code         string `json:"code"`
	RefreshToken string `json:"refreshToken"`
	GrantType    string `json:"grantType"`
}

func NewUser() *User {
	return &User{}
}

type User struct {
	// 用户的钉钉昵称
	Nick string `json:"nick"`
	// 头像URL
	AvatarUrl string `json:"avatarUrl"`
	// 用户的手机号 如果要获取用户手机号，需要在开发者后台申请个人手机号信息权限
	Mobile string `json:"mobile"`
	// 用户的openId
	OpenId string `json:"openId"`
	// 用户的unionId
	UnionId string `json:"unionId"`
	// 用户的个人邮箱
	Email string `json:"email"`
	// 手机号对应的国家号
	StateCode string `json:"手机号对应的国家号"`
}

func (u *User) Username() string {
	if u.Email != "" {
		ud := strings.Split(u.Email, "@")
		if len(ud) == 2 {
			return ud[1]
		}
	}

	if u.Mobile != "" {
		return u.Mobile
	}

	return u.Nick
}

func (u *User) ToProfile() *user.Profile {
	p := user.NewProfile()
	p.RealName = u.Nick
	p.Avatar = u.AvatarUrl
	p.Email = u.Email
	p.Phone = u.Mobile
	return p
}

func (u *User) ToCreateUserRequest(domain, password, descriptoin string) *user.CreateUserRequest {
	req := &user.CreateUserRequest{
		Provider:    user.PROVIDER_DINGDING,
		Type:        user.TYPE_SUB,
		CreateFrom:  user.CREATE_FROM_PRIMARY_ACCOUNT,
		Domain:      domain,
		Username:    u.Username(),
		Password:    password,
		Description: descriptoin,
		Dingding: &user.DingDing{
			OpenId:  u.OpenId,
			UnionId: u.UnionId,
		},
	}
	return req
}
