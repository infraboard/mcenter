package wechatwork

import (
	"strings"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/user"
)

// 获取access_token https://developer.work.weixin.qq.com/document/path/91039
type GetAppAccessTokenRequest struct {
	// 企业ID
	CorpId string `json:"corpid"`
	// 应用的凭证密钥，注意应用需要是启用状态
	CorpSecret string `json:"corpsecret"`
}

type User struct {
	// 成员UserID
	UserId string `json:"userid"`
	// 性别。0表示未定义，1表示男性，2表示女性。仅在用户同意snsapi_privateinfo授权时返回真实值，否则返回0.
	Gender int `json:"gender"`
	// 头像url。仅在用户同意snsapi_privateinfo授权时返回
	AvatarUrl string `json:"avatar"`
	// 手机，仅在用户同意snsapi_privateinfo授权时返回，第三方应用不可获取
	Mobile string `json:"mobile"`
	// 邮箱，仅在用户同意snsapi_privateinfo授权时返回，第三方应用不可获取
	Email string `json:"email"`
	// 企业邮箱，仅在用户同意snsapi_privateinfo授权时返回，第三方应用不可获取
	BizMail string `json:"biz_mail"`
	// 仅在用户同意snsapi_privateinfo授权时返回，第三方应用不可获取
	Address string `json:"address"`
	// 员工个人二维码（扫描可添加为外部联系人），仅在用户同意snsapi_privateinfo授权时返回
	QrCode string `json:"qr_code"`
}

func (u *User) Username() string {
	if u.BizMail != "" {
		ud := strings.Split(u.BizMail, "@")
		if len(ud) == 2 {
			return ud[1]
		}
	}

	if u.Email != "" {
		ud := strings.Split(u.Email, "@")
		if len(ud) == 2 {
			return ud[1]
		}
	}

	if u.Mobile != "" {
		return u.Mobile
	}

	return ""
}

func (u *User) ToProfile() *user.Profile {
	p := user.NewProfile()
	p.Avatar = u.AvatarUrl
	p.Email = u.Email
	p.Phone = u.Mobile
	return p
}

func (u *User) ToCreateUserRequest(domain, password, descriptoin string) *user.CreateUserRequest {
	req := &user.CreateUserRequest{
		Provider:    user.PROVIDER_WECHAT_WORK,
		Type:        user.TYPE_SUB,
		CreateBy:    user.CREATE_BY_ADMIN,
		Domain:      domain,
		Username:    u.Username(),
		Password:    password,
		Description: descriptoin,
		Wechatwork: &user.WechatWork{
			UserId: u.UserId,
		},
	}
	return req
}

type Response struct {
	// 返回码
	Code int `json:"errcode"`
	// 对返回码的文本描述内容
	Message string `json:"errmsg"`
}

func NewUserInfoResponse() *UserInfoResponse {
	return &UserInfoResponse{}
}

type UserInfoResponse struct {
	Response
	UserInfo
}

type UserInfo struct {
	// 成员UserID
	UserId string `json:"userid"`
	// 成员票据，最大为512字节，有效期为1800s scope为snsapi_privateinfo，且用户在应用可见范围之内时返回此参数。
	UserTicket string `json:"user_ticket"`
}

func NewUserDetailResponse() *UserResponse {
	return &UserResponse{}
}

type UserResponse struct {
	Response
	User
}

func NewGetTokenResponse() *GetTokenResponse {
	return &GetTokenResponse{}
}

type GetTokenResponse struct {
	Response
	domain.WechatWorkAccessToken
}
