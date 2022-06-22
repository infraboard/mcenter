package setting

import "context"

type Service interface {
	GetSetting(context.Context) (*Setting, error)
	UpdateSetting(context.Context, *Setting) (*Setting, error)
}
