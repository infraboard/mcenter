package rest

import "github.com/go-resty/resty/v2"

func NewClient() *ClientSet {
	return &ClientSet{
		c: resty.New(),
	}
}

type ClientSet struct {
	c *resty.Client
}

func (c *ClientSet) Instance() {

}
