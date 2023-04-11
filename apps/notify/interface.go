package notify

const (
	AppName = "notify"
)

type Service interface {
	RPCServer
}

// NewSendMailRequest todo
func NewSendNotifyRequest() *SendNotifyRequest {
	return &SendNotifyRequest{
		NotifyTye:    NOTIFY_TYPE_MAIL,
		Users:        []string{},
		SmsRequest:   NewSMSRequest(),
		VoiceRequest: NewVoiceRequest(),
	}
}

// NewSendSMSRequest todo
func NewSendMailRequest(title, content string, users ...string) *SendNotifyRequest {
	req := NewSendNotifyRequest()
	req.AddUser(users...)
	req.Title = title
	req.Content = content
	return req
}

// NewSendSMSRequest todo
func NewSendSMSRequest(templateId string, templateParams []string, users ...string) *SendNotifyRequest {
	req := NewSendNotifyRequest()
	req.NotifyTye = NOTIFY_TYPE_SMS
	req.SmsRequest.TemplateId = templateId
	req.AddSmsParams(templateParams...)
	req.AddUser(users...)
	return req
}

// AddParams todo
func (req *SendNotifyRequest) AddSmsParams(params ...string) {
	req.SmsRequest.TemplateParams = append(req.SmsRequest.TemplateParams, params...)
}

// AddPhone todo
func (req *SendNotifyRequest) AddUser(users ...string) {
	req.Users = append(req.Users, users...)
}

func NewSMSRequest() *SMSRequest {
	return &SMSRequest{
		TemplateParams: []string{},
	}
}

func NewVoiceRequest() *VoiceRequest {
	return &VoiceRequest{
		TemplateParams: []string{},
	}
}

func NewRecord(req *SendNotifyRequest) *Record {
	return &Record{
		Request:  req,
		Response: []*SendResponse{},
	}
}

func (r *Record) AddResponse(items ...*SendResponse) {
	r.Response = append(r.Response, items...)
}

func (r *Record) Targets() (targets []string) {
	if r.Response == nil {
		return
	}
	for i := range r.Response {
		item := r.Response[i]
		targets = append(targets, item.Target)
	}
	return
}

func NewSendResponse(t string) *SendResponse {
	return &SendResponse{
		Target: t,
		Extra:  map[string]string{},
	}
}

func (r *SendResponse) SendError(err error) {
	r.Message = err.Error()
}

func (r *SendResponse) SendSuccess() {
	r.Success = true
}
