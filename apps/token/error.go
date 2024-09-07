package token

import (
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/application"
)

var (
	ErrTokenUnauthorized = exception.NewUnauthorized("Auth Header Required, Format: %s: Bearer ${access_token}", ACCESS_TOKEN_HEADER_KEY)
)

var (
	ErrCodeUnauthorized = exception.NewApiException(1401, "验证码校验失败").
		WithMessagef("Code Header Required, Format: %s: ${code}", CODE_HEADER_KEY).
		WithNamespace(application.Get().AppName)
)
