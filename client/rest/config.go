package rest

func NewDefaultConfig() *Config {
	return &Config{
		Address:    "127.0.0.1:8010",
		PathPrefix: "/mcenter/api/v1/",
	}
}

type Config struct {
	Token      string `json:"token" toml:"token" yaml:"token" env:"KEYAUTH_TOKEN"`
	Address    string
	PathPrefix string
}
