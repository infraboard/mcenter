package gitlab

import (
	"context"
	"strconv"

	"github.com/infraboard/mcube/v2/client/negotiator"
	"github.com/infraboard/mcube/v2/client/rest"
)

func newProjectV4(gitlab *GitlabV4) *ProjectV4 {
	return &ProjectV4{
		client: gitlab.client.Group("projects"),
	}
}

type ProjectV4 struct {
	client *rest.RESTClient
}

// Get a list of all visible projects across GitLab for the authenticated user.
// 参考文档: https://docs.gitlab.com/ee/api/projects.html#list-all-projects
func (p *ProjectV4) ListProjects(ctx context.Context, in *ListProjectRequest) (*ProjectSet, error) {
	set := NewProjectSet()

	var total string
	err := p.client.
		Get("/").
		Param("owned", strconv.FormatBool(in.Owned)).
		Param("simple", strconv.FormatBool(in.Simple)).
		Param("page", in.PageNumerToString()).
		Param("per_page", in.PageSizeToString()).
		Param("order_by", "created_at").
		Param("sort", "desc").
		Param("search", in.Keywords).
		Do(ctx).
		Header(RESPONSE_HEADER_X_TOTAL, &total).
		Into(&set.Items)

	set.SetTotalFromString(total)
	if err != nil {
		return nil, err
	}
	return set, nil
}

// Get languages used in a project with percentage value.
// 参考文档: https://docs.gitlab.com/ee/api/projects.html#languages
func (p *ProjectV4) ListProjectLanguage(ctx context.Context, pid string) (*ProjectLanguageSet, error) {
	resp := map[string]float64{}

	err := p.client.
		Get(pid).
		Suffix("languages").
		Do(ctx).
		Into(&resp)

	if err != nil {
		return nil, err
	}

	set := NewProjectLanguageSet(resp)
	return set, nil
}

// Adds a hook to a specified project.
// 参考文档: https://docs.gitlab.com/ce/api/projects.html#add-project-hook
func (p *ProjectV4) AddProjectHook(ctx context.Context, req *AddProjectHookRequest) (
	*AddProjectHookResponse, error) {
	ins := NewAddProjectHookResponse()
	err := p.client.
		Post(req.ProjectID).
		Suffix("hooks").
		Header(rest.CONTENT_TYPE_HEADER, string(negotiator.MIME_POST_FORM)).
		Body(req.WebHook.FormValue()).
		Do(ctx).
		Into(ins)

	if err != nil {
		return nil, err
	}
	return ins, nil
}

// Removes a hook from a project.
// This is an idempotent method and can be called multiple times. Either the hook is available or not.
// 参考文档: https://docs.gitlab.com/ce/api/projects.html#delete-project-hook
func (p *ProjectV4) DeleteProjectHook(ctx context.Context, req *DeleteProjectHookReqeust) error {
	err := p.client.
		Delete(req.ProjectID).
		Suffix("hooks").
		Suffix(req.HookID).
		Do(ctx).
		Error()

	if err != nil {
		return err
	}
	return nil
}

// Get a list of repository branches from a project, sorted by name alphabetically.
// 参考文档: https://docs.gitlab.com/ee/api/branches.html#list-repository-branches
func (p *ProjectV4) ListProjectBranch(ctx context.Context, in *ListProjectBranchRequest) (*BranchSet, error) {
	set := NewBranchSet()

	var total string
	err := p.client.
		Get(in.ProjectId).
		Suffix("repository").
		Suffix("branches").
		Param("page", in.PageNumerToString()).
		Param("per_page", in.PageSizeToString()).
		Param("search", in.Keywords).
		Do(ctx).
		Header(RESPONSE_HEADER_X_TOTAL, &total).
		Into(&set.Items)

	set.SetTotalFromString(total)

	if err != nil {
		return nil, err
	}

	return set, nil
}

// Get a single project repository branch.
// 参考文档: https://docs.gitlab.com/ee/api/branches.html#get-single-repository-branch
func (p *ProjectV4) GetProjectBranch(ctx context.Context, in *GetProjectBranchRequest) (*Branch, error) {
	ins := NewBranch()
	err := p.client.
		Get(in.ProjectId).
		Suffix("repository").
		Suffix("branches").
		Suffix(in.Branch).
		Do(ctx).
		Into(ins)
	if err != nil {
		return nil, err
	}
	return ins, nil
}
