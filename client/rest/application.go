package rest

import (
	"context"

	"github.com/infraboard/mcenter/apps/application"
)

type ApplicationService interface {
	ValidateCredential(context.Context, *application.ValidateCredentialRequest) (*application.Application, error)
	CreateApplication(context.Context, *application.CreateApplicationRequest) (*application.Application, error)
	UpdateApplication(context.Context, *application.UpdateApplicationRequest) (*application.Application, error)
	QueryApplication(context.Context, *application.QueryApplicationRequest) (*application.ApplicationSet, error)
	DescribeApplication(context.Context, *application.DescribeApplicationRequest) (*application.Application, error)
	DeleteApplication(context.Context, *application.DeleteApplicationRequest) (*application.Application, error)
	RefreshCredential(context.Context, *application.DescribeApplicationRequest) (*application.Application, error)
}
