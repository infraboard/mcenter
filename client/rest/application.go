package rest

import (
	"context"

	"github.com/infraboard/mcenter/apps/application"
	"github.com/infraboard/mcube/client/rest"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/pb/request"
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
	client *rest.RESTClient
}

func (i *appImpl) ValidateCredential(ctx context.Context, req *application.ValidateCredentialRequest) (
	*application.Application, error) {
	ins := application.NewDefaultApplication()
	resp := response.NewData(ins)

	err := i.client.
		Post("application").
		Body(req).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *appImpl) CreateApplication(ctx context.Context, req *application.CreateApplicationRequest) (
	*application.Application, error) {
	ins := application.NewDefaultApplication()
	resp := response.NewData(ins)

	err := i.client.
		Post("application").
		Body(req).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, err
	}
	return ins, nil
}

func (i *appImpl) UpdateApplication(ctx context.Context, req *application.UpdateApplicationRequest) (
	*application.Application, error) {
	ins := application.NewDefaultApplication()
	resp := response.NewData(ins)

	var err error
	switch req.UpdateMode {
	case request.UpdateMode_PUT:
		err = i.client.
			Put("application/" + req.Id).
			Body(req.Spec).
			Do(ctx).
			Into(resp)
	case request.UpdateMode_PATCH:
		err = i.client.
			Patch("application/" + req.Id).
			Body(req.Spec).
			Do(ctx).
			Into(resp)
	}

	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, err
	}

	return ins, nil
}

func (i *appImpl) QueryApplication(ctx context.Context, req *application.QueryApplicationRequest) (
	*application.ApplicationSet, error) {
	set := application.NewApplicationSet()
	resp := response.NewData(set)

	err := i.client.
		Get("application").
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, err
	}

	return set, nil
}

func (i *appImpl) DescribeApplication(ctx context.Context, req *application.DescribeApplicationRequest) (
	*application.Application, error) {
	ins := application.NewDefaultApplication()
	resp := response.NewData(ins)

	err := i.client.
		Get("application/" + req.Id).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, err
	}
	return ins, nil
}

func (i *appImpl) DeleteApplication(ctx context.Context, req *application.DeleteApplicationRequest) (
	*application.Application, error) {
	ins := application.NewDefaultApplication()
	resp := response.NewData(ins)

	err := i.client.
		Delete("application/" + req.Id).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, err
	}
	return ins, nil
}

func (i *appImpl) RefreshCredential(ctx context.Context, req *application.DescribeApplicationRequest) (
	*application.Application, error) {
	return nil, nil
}
