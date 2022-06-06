package rest

type Config struct {
	Token   string `json:"token" toml:"token" yaml:"token" env:"KEYAUTH_TOKEN"`
	Address string
}
