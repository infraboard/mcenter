package rest

import (
	"github.com/caarlos0/env/v6"
)

var (
	client *ClientSet
)

func C() *ClientSet {
	if client == nil {
		panic("mcenter rest client config not load")
	}
	return client
}

// 通过环境变量加载rest客户端
func LoadClientFromEnv() error {
	conf := NewDefaultConfig()
	err := env.Parse(conf)
	if err != nil {
		return err
	}
	return LoadClientFromConfig(conf)
}

// 通过配置对象加载rest客户端
func LoadClientFromConfig(conf *Config) error {
	c, err := NewClient(conf)
	if err != nil {
		return err
	}
	client = c
	return nil
}
