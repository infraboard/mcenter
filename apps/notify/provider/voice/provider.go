package voice

import (
	"context"

	"github.com/infraboard/mcenter/apps/notify"
)

var (
	// m is a map from scheme to notifyer.
	m = make(map[string]Notifyer)
)

type Notifyer interface {
	Init() error
	Vendor() string
	VoiceNotifyer
}

type VoiceNotifyer interface {
	Call(context.Context, *notify.SendVoiceRequest) (*notify.SendVoiceResponse, error)
}

// 注册令牌颁发器
func Registe(i Notifyer) {
	m[i.Vendor()] = i
}

func GetVoiceNotifyer(vendor string) VoiceNotifyer {
	if v, ok := m[vendor]; ok {
		return v
	}

	return nil
}
