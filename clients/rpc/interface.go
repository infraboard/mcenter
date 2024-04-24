package rpc

import (
	"github.com/infraboard/mcube/v2/ioc"
)

const (
	AppName = "mcenter"
)

func C() *ClientSet {
	return ioc.Config().Get(AppName).(*Mcenter).cs
}

func Config() *Mcenter {
	return ioc.Config().Get(AppName).(*Mcenter)
}
