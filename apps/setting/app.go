package setting

import (
	"encoding/json"

	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcenter/apps/notify/provider/mail"
	"github.com/infraboard/mcenter/apps/notify/provider/sms"
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
			SMS:   sms.NewDefaultSMS(),
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
	SMS *sms.SMS `bson:"sms" json:"sms"`
}
