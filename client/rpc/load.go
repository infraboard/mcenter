package rpc

import "github.com/caarlos0/env/v6"

var (
	client *ClientSet
)

func C() *ClientSet {
	if client == nil {
		panic("mcenter rpc client config not load")
	}
	return client
}

func LoadClientFromEnv() error {
	conf := NewDefaultConfig()
	err := env.Parse(conf)
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
