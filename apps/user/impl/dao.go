package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/v2/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) save(ctx context.Context, u *user.User) error {
	if _, err := s.col.InsertOne(ctx, u); err != nil {
		return exception.NewInternalServerError("inserted user(%s) document error, %s",
			u.Meta.Id, err)
	}

	return nil
}

func (s *service) update(ctx context.Context, ins *user.User) error {
	if _, err := s.col.UpdateByID(ctx, ins.Meta.Id, bson.M{"$set": ins}); err != nil {
		return exception.NewInternalServerError("inserted user(%s) document error, %s",
			ins.Meta.Id, err)
	}

	return nil
}

func (s *service) delete(ctx context.Context, set *user.UserSet) error {
	if set == nil || len(set.Items) == 0 {
		return fmt.Errorf("user is nil")
	}

	// 专门优化单个删除
	var result *mongo.DeleteResult
	var err error
	if len(set.Items) == 1 {
		result, err = s.col.DeleteMany(ctx, bson.M{"_id": set.UserIds()[0]})
	} else {
		result, err = s.col.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": set.UserIds()}})
	}

	if err != nil {
		return exception.NewInternalServerError("delete user(%s) error, %s", set, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("user %s not found", set)
	}

	return nil
}

func newQueryRequest(r *user.QueryUserRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

type queryRequest struct {
	*user.QueryUserRequest
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

	if r.Domain != "" {
		filter["domain"] = r.Domain
	}
	if r.Provider != nil {
		filter["provider"] = r.Provider
	}
	if r.Type != nil {
		filter["type"] = r.Type
	}
	if len(r.UserIds) > 0 {
		filter["_id"] = bson.M{"$in": r.UserIds}
	}
	if r.Keywords != "" {
		filter["username"] = bson.M{"$regex": r.Keywords, "$options": "im"}
	}
	if r.Labels != nil {
		for k, v := range r.Labels {
			filter["labels."+k] = v
		}
	}

	if len(r.ExtraUserIds) > 0 {
		filter = bson.M{"$or": bson.A{
			filter,
			bson.M{"_id": bson.M{"$in": r.ExtraUserIds}},
		}}
	}

	return filter
}
