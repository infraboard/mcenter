package provider

import "github.com/infraboard/mcenter/apps/scm"

var (
	// m is a map from scheme to scm operator.
	m = make(map[scm.PROVIDER]Operator)
)

type Operator interface {
	ListProjects() (*scm.ProjectSet, error)
	AddProjectHook(*AddProjectHookRequest) (*AddProjectHookResponse, error)
	DeleteProjectHook(*DeleteProjectReqeust) error
}
