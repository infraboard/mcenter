package gorestful

import "github.com/infraboard/mcube/v2/ioc"

const (
	AppName = "mcenter_gorestful_auth_middleware"
)

func Get() *HttpAuther {
	return ioc.Config().Get(AppName).(*HttpAuther)
}
