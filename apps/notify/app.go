package notify

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	AppName = "notify"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewSendSMSRequest todo
func NewSendSMSRequest() *SendSMSRequest {
	return &SendSMSRequest{}
}

// AddParams todo
func (req *SendSMSRequest) AddParams(params ...string) {
	req.TemplateParams = append(req.TemplateParams, params...)
}

// AddPhone todo
func (req *SendSMSRequest) AddPhone(params ...string) {
	req.PhoneNumbers = append(req.PhoneNumbers, params...)
}

// Validate todo
func (req *SendSMSRequest) Validate() error {
	return validate.Struct(req)
}

// InjectDefaultIsoCode todo
func (req *SendSMSRequest) InjectDefaultIsoCode() {
	for i, number := range req.PhoneNumbers {
		if !strings.HasPrefix(number, "+") {
			req.PhoneNumbers[i] = "+86" + number
		}
	}
}
