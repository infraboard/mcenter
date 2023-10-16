package rpc

import (
	"github.com/infraboard/mcube/ioc"
)

const (
	MCENTER = "mcenter_client"
)

func C() *ClientSet {
	return ioc.Config().Get(MCENTER).(*Mcenter).cs
}

func Config() *Mcenter {
	return ioc.Config().Get(MCENTER).(*Mcenter)
}
