package notify

import "github.com/caarlos0/env/v6"

func NewVoiceConfig() *VoiceConfig {
	return &VoiceConfig{
		Provider: PROVIDER_TENCENT,
		Tencent:  DefaultTencentVoiceConfig(),
	}
}

func DefaultTencentVoiceConfig() *TencentVoiceConfig {
	return &TencentVoiceConfig{
		Endpoint:   "vms.tencentcloudapi.com",
		Region:     "ap-guangzhou",
		SignMethod: "TC3-HMAC-SHA256",
		ReqMethod:  "POST",
	}
}

func LoadTencentVoiceConfigFromEnv() (*TencentVoiceConfig, error) {
	conf := DefaultTencentVoiceConfig()
	err := env.Parse(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
