package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/upsert/#upsert
func (i *impl) upsert(ctx context.Context, ins *instance.Instance) error {
	filter := bson.D{{Key: "_id", Value: ins.Id}}
	update := bson.D{{Key: "$set", Value: ins}}
	opts := options.Update().SetUpsert(true)
	if _, err := i.col.UpdateOne(ctx, filter, update, opts); err != nil {
		return exception.NewInternalServerError("inserted instance(%s) document error, %s",
			ins.FullName(), err)
	}
	return nil
}

func newSerchRequest(r *instance.SearchRequest) *searchRequest {
	return &searchRequest{
		r,
	}
}

type searchRequest struct {
	*instance.SearchRequest
}

func (r *searchRequest) FindOptions() *options.FindOptions {
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

func (r *searchRequest) FindFilter() bson.M {
	filter := bson.M{}
	if r.Domain != "" {
		filter["domain"] = r.Domain
	}
	if r.Namespace != "" {
		filter["namespace"] = r.Namespace
	}
	if r.ApplicationName != "" {
		filter["application_name"] = r.ApplicationName
	}
	if r.Region != "" {
		filter["registry_info.region"] = r.Region
	}
	if r.Environment != "" {
		filter["registry_info.environment"] = r.Environment
	}
	if r.Group != "" {
		filter["registry_info.group"] = r.Group
	}
	return filter
}

func (i *impl) search(ctx context.Context, req *searchRequest) (*instance.InstanceSet, error) {
	resp, err := i.col.Find(ctx, req.FindFilter(), req.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find instance error, error is %s", err)
	}

	set := instance.NewInstanceSet()
	// 循环
	for resp.Next(ctx) {
		ins := instance.NewDefaultInstance()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode instance error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, req.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get instance count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}
