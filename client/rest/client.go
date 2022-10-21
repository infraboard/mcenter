package rest

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/mcube/client/rest"
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

func LoadClientFromEnv() error {
	conf := NewDefaultConfig()
	err := env.Parse(conf)
	if err != nil {
		return err
	}
	fmt.Println(conf.Address, conf.PathPrefix)

	c, err := NewClient(conf)
	if err != nil {
		return err
	}
	client = c
	return nil
}

func NewClient(conf *Config) (*ClientSet, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	c := rest.NewRESTClient()
	c.SetBearerTokenAuth(conf.Token)
	c.SetBaseURL(conf.Address + conf.PathPrefix)
	return &ClientSet{
		c: c,
	}, nil
}

type ClientSet struct {
	c *rest.RESTClient
}

func (c *ClientSet) Service() MetaService {
	return &svcImpl{client: c.c}
}

func (c *ClientSet) Instance() InstanceService {
	return &insImpl{client: c.c}
}

func (c *ClientSet) Token() TokenService {
	return &tokenImpl{client: c.c}
}
