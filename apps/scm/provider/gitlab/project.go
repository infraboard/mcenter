package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/infraboard/mcenter/apps/scm"
	"github.com/infraboard/mcenter/apps/scm/provider"
)

func NewProjectSet() *scm.ProjectSet {
	return &scm.ProjectSet{
		Items: []*scm.Project{},
	}
}

// https://gitlab.com/api/v4/projects?owned=true
// https://docs.gitlab.com/ce/api/projects.html
func (r *SCM) ListProjects() (*scm.ProjectSet, error) {
	projectURL := r.resourceURL("projects", map[string]string{"owned": "true", "simple": "true"})
	req, err := r.newJSONRequest("GET", projectURL)
	if err != nil {
		return nil, err
	}

	// 发起请求
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取body
	bytesB, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respString := string(bytesB)

	if (resp.StatusCode / 100) != 2 {
		return nil, fmt.Errorf("status code[%d] is not 200, response %s", resp.StatusCode, respString)
	}

	set := NewProjectSet()
	if err := json.Unmarshal(bytesB, &set.Items); err != nil {
		return nil, err
	}

	return set, nil
}

// POST /projects/:id/hooks
// https://docs.gitlab.com/ce/api/projects.html#add-project-hook
func (r *SCM) AddProjectHook(in *provider.AddProjectHookRequest) (*provider.AddProjectHookResponse, error) {
	addHookURL := r.resourceURL(fmt.Sprintf("projects/%d/hooks", in.ProjectID), nil)
	req, err := r.newFormReqeust("POST", addHookURL, strings.NewReader(in.Hook.FormValue().Encode()))
	if err != nil {
		return nil, err
	}

	// 发起请求
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取body
	bytesB, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respString := string(bytesB)

	if (resp.StatusCode / 100) != 2 {
		return nil, fmt.Errorf("status code[%d] is not 200, response %s", resp.StatusCode, respString)
	}

	ins := provider.NewAddProjectHookResponse()
	if err := json.Unmarshal(bytesB, &ins); err != nil {
		return nil, err
	}

	return ins, nil
}

// DELETE /projects/:id/hooks/:hook_id
// https://docs.gitlab.com/ce/api/projects.html#delete-project-hook
func (r *SCM) DeleteProjectHook(in *provider.DeleteProjectReqeust) error {
	addHookURL := r.resourceURL(fmt.Sprintf("projects/%d/hooks/%d", in.ProjectID, in.HookID), nil)
	req, err := r.newFormReqeust("DELETE", addHookURL, nil)
	if err != nil {
		return err
	}

	// 发起请求
	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取body
	bytesB, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respString := string(bytesB)

	if (resp.StatusCode / 100) != 2 {
		return fmt.Errorf("status code[%d] is not 200, response %s", resp.StatusCode, respString)
	}

	return nil
}
