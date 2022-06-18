package setting

import (
	"github.com/infraboard/mcenter/apps/notify/provider/mail"
	"github.com/infraboard/mcenter/apps/notify/provider/sms/ali"
	"github.com/infraboard/mcenter/apps/notify/provider/sms/tencent"
)

const (
	AppName = "setting"
)

// Setting 系统配置
type Setting struct {
	// 配置版本
	Version string `bson:"_id" json:"version"`
	// 系统通知
	Notify *Notify `bson:"verify" json:"verify"`
}

type Notify struct {
	// 通知方式
	Type NOTIFY_TYPE `bson:"type" json:"type"`
	// 邮件通知配置
	Email *mail.Config `bson:"email" json:"email"`
	// 短信通知配置
	SMS *SMS `bson:"sms" json:"sms"`
	// 验证码配置
	Code *Code `bson:"code" json:"code"`
}

type Code struct {
	// 验证码默认过期时间
	ExpireMinutes uint `bson:"expire_minutes" json:"expire_minutes" validate:"required"`
	// 邮件通知时的模板
	MailTemplate string `bson:"mail_template" json:"mail_template"`
	// 短信通知时的云商模板ID
	SmsTemplateID string `bson:"sms_template_id" json:"sms_template_id"`
}

type SMS struct {
	// 短信服务商
	Provider SMS_PROVIDER `bson:"Provider" json:"Provider"`
	// 腾讯短信服务配置
	TencentConfig *tencent.Config `bson:"tencent_config" json:"tencent_config"`
	// 阿里云短信服务配置
	AliConfig *ali.Config `bson:"ali_config" json:"ali_config"`
}
