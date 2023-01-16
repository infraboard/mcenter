package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/deploy"
	"github.com/infraboard/mcube/exception"
)

func (i *impl) CreateDeploy(ctx context.Context, in *deploy.CreateDeployRequest) (
	*deploy.Deploy, error) {
	ins, err := deploy.New(in)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a deploy document error, %s", err)
	}
	return ins, nil
}

func (i *impl) QueryDeploy(ctx context.Context, in *deploy.QueryDeployRequest) (
	*deploy.DeploySet, error) {
	r := newQueryRequest(in)
	resp, err := i.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find deploy error, error is %s", err)
	}

	set := deploy.NewDeploySet()
	// 循环
	for resp.Next(ctx) {
		ins := deploy.NewDefaultDeploy()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode deploy error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get deploy count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (i *impl) UpdateDeploy(ctx context.Context, in *deploy.CreateDeployRequest) (
	*deploy.Deploy, error) {
	return nil, nil
}

func (i *impl) DeleteDeploy(ctx context.Context, in *deploy.DeleteDeployRequest) (
	*deploy.Deploy, error) {
	return nil, nil
}
