package setting

import (
	"encoding/json"

	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/notify/provider/mail"
	sms_ali "github.com/infraboard/mcenter/apps/notify/provider/sms/ali"
	sms_tencent "github.com/infraboard/mcenter/apps/notify/provider/sms/tencent"
	voice_ali "github.com/infraboard/mcenter/apps/notify/provider/voice/ali"
	voice_tencent "github.com/infraboard/mcenter/apps/notify/provider/voice/tencent"
)

const (
	AppName = "setting"
)

const (
	// DEFAULT_CONFIG_VERSION 默认版本
	DEFAULT_CONFIG_VERSION = "v1"
)

// NewDefaultConfig todo
func NewDefaultSetting() *Setting {
	return &Setting{
		Version: DEFAULT_CONFIG_VERSION,
		Notify: &Notify{
			Email: mail.NewDefaultConfig(),
			SMS:   NewDefaultSmsSetting(),
			Voice: NewDefaultVoiceSetting(),
		},
		Code: code.NewDefaultCodeSetting(),
	}
}

// Setting 系统配置
type Setting struct {
	// 配置版本
	Version string `bson:"_id" json:"version"`
	// 系统通知配置
	Notify *Notify `bson:"verify" json:"verify"`
	// 验证码配置
	Code *code.CodeSetting `bson:"code" json:"code"`
}

func (s *Setting) JsonFormat() string {
	d, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(d)
}

type Notify struct {
	// 邮件通知配置
	Email *mail.Config `bson:"email" json:"email"`
	// 短信通知配置
	SMS *SmsSetting `bson:"sms" json:"sms"`
	// 语音通知配置
	Voice *VoiceSetting `bson:"voice" json:"voice"`
}

func NewDefaultSmsSetting() *SmsSetting {
	return &SmsSetting{
		Provider:      notify.SMS_PROVIDER_TENCENT,
		TencentConfig: sms_tencent.NewDefaultConfig(),
		AliConfig:     sms_ali.NewDefaultConfig(),
	}
}

type SmsSetting struct {
	// 短信服务商
	Provider notify.SMS_PROVIDER `bson:"Provider" json:"Provider"`
	// 腾讯短信服务配置
	TencentConfig *sms_tencent.Config `bson:"tencent_config" json:"tencent_config"`
	// 阿里云短信服务配置
	AliConfig *sms_ali.Config `bson:"ali_config" json:"ali_config"`
}

func NewDefaultVoiceSetting() *VoiceSetting {
	return &VoiceSetting{
		TencentConfig: voice_tencent.DefaultConfig(),
		AliConfig:     voice_ali.DefaultConfig(),
	}
}

type VoiceSetting struct {
	// 腾讯短信服务配置
	TencentConfig *voice_tencent.Config `bson:"tencent_config" json:"tencent_config"`
	// 阿里云短信服务配置
	AliConfig *voice_ali.Config `bson:"ali_config" json:"ali_config"`
}
