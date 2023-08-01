package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreateUser(ctx context.Context, req *user.CreateUserRequest) (
	*user.User, error) {
	u, err := user.New(req)
	if err != nil {
		return nil, err
	}

	// 如果是管理员创建的账号需要用户自己重置密码
	if req.CreateFrom.IsIn(user.CREATE_FROM_PRIMARY_ACCOUNT) {
		u.Password.SetNeedReset("admin created user need reset when first login")
	}

	if err := s.save(ctx, u); err != nil {
		return nil, err
	}

	u.Password = nil
	return u, nil
}

// 查询用户列表
func (s *service) QueryUser(ctx context.Context, req *user.QueryUserRequest) (
	*user.UserSet, error) {
	r := newQueryRequest(req)
	resp, err := s.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find user error, error is %s", err)
	}

	set := user.NewUserSet()
	// 循环
	if !req.SkipItems {
		for resp.Next(ctx) {
			ins := user.NewDefaultUser()
			if err := resp.Decode(ins); err != nil {
				return nil, exception.NewInternalServerError("decode user error, error is %s", err)
			}
			ins.Desensitize()
			set.Add(ins)
		}
	}

	// count
	count, err := s.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get user count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

// 查询用户详情
func (s *service) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (
	*user.User, error) {
	filter := bson.M{}
	switch req.DescribeBy {
	case user.DESCRIBE_BY_USER_ID:
		filter["_id"] = req.Id
	case user.DESCRIBE_BY_FEISHU_USER_ID:
		filter["feishu.user_id"] = req.Id
	case user.DESCRIBE_BY_USER_NAME:
		filter["username"] = req.Username
	default:
		return nil, exception.NewBadRequest("unknow desribe by %s", req.DescribeBy)
	}

	ins := user.NewDefaultUser()
	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("user %s not found", req)
		}

		return nil, exception.NewInternalServerError("user %s error, %s", req, err)
	}

	ins.SetupDefault()
	return ins, nil
}

// 修改用户信息
func (s *service) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (
	*user.User, error) {
	ins, err := s.DescribeUser(ctx, user.NewDescriptUserRequestById(req.UserId))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case request.UpdateMode_PUT:
		ins.Update(req)
	case request.UpdateMode_PATCH:
		ins.Spec.Description = req.Description
		err := ins.Patch(req)
		if err != nil {
			return nil, exception.NewBadRequest("patch error, %s", err)
		}
	}

	if err := s.update(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

// 删除用户
func (s *service) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (
	*user.UserSet, error) {
	// 判断这些要删除的用户是否存在
	queryReq := user.NewQueryUserRequest()
	queryReq.UserIds = req.UserIds
	set, err := s.QueryUser(ctx, queryReq)
	if err != nil {
		return nil, err
	}

	noExist := []string{}
	for _, uid := range req.UserIds {
		if !set.HasUser(uid) {
			noExist = append(noExist, uid)
		}
	}
	if len(noExist) > 0 {
		return nil, exception.NewBadRequest("user %v not found", req.UserIds)
	}

	if err := s.delete(ctx, set); err != nil {
		return nil, err
	}
	return set, nil
}

// 修改用户密码, 用户需要知道原先密码
func (s *service) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest) (
	*user.Password, error) {
	ins, err := s.DescribeUser(ctx, user.NewDescriptUserRequestById(req.UserId))
	if err != nil {
		return nil, err
	}

	err = ins.Password.CheckPassword(req.OldPass)
	if err != nil {
		return nil, err
	}

	pass, err := user.NewHashedPassword(req.NewPass)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	pass.NeedReset = req.IsReset
	pass.ResetReason = req.ResetReason
	ins.Password = pass

	err = s.update(ctx, ins)
	if err != nil {
		return nil, err
	}
	return pass, nil
}

// 重置密码, 无需知道原先密码, 主账号执行
func (s *service) ResetPassword(ctx context.Context, req *user.ResetPasswordRequest) (
	*user.Password, error) {
	ins, err := s.DescribeUser(ctx, user.NewDescriptUserRequestById(req.UserId))
	if err != nil {
		return nil, err
	}

	pass, err := user.NewHashedPassword(req.NewPass)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	pass.NeedReset = req.IsReset
	pass.ResetReason = req.ResetReason
	ins.Password = pass

	err = s.update(ctx, ins)
	if err != nil {
		return nil, err
	}
	return pass, nil
}
