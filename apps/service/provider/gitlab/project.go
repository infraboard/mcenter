package gitlab

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/client/negotiator"
	"github.com/infraboard/mcube/client/rest"
)

func newProjectV4(gitlab *GitlabV4) *ProjectV4 {
	return &ProjectV4{
		client: gitlab.client.Group("projects"),
	}
}

type ProjectV4 struct {
	client *rest.RESTClient
}

// 参考文档: https://docs.gitlab.com/ce/api/projects.html
func (p *ProjectV4) ListProjects(ctx context.Context) (*ProjectSet, error) {
	set := NewProjectSet()

	err := p.client.
		Get("/").
		Param("owned", "true").
		Param("simple", "true").
		Do(ctx).
		Into(&set.Items)

	if err != nil {
		return nil, err
	}
	return set, nil
}

// 参考文档: https://docs.gitlab.com/ce/api/projects.html#add-project-hook
func (p *ProjectV4) AddProjectHook(ctx context.Context, req *AddProjectHookRequest) (
	*AddProjectHookResponse, error) {
	ins := NewAddProjectHookResponse()
	err := p.client.
		Post(fmt.Sprintf("%d/hooks", req.ProjectID)).
		Header(rest.CONTENT_TYPE_HEADER, string(negotiator.MIME_POST_FORM)).
		Body(req.WebHook.FormValue()).
		Do(ctx).
		Into(ins)

	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 参考文档: https://docs.gitlab.com/ce/api/projects.html#delete-project-hook
func (p *ProjectV4) DeleteProjectHook(ctx context.Context, req *DeleteProjectHookReqeust) error {
	err := p.client.
		Delete(fmt.Sprintf("%d/hooks/%d", req.ProjectID, req.HookID)).
		Do(ctx).
		Error()

	if err != nil {
		return err
	}
	return nil
}
