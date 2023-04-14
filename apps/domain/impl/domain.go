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

// 查询域列表
func (s *service) QueryDoamin(ctx context.Context, in *domain.QueryDomainRequest) (*domain.DomainSet, error) {
	r := newQueryRequest(in)
	resp, err := s.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find user error, error is %s", err)
	}

	set := domain.NewDomainSet()
	// 循环
	for resp.Next(ctx) {
		ins := domain.NewDefaultDomain()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode user error, error is %s", err)
		}
		ins.Desensitize()
		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get user count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (s *service) CreateDomain(ctx context.Context, req *domain.CreateDomainRequest) (*domain.Domain, error) {
	d, err := domain.New(req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	if _, err := s.col.InsertOne(ctx, d); err != nil {
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
		filter["spec.name"] = req.Name
	}

	d := domain.NewDefaultDomain()
	if err := s.col.FindOne(ctx, filter).Decode(d); err != nil {
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

	var descReq *domain.DescribeDomainRequest
	if req.Name != "" {
		descReq = domain.NewDescribeDomainRequestByName(req.Name)
	}
	if req.Id != "" {
		descReq = domain.NewDescribeDomainRequestById(req.Id)
	}

	d, err := s.DescribeDomain(ctx, descReq)
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
		if err := d.Spec.Validate(); err != nil {
			return nil, err
		}
	default:
		return nil, exception.NewBadRequest("unknown update mode: %s", req.UpdateMode)
	}

	d.Meta.UpdateAt = time.Now().Unix()
	_, err = s.col.UpdateOne(ctx, bson.M{"_id": d.Meta.Id}, bson.M{"$set": d})
	if err != nil {
		return nil, exception.NewInternalServerError("update domain(%s) error, %s", d.Meta.Id, err)
	}

	return d, nil
}
