package tencent

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
)

const (
	// DEFAULT_TENCENT_SMS_ENDPOINT todo
	DEFAULT_TENCENT_SMS_ENDPOINT = "sms.tencentcloudapi.com"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// LoadConfigFromEnv todo
func LoadConfigFromEnv() (*Config, error) {
	cfg := NewDefaultConfig()
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("load config from env, %s", err.Error())
	}
	return cfg, nil
}

// NewDefaultConfig todo
func NewDefaultConfig() *Config {
	return &Config{}
}

// 接口和相关文档请参考https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2019-07-11&Action=SendSms&SignVersion=
type Config struct {
	Endpoint  string `bson:"endpoint" json:"endpoint" env:"SMS_TENCENT_ENDPOINT"`
	SecretID  string `bson:"secret_id" json:"secret_id" validate:"required,lte=64" env:"SMS_TENCENT_SECRET_ID"`
	SecretKey string `bson:"secret_key" json:"secret_key" validate:"required,lte=64" env:"SMS_TENCENT_SECRET_KEY"`
	AppID     string `bson:"app_id" json:"app_id" validate:"required,lte=64" env:"SMS_TENCENT_APPID"`
	Sign      string `bson:"sign" json:"sign" validate:"required,lte=128" env:"SMS_TENCENT_SIGN"`
}

// Desensitize 脱敏
func (c *Config) Desensitize() {
	c.SecretKey = ""
}

// Validate todo
func (s *Config) Validate() error {
	return validate.Struct(s)
}

// GetEndpoint todo
func (s *Config) GetEndpoint() string {
	if s.Endpoint != "" {
		return s.Endpoint
	}

	return DEFAULT_TENCENT_SMS_ENDPOINT
}
