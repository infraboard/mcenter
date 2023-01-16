package impl

import (
	"github.com/infraboard/mcenter/apps/deploy"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryRequest(r *deploy.QueryDeployRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

type queryRequest struct {
	*deploy.QueryDeployRequest
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

	if len(r.Ids) > 0 {
		filter["_id"] = bson.M{"$in": r.Ids}
	}

	return filter
}
