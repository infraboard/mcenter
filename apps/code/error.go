package code

import (
	"github.com/infraboard/mcenter/version"
	"github.com/infraboard/mcube/exception"
)

var (
	ErrUnauthorized = exception.NewAPIException(version.ServiceName, 1401, "验证码校验失败", "Code Header Required, Format: %s: ${code}", CODE_HEADER_KEY)
)
