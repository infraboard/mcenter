package provider

import (
	"context"

	"github.com/infraboard/mcenter/apps/scm"
)

var (
	// m is a map from scheme to scm operator.
	m = make(map[scm.PROVIDER]Operator)
)

type Operator interface {
	ListProjects(context.Context) (*scm.ProjectSet, error)
	AddProjectHook(context.Context, *AddProjectHookRequest) (*AddProjectHookResponse, error)
	DeleteProjectHook(context.Context, *DeleteProjectReqeust) error
}
