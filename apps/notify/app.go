package notify

const (
	AppName = "notify"
)

// NewSendMailRequest todo
func NewSendMailRequest(users []string, title, content string) *SendMailRequest {
	return &SendMailRequest{
		Users:   users,
		Title:   title,
		Content: content,
	}
}

// NewSendSMSRequest todo
func NewSendSMSRequest() *SendSMSRequest {
	return &SendSMSRequest{}
}

// AddParams todo
func (req *SendSMSRequest) AddParams(params ...string) {
	req.TemplateParams = append(req.TemplateParams, params...)
}

// AddPhone todo
func (req *SendSMSRequest) AddUser(users ...string) {
	req.Users = append(req.Users, users...)
}
