package impl

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/infraboard/mcenter/apps/application"
)

func (i *impl) ValidateCredential(ctx context.Context, req *application.ValidateCredentialRequest) (
	*application.Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCredential not implemented")
}

func (i *impl) CreateService(ctx context.Context, req *application.CreateMicroRequest) (
	*application.Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}

func (i *impl) QueryService(ctx context.Context, req *application.QueryMicroRequest) (
	*application.Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryService not implemented")
}

func (i *impl) DescribeService(ctx context.Context, req *application.DescribeMicroRequest) (
	*application.Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeService not implemented")
}

func (i *impl) DeleteService(ctx context.Context, req *application.DeleteMicroRequest) (
	*application.Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}

func (i *impl) RefreshCredential(ctx context.Context, req *application.DescribeMicroRequest) (
	*application.Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshCredential not implemented")
}
