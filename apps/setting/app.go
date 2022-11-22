package setting

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/infraboard/mcenter/apps/notify"
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
			Type:  notify.NOTIFY_TYPE_MAIL,
			Email: mail.NewDefaultConfig(),
			SMS:   sms.NewDefaultSMS(),
			Code:  NewDefaultCode(),
		},
	}
}

// Setting 系统配置
type Setting struct {
	// 配置版本
	Version string `bson:"_id" json:"version"`
	// 系统通知
	Notify *Notify `bson:"verify" json:"verify"`
}

func (s *Setting) JsonFormat() string {
	d, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(d)
}

type Notify struct {
	// 通知方式
	Type notify.NOTIFY_TYPE `bson:"type" json:"type"`
	// 邮件通知配置
	Email *mail.Config `bson:"email" json:"email"`
	// 短信通知配置
	SMS *sms.SMS `bson:"sms" json:"sms"`
	// 验证码配置
	Code *Code `bson:"code" json:"code"`
}

// NewDefaultConfig todo
func NewDefaultCode() *Code {
	return &Code{
		ExpireMinutes: 10,
		MailTemplate:  "您的动态验证码为：{1}，{2}分钟内有效！，如非本人操作，请忽略本邮件！",
	}
}

type Code struct {
	// 验证码默认过期时间
	ExpireMinutes uint32 `bson:"expire_minutes" json:"expire_minutes" validate:"required"`
	// 邮件通知时的模板
	MailTemplate string `bson:"mail_template" json:"mail_template"`
	// 短信通知时的云商模板ID
	SmsTemplateID string `bson:"sms_template_id" json:"sms_template_id"`
}

// RenderMailCentent todo
func (c *Code) RenderMailCentent(code string, expireMinite uint32) string {
	// 如果为0 则使用默认值
	if expireMinite == 0 {
		expireMinite = c.ExpireMinutes
	}

	t1 := strings.ReplaceAll(c.MailTemplate, "{1}", code)
	return strings.ReplaceAll(t1, "{2}", fmt.Sprintf("%d", expireMinite))
}
