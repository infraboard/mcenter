package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/v2/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/token"
)

func (i *impl) save(ctx context.Context, ins *service.Service) error {
	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted Service(%s) document error, %s",
			ins.Spec.Name, err)
	}
	return nil
}

func (i *impl) update(ctx context.Context, ins *service.Service) error {
	if _, err := i.col.UpdateByID(ctx, ins.Meta.Id, ins); err != nil {
		return exception.NewInternalServerError("inserted Service(%s) document error, %s",
			ins.Spec.Name, err)
	}

	return nil
}

func newQueryRequest(r *service.QueryServiceRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

type queryRequest struct {
	*service.QueryServiceRequest
}

func (r *queryRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort: bson.D{
			{Key: "create_at", Value: -1},
		},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryRequest) FindFilter() bson.M {
	filter := bson.M{}
	token.MakeMongoFilter(filter, r.Scope)
	policy.MakeMongoFilter(filter, "labels", r.Filters)

	if len(r.RepositorySshUrls) > 0 {
		filter["code_repository.ssh_url"] = bson.M{"$in": r.RepositorySshUrls}
	}

	if r.Keywords != "" {
		filter["name"] = bson.M{"$regex": r.Keywords, "$options": "im"}
	}

	if len(r.Ids) > 0 {
		filter["_id"] = bson.M{"$in": r.Ids}
	}
	return filter
}

func (i *impl) delete(ctx context.Context, ins *service.Service) error {
	if ins == nil || ins.Meta.Id == "" {
		return fmt.Errorf("service is nil")
	}

	result, err := i.col.DeleteOne(ctx, bson.M{"_id": ins.Meta.Id})
	if err != nil {
		return exception.NewInternalServerError("delete Service(%s) error, %s", ins.Meta.Id, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("service %s not found", ins.Meta.Id)
	}

	return nil
}
