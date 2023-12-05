package rpc

import (
	"github.com/infraboard/mcube/v2/ioc"
)

const (
	MCENTER = "mcenter"
)

func C() *ClientSet {
	return ioc.Config().Get(MCENTER).(*Mcenter).cs
}

func Config() *Mcenter {
	return ioc.Config().Get(MCENTER).(*Mcenter)
}
