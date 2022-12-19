package code

import (
	"fmt"
	"strings"

	"github.com/infraboard/mcenter/apps/notify"
)

// NewDefaultCodeSetting todo
func NewDefaultCodeSetting() *CodeSetting {
	return &CodeSetting{
		NotifyType:    notify.NOTIFY_TYPE_MAIL,
		ExpireMinutes: 10,
		MailTemplate:  "您的动态验证码为：{1}，{2}分钟内有效！，如非本人操作，请忽略本邮件！",
	}
}

type CodeSetting struct {
	// 通知方式
	NotifyType notify.NOTIFY_TYPE `bson:"type" json:"type"`
	// 验证码默认过期时间
	ExpireMinutes uint32 `bson:"expire_minutes" json:"expire_minutes" validate:"required"`
	// 邮件通知时的模板
	MailTemplate string `bson:"mail_template" json:"mail_template"`
	// 短信通知时的云商模板ID
	SmsTemplateID string `bson:"sms_template_id" json:"sms_template_id"`
}

// RenderMailCentent todo
func (c *CodeSetting) RenderMailCentent(code string, expireMinite uint32) string {
	// 如果为0 则使用默认值
	if expireMinite == 0 {
		expireMinite = c.ExpireMinutes
	}

	t1 := strings.ReplaceAll(c.MailTemplate, "{1}", code)
	return strings.ReplaceAll(t1, "{2}", fmt.Sprintf("%d", expireMinite))
}
