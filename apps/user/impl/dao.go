package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) save(ctx context.Context, u *user.User) error {
	if _, err := s.col.InsertOne(context.TODO(), u); err != nil {
		return exception.NewInternalServerError("inserted user(%s) document error, %s",
			u.Id, err)
	}

	return nil
}

func (s *service) update(ctx context.Context, ins *user.User) error {
	if _, err := s.col.UpdateByID(ctx, ins.Id, ins); err != nil {
		return exception.NewInternalServerError("inserted user(%s) document error, %s",
			ins.Id, err)
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
