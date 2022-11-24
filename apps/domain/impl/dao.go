package impl

import (
	"github.com/infraboard/mcenter/apps/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryRequest(r *domain.QueryDomainRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

type queryRequest struct {
	*domain.QueryDomainRequest
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

	if len(r.Names) > 0 {
		filter["spec.name"] = bson.M{"$in": r.Names}
	}

	return filter
}
