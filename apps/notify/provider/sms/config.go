package sms

import (
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/notify/provider/sms/ali"
	"github.com/infraboard/mcenter/apps/notify/provider/sms/tencent"
)

func NewDefaultSMS() *SMS {
	return &SMS{
		Provider:      notify.SMS_PROVIDER_TENCENT,
		TencentConfig: tencent.NewDefaultConfig(),
		AliConfig:     ali.NewDefaultConfig(),
	}
}

type SMS struct {
	// 短信服务商
	Provider notify.SMS_PROVIDER `bson:"Provider" json:"Provider"`
	// 腾讯短信服务配置
	TencentConfig *tencent.Config `bson:"tencent_config" json:"tencent_config"`
	// 阿里云短信服务配置
	AliConfig *ali.Config `bson:"ali_config" json:"ali_config"`
}
