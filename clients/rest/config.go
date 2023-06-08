package rest

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

func NewDefaultConfig() *Config {
	return &Config{
		Address:    "http://127.0.0.1:8010",
		PathPrefix: "/mcenter/api/v1/",
	}
}

type Config struct {
	Token      string `json:"token" toml:"token" yaml:"token" env:"MCENTER_TOKEN"`
	Address    string `json:"address" toml:"address" yaml:"address" env:"MCENTER_HTTP_ADDRESS" validate:"required"`
	PathPrefix string `json:"path_prefix" toml:"path_prefix" yaml:"path_prefix" env:"MCENTER_HTTP_PATH_PREFIX" validate:"required"`
}

func (c *Config) Validate() error {
	return validate.Struct(c)
}
