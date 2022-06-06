package rest

import (
	"context"

	"github.com/go-resty/resty/v2"
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

type appImpl struct {
	client *resty.Client
}

func (i *appImpl) ValidateCredential(ctx context.Context, req *application.ValidateCredentialRequest) (
	*application.Application, error) {
	return nil, nil
}

func (i *appImpl) CreateApplication(ctx context.Context, req *application.CreateApplicationRequest) (
	*application.Application, error) {
	return nil, nil
}

func (i *appImpl) UpdateApplication(ctx context.Context, req *application.UpdateApplicationRequest) (
	*application.Application, error) {
	return nil, nil
}

func (i *appImpl) QueryApplication(ctx context.Context, req *application.QueryApplicationRequest) (
	*application.ApplicationSet, error) {
	i.client.R()
	return nil, nil
}

func (i *appImpl) DescribeApplication(ctx context.Context, req *application.DescribeApplicationRequest) (
	*application.Application, error) {
	return nil, nil
}

func (i *appImpl) DeleteApplication(ctx context.Context, req *application.DeleteApplicationRequest) (
	*application.Application, error) {
	return nil, nil
}

func (i *appImpl) RefreshCredential(ctx context.Context, req *application.DescribeApplicationRequest) (
	*application.Application, error) {
	return nil, nil
}
