package rpc

import (
	"github.com/caarlos0/env/v6"
)

var (
	client *ClientSet
)

func C() *ClientSet {
	if client == nil {
		panic("mcenter rpc client config not load")
	}
	return client
}

func HasLoaded() bool {
	return client != nil
}

func NewConfigFromEnv() (*Config, error) {
	conf := NewDefaultConfig()
	err := env.Parse(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func LoadClientFromEnv() error {
	conf, err := NewConfigFromEnv()
	if err != nil {
		return err
	}

	return LoadClientFromConfig(conf)
}

func LoadClientFromConfig(conf *Config) error {
	c, err := NewClient(conf)
	if err != nil {
		return err
	}
	client = c
	return nil
}
