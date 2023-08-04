package token

import (
	"github.com/infraboard/mcenter/version"
	"github.com/infraboard/mcube/exception"
)

var (
	ErrTokenUnauthorized = exception.NewUnauthorized("Auth Header Required, Format: %s: Bearer ${access_token}", ACCESS_TOKEN_HEADER_KEY)
)

var (
	ErrCodeUnauthorized = exception.NewAPIException(version.ServiceName, 1401, "验证码校验失败", "Code Header Required, Format: %s: ${code}", CODE_HEADER_KEY)
)
