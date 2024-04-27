package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/v2/exception"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/policy"
)

func (s *impl) CreateNamespace(ctx context.Context, req *namespace.CreateNamespaceRequest) (
	*namespace.Namespace, error) {
	ins, err := namespace.New(req)
	if err != nil {
		return nil, err
	}

	if req.ParentId != "" {
		c, err := s.counter.GetNextSequenceValue(req.ParentId)
		if err != nil {
			return nil, err
		}
		ins.Meta.Id = fmt.Sprintf("%s-%d", req.ParentId, c.Value)
	}

	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted namespace(%s) document error, %s",
			ins.Spec.Name, err)
	}

	return ins, nil
}

func (s *impl) QueryNamespace(ctx context.Context, req *namespace.QueryNamespaceRequest) (
	*namespace.NamespaceSet, error) {
	r := newPaggingQuery(req)
	set := namespace.NewNamespaceSet()

	filter := r.FindFilter()
	s.log.Debug().Msgf("query filter: %s", filter)
	resp, err := s.col.Find(ctx, filter, r.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find namespace error, error is %s", err)
	}

	// 循环
	for resp.Next(ctx) {
		ins := namespace.NewDefaultNamespace()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode namespace error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get namespace count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *impl) DescribeNamespace(ctx context.Context, req *namespace.DescriptNamespaceRequest) (
	*namespace.Namespace, error) {
	r, err := newDescribeQuery(req)
	if err != nil {
		return nil, err
	}

	ins := namespace.NewDefaultNamespace()
	if err := s.col.FindOne(ctx, r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("namespace %s not found", req)
		}

		return nil, exception.NewInternalServerError("find namespace %s error, %s", req.Name, err)
	}

	return ins, nil
}

func (s *impl) DeleteNamespace(ctx context.Context, req *namespace.DeleteNamespaceRequest) (*namespace.Namespace, error) {
	ns, err := s.DescribeNamespace(ctx, namespace.NewDescriptNamespaceRequestByName(req.Domain, req.Name))
	if err != nil {
		return nil, err
	}

	r, err := newDeleteRequest(req)
	if err != nil {
		return nil, err
	}

	_, err = s.col.DeleteOne(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("delete namespace(%s) error, %s", req.Name, err)
	}

	// 清除空间管理的所有策略
	_, err = s.policy.DeletePolicy(ctx, policy.NewDeletePolicyRequestWithNamespace(req.Domain, req.Name))
	if err != nil {
		s.log.Error().Msgf("delete namespace policy error, %s", err)
	}

	return ns, nil
}
