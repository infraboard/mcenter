package voice

import (
	"github.com/infraboard/mcenter/apps/notify/provider/voice/ali"
	"github.com/infraboard/mcenter/apps/notify/provider/voice/tencent"
)

func NewDefaultVoiceSetting() *VoiceSetting {
	return &VoiceSetting{
		TencentConfig: tencent.DefaultConfig(),
		AliConfig:     ali.DefaultConfig(),
	}
}

type VoiceSetting struct {
	// 腾讯短信服务配置
	TencentConfig *tencent.Config `bson:"tencent_config" json:"tencent_config"`
	// 阿里云短信服务配置
	AliConfig *ali.Config `bson:"ali_config" json:"ali_config"`
}
