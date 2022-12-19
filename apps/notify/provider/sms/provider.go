package sms

import (
	"context"

	"github.com/infraboard/mcenter/apps/notify"
)

type SmsNotifyer interface {
	Send(context.Context, *notify.SendSMSRequest) error
}
