package rest

import (
	"context"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcube/client/rest"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/pb/request"
)

type MetaService interface {
	ValidateCredential(context.Context, *service.ValidateCredentialRequest) (*service.Service, error)
	CreateService(context.Context, *service.CreateServiceRequest) (*service.Service, error)
	UpdateService(context.Context, *service.UpdateServiceRequest) (*service.Service, error)
	QueryService(context.Context, *service.QueryServiceRequest) (*service.ServiceSet, error)
	DescribeService(context.Context, *service.DescribeServiceRequest) (*service.Service, error)
	DeleteService(context.Context, *service.DeleteServiceRequest) (*service.Service, error)
	RefreshCredential(context.Context, *service.DescribeServiceRequest) (*service.Service, error)
}

type appImpl struct {
	client *rest.RESTClient
}

func (i *appImpl) ValidateCredential(ctx context.Context, req *service.ValidateCredentialRequest) (
	*service.Service, error) {
	ins := service.NewDefaultService()
	resp := response.NewData(ins)

	err := i.client.
		Post("service").
		Body(req).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *appImpl) CreateService(ctx context.Context, req *service.CreateServiceRequest) (
	*service.Service, error) {
	ins := service.NewDefaultService()
	resp := response.NewData(ins)

	err := i.client.
		Post("service").
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

func (i *appImpl) UpdateService(ctx context.Context, req *service.UpdateServiceRequest) (
	*service.Service, error) {
	ins := service.NewDefaultService()
	resp := response.NewData(ins)

	var err error
	switch req.UpdateMode {
	case request.UpdateMode_PUT:
		err = i.client.
			Put("service/" + req.Id).
			Body(req.Spec).
			Do(ctx).
			Into(resp)
	case request.UpdateMode_PATCH:
		err = i.client.
			Patch("service/" + req.Id).
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

func (i *appImpl) QueryService(ctx context.Context, req *service.QueryServiceRequest) (
	*service.ServiceSet, error) {
	set := service.NewServiceSet()
	resp := response.NewData(set)

	err := i.client.
		Get("service").
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

func (i *appImpl) DescribeService(ctx context.Context, req *service.DescribeServiceRequest) (
	*service.Service, error) {
	ins := service.NewDefaultService()
	resp := response.NewData(ins)

	err := i.client.
		Get("service/" + req.Id).
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

func (i *appImpl) DeleteService(ctx context.Context, req *service.DeleteServiceRequest) (
	*service.Service, error) {
	ins := service.NewDefaultService()
	resp := response.NewData(ins)

	err := i.client.
		Delete("service/" + req.Id).
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

func (i *appImpl) RefreshCredential(ctx context.Context, req *service.DescribeServiceRequest) (
	*service.Service, error) {
	return nil, nil
}
