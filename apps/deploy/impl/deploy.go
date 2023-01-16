package impl

import (
	"context"
	"time"

	"github.com/imdario/mergo"
	"github.com/infraboard/mcenter/apps/deploy"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (i *impl) DescribeDeploy(ctx context.Context, req *deploy.DescribeDeployRequest) (*deploy.Deploy, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	filter := bson.M{}
	switch req.DescribeBy {
	case deploy.DESCRIBE_BY_ID:
		filter["_id"] = req.DescribeValue
	}

	d := deploy.NewDefaultDeploy()
	if err := i.col.FindOne(ctx, filter).Decode(d); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("deploy %s not found", req)
		}

		return nil, exception.NewInternalServerError("find deploy %s error, %s", req.DescribeValue, err)
	}

	return d, nil
}

func (i *impl) UpdateDeploy(ctx context.Context, in *deploy.UpdateDeployRequest) (
	*deploy.Deploy, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := deploy.NewDescribeDeployRequest(in.Id)
	d, err := i.DescribeDeploy(ctx, req)
	if err != nil {
		return nil, err
	}

	switch in.UpdateMode {
	case request.UpdateMode_PUT:
		d.Spec = in.Spec
	case request.UpdateMode_PATCH:
		if err := mergo.MergeWithOverwrite(d.Spec, in.Spec); err != nil {
			return nil, err
		}
		if err := d.Spec.Validate(); err != nil {
			return nil, err
		}
	default:
		return nil, exception.NewBadRequest("unknown update mode: %s", in.UpdateMode)
	}

	d.UpdateAt = time.Now().UnixMilli()
	_, err = i.col.UpdateOne(ctx, bson.M{"_id": d.Id}, bson.M{"$set": d})
	if err != nil {
		return nil, exception.NewInternalServerError("update deploy(%s) error, %s", d.Id, err)
	}

	return d, nil
}

func (i *impl) DeleteDeploy(ctx context.Context, in *deploy.DeleteDeployRequest) (
	*deploy.Deploy, error) {

	req := deploy.NewDescribeDeployRequest(in.Id)
	d, err := i.DescribeDeploy(ctx, req)
	if err != nil {
		return nil, err
	}

	_, err = i.col.DeleteOne(ctx, bson.M{"_id": in.Id})
	if err != nil {
		return nil, exception.NewInternalServerError("delete deploy(%s) error, %s", in.Id, err)
	}

	return d, nil
}
