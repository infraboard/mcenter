package notify

func NewNotifySetting() *NotifySetting {
	return &NotifySetting{
		MailConfig:  NewMailConfig(),
		SmsConfig:   NewSmsConfig(),
		VoiceConfig: NewVoiceConfig(),
	}
}
