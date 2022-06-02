package client

import "github.com/infraboard/mcenter/client/auth"

// NewDefaultConfig todo
func NewDefaultConfig() *Config {
	return &Config{
		Address: "localhost:18050",
	}
}

// Config 客户端配置
type Config struct {
	Address      string `json:"adress" toml:"adress" yaml:"adress" env:"MCENTER_ADDRESS"`
	ClientID     string `json:"client_id" toml:"client_id" yaml:"client_id" env:"MCENTER_CLINET_ID"`
	ClientSecret string `json:"client_secret" toml:"client_secret" yaml:"client_secret" env:"MCENTER_CLIENT_SECRET"`
}

func (c *Config) Credentials() *auth.Authentication {
	return auth.NewAuthentication(c.ClientID, c.ClientSecret)
}
