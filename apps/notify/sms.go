package notify

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/mcube/v2/ioc/config/validator"
)

const (
	// DEFAULT_TENCENT_SMS_ENDPOINT todo
	DEFAULT_TENCENT_SMS_ENDPOINT = "sms.tencentcloudapi.com"
)

func NewSmsConfig() *SmsConfig {
	return &SmsConfig{
		Enabled:  false,
		Provider: PROVIDER_TENCENT,
		Tencent:  NewTencentSmsConfig(),
	}
}

// LoadTencentSmsConfigFromEnv todo
func LoadTencentSmsConfigFromEnv() (*TencentSmsConfig, error) {
	cfg := NewTencentSmsConfig()
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("load config from env, %s", err.Error())
	}
	return cfg, nil
}

// NewTencentSmsConfig todo
func NewTencentSmsConfig() *TencentSmsConfig {
	return &TencentSmsConfig{
		Endpoint: DEFAULT_TENCENT_SMS_ENDPOINT,
	}
}

// Desensitize 脱敏
func (c *TencentSmsConfig) Desensitize() {
	c.SecretKey = ""
}

// Validate todo
func (s *TencentSmsConfig) Validate() error {
	return validator.Validate(s)
}

// GetEndpoint todo
func (s *TencentSmsConfig) GetEndpointWithDefault() string {
	if s.Endpoint != "" {
		return s.Endpoint
	}

	return DEFAULT_TENCENT_SMS_ENDPOINT
}
