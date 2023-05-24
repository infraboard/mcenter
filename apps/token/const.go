package token

const (
	// token默认过期时长
	DEFAULT_ACCESS_TOKEN_EXPIRE_SECOND = 3600
	// 刷新token默认过期时间
	DEFAULT_REFRESH_TOKEN_EXPIRE_SECOND = DEFAULT_ACCESS_TOKEN_EXPIRE_SECOND * 4
)

const (
	// 携带Token的Header 名称, 标准格式: Authorization: bearer xxxx-token
	ACCESS_TOKEN_HEADER_KEY = "Authorization"
	// token也可以放到Cookie中, 只是cookie本身并不是太安全, 如果Authorization头没有才考虑从
	// cookie中获取
	ACCESS_TOKEN_COOKIE_KEY = "mcenter.access_token"
)

const (
	// 权限校验时, 校验访问者的Token的权限时，除了携带自己的身份凭证之外(AccessToken)，
	// 还需要额外携带被校验用户的Token(ValidatedToken)
	VALIDATE_TOKEN_HEADER_KEY = "X-VALIDATE-TOKEN"
)

const (
	TOKEN_ATTRIBUTE_NAME = "token"
)
