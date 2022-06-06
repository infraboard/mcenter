package rest

import "github.com/go-resty/resty/v2"

func NewClient(conf *Config) *ClientSet {
	c := resty.New()
	c.SetAuthToken(conf.Token)
	c.SetBaseURL(conf.Address)
	return &ClientSet{
		c: c,
	}
}

type ClientSet struct {
	c *resty.Client
}

func (c *ClientSet) Application() ApplicationService {
	return &appImpl{client: c.c}
}

func (c *ClientSet) Instance() InstanceService {
	return &insImpl{client: c.c}
}
