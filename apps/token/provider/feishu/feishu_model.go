package feishu

import (
	"strings"

	"github.com/infraboard/mcenter/apps/user"
)

func NewUser() *User {
	return &User{}
}

type User struct {
	// sub	用户在应用内的唯一标识，等同于open_id	string	是
	Sub string `json:"sub"`
	// name	用户姓名	string	是
	Name string `json:"name"`
	// picture	用户头像，等同于avatar_url	string	是
	Picture string `json:"picture"`
	// open_id	用户在应用内的唯一标识, 等同于sub	string	是
	OpenId string `json:"open_id"`
	// union_id	用户统一ID，在同一租户开发的所有应用内的唯一标识	string	是
	UnionId string `json:"union_id"`
	// en_name	用户英文名称	string	是
	EnName string `json:"en_name"`
	// tenant_key	当前企业标识	string	是
	TenantKey string `json:"tenant_key"`
	// avatar_url	用户头像，等同于picture	string	是
	AvatarUrl string `json:"avatar_url"`
	// avatar_thumb	用户头像 72x72	string	是
	AvatarThumb string `json:"avatar_thumb"`
	// avatar_middle	用户头像 240x240	string	是
	AvatarMiddle string `json:"avatar_middle"`
	// avatar_big	用户头像 640x640	string	是
	AvatarBig string `json:"avatar_big"`
	// user_id	用户 user id，申请了邮箱获取权限(获取用户 user ID)的应用会返回该字段	string	否
	UserId string `json:"user_id"`
	// email	用户邮箱，申请了邮箱获取权限(获取用户邮箱信息)的应用会返回该字段	string	否
	Email string `json:"email"`
	// mobile	用户手机号，申请了手机号获取权限(获取用户手机号)的应用会返回该字段
	Mobile string `json:"mobile"`
}

func (u *User) Username() string {
	if u.Email != "" {
		ud := strings.Split(u.Email, "@")
		if len(ud) == 2 {
			return ud[1]
		}
	}

	if u.EnName != "" {
		return u.EnName
	}

	return u.Name
}

func (u *User) ToProfile() *user.Profile {
	p := user.NewProfile()
	p.RealName = u.Name
	p.Avatar = u.AvatarUrl
	p.Email = u.Email
	p.Phone = u.Mobile
	return p
}

func (u *User) ToCreateUserRequest(domain, password, descriptoin string) *user.CreateUserRequest {
	req := &user.CreateUserRequest{
		Provider:    user.PROVIDER_FEISHU,
		Type:        user.TYPE_SUB,
		CreateFrom:  user.CREATE_FROM_PRIMARY_ACCOUNT,
		Domain:      domain,
		Username:    u.Username(),
		Password:    password,
		Description: descriptoin,
		Feishu: &user.Feishu{
			OpenId:  u.OpenId,
			UnionId: u.UnionId,
			UserId:  u.UserId,
		},
	}
	return req
}
