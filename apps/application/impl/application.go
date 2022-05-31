package impl

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/infraboard/mcenter/apps/application"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
)

func (i *impl) ValidateCredential(ctx context.Context, req *application.ValidateCredentialRequest) (
	*application.Application, error) {
	return nil, nil
}

func (i *impl) CreateApplication(ctx context.Context, req *application.CreateApplicationRequest) (
	*application.Application, error) {
	ins, err := application.NewApplication(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create book error, %s", err)
	}

	if err := i.save(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) UpdateApplication(ctx context.Context, req *application.UpdateApplicationRequest) (
	*application.Application, error) {
	ins, err := i.DescribeApplication(ctx, application.NewDescribeApplicationRequest(req.Id))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case request.UpdateMode_PUT:
		ins.Update(req)
	case request.UpdateMode_PATCH:
		err := ins.Patch(req)
		if err != nil {
			return nil, err
		}
	}

	// 校验更新后数据合法性
	if err := ins.Spec.Validate(); err != nil {
		return nil, err
	}

	if err := i.update(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) QueryApplication(ctx context.Context, req *application.QueryApplicationRequest) (
	*application.ApplicationSet, error) {
	query := newQueryRequest(req)
	return i.query(ctx, query)
}

func (i *impl) DescribeApplication(ctx context.Context, req *application.DescribeApplicationRequest) (
	*application.Application, error) {
	return i.get(ctx, req.Id)
}

func (i *impl) DeleteApplication(ctx context.Context, req *application.DeleteApplicationRequest) (
	*application.Application, error) {
	ins, err := i.DescribeApplication(ctx, application.NewDescribeApplicationRequest(req.Id))
	if err != nil {
		return nil, err
	}

	if err := i.delete(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) RefreshCredential(ctx context.Context, req *application.DescribeApplicationRequest) (
	*application.Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshCredential not implemented")
}
