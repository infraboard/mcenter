package impl

import (
	"context"
	"time"

	"github.com/imdario/mergo"
	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreateDomain(ctx context.Context, req *domain.CreateDomainRequest) (*domain.Domain, error) {
	d, err := domain.New(req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	if _, err := s.col.InsertOne(context.TODO(), d); err != nil {
		return nil, exception.NewInternalServerError("inserted a domain document error, %s", err)
	}

	return d, nil
}

func (s *service) DescribeDomain(ctx context.Context, req *domain.DescribeDomainRequest) (*domain.Domain, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	filter := bson.M{}
	switch req.DescribeBy {
	case domain.DESCRIBE_BY_ID:
		filter["_id"] = req.Id
	case domain.DESCRIBE_BY_NAME:
		filter["name"] = req.Name
	}

	d := domain.NewDefault()
	if err := s.col.FindOne(context.TODO(), filter).Decode(d); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("domain %s not found", req)
		}

		return nil, exception.NewInternalServerError("find domain %s error, %s", req.Id, err)
	}

	return d, nil
}

func (s *service) UpdateDomain(ctx context.Context, req *domain.UpdateDomainRequest) (*domain.Domain, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	d, err := s.DescribeDomain(ctx, domain.NewDescribeDomainRequest(req.Id))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case request.UpdateMode_PUT:
		d.Spec = req.Spec
	case request.UpdateMode_PATCH:
		if err := mergo.MergeWithOverwrite(d.Spec, req.Spec); err != nil {
			return nil, err
		}
	default:
		return nil, exception.NewBadRequest("unknown update mode: %s", req.UpdateMode)
	}

	d.UpdateAt = time.Now().UnixMilli()
	_, err = s.col.UpdateOne(context.TODO(), bson.M{"_id": d.Id}, bson.M{"$set": d})
	if err != nil {
		return nil, exception.NewInternalServerError("update domain(%s) error, %s", d.Id, err)
	}

	return d, nil
}
