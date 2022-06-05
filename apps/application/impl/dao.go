package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/mcenter/apps/application"
)

func (i *impl) save(ctx context.Context, ins *application.Application) error {
	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted application(%s) document error, %s",
			ins.Spec.Name, err)
	}
	return nil
}

func (i *impl) update(ctx context.Context, ins *application.Application) error {
	if _, err := i.col.UpdateByID(ctx, ins.Id, ins); err != nil {
		return exception.NewInternalServerError("inserted application(%s) document error, %s",
			ins.Spec.Name, err)
	}

	return nil
}

func newQueryRequest(r *application.QueryApplicationRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

type queryRequest struct {
	*application.QueryApplicationRequest
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
	// if r.Keywords != "" {
	// 	filter["$or"] = bson.A{
	// 		bson.M{"data.name": bson.M{"$regex": r.Keywords, "$options": "im"}},
	// 		bson.M{"data.author": bson.M{"$regex": r.Keywords, "$options": "im"}},
	// 	}
	// }
	return filter
}

func (i *impl) query(ctx context.Context, req *queryRequest) (*application.ApplicationSet, error) {
	resp, err := i.col.Find(ctx, req.FindFilter(), req.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find book error, error is %s", err)
	}

	ApplicationSet := application.NewApplicationSet()
	// 循环
	for resp.Next(ctx) {
		ins := application.NewDefaultApplication()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode book error, error is %s", err)
		}

		ApplicationSet.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, req.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get application count error, error is %s", err)
	}
	ApplicationSet.Total = count

	return ApplicationSet, nil
}

func (i *impl) get(ctx context.Context, req *application.DescribeApplicationRequest) (*application.Application, error) {
	filter := bson.M{}
	switch req.DescribeBy {
	case application.DescribeBy_APP_ID:
		filter["_id"] = req.Id
	case application.DescribeBy_APP_CLIENT_ID:
		filter["credential.client_id"] = req.ClientId
	case application.DescribeBy_APP_NAME:
		filter["spec.name"] = req.Name
	}

	ins := application.NewDefaultApplication()
	if err := i.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("application %s not found", req)
		}

		return nil, exception.NewInternalServerError("find application %s error, %s", req, err)
	}

	return ins, nil
}

func (i *impl) delete(ctx context.Context, ins *application.Application) error {
	if ins == nil || ins.Id == "" {
		return fmt.Errorf("application is nil")
	}

	result, err := i.col.DeleteOne(ctx, bson.M{"_id": ins.Id})
	if err != nil {
		return exception.NewInternalServerError("delete application(%s) error, %s", ins.Id, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("application %s not found", ins.Id)
	}

	return nil
}
