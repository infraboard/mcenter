package notify

func NewMailConfig() *MailConfig {
	return &MailConfig{
		Enabled: false,
	}
}
