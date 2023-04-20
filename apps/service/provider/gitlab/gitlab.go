package gitlab

import (
	"github.com/infraboard/mcube/client/rest"
)

func NewGitlabV4(conf *Config) *GitlabV4 {
	client := rest.NewRESTClient()
	client.EnableTrace()
	client.SetBaseURL(conf.Address + "/api/v4")
	client.SetHeader(rest.CONTENT_TYPE_HEADER, "application/json")
	client.SetHeader("PRIVATE-TOKEN", conf.PrivateToken)

	return &GitlabV4{
		conf:   conf,
		client: client,
	}
}

type GitlabV4 struct {
	conf   *Config
	client *rest.RESTClient
}

func (g *GitlabV4) Project() *ProjectV4 {
	return newProjectV4(g)
}
