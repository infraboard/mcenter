package wechatwork

// 获取access_token https://developer.work.weixin.qq.com/document/path/91039
type GetAppAccessTokenRequest struct {
	// 企业ID
	CorpId string `json:"corpid"`
	// 应用的凭证密钥，注意应用需要是启用状态
	CorpSecret string `json:"corpsecret"`
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
