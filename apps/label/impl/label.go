package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/label"
	"github.com/infraboard/mcube/exception"
)

// 创建标签
func (i *impl) CreateLabel(ctx context.Context, in *label.CreateLabelRequest) (
	*label.Label, error) {
	ins, err := label.New(in)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a label document error, %s", err)
	}
	return ins, nil
}

// 查询标签列表
func (i *impl) QueryLabel(ctx context.Context, in *label.QueryLabelRequest) (
	*label.LabelSet, error) {
	r := newQueryRequest(in)
	resp, err := i.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find label error, error is %s", err)
	}

	set := label.NewLabelSet()
	// 循环
	for resp.Next(ctx) {
		ins := label.NewDefaultLabel()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode label error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get label count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

// 修改标签
func (i *impl) UpdateLabel(ctx context.Context, in *label.UpdateLabelRequest) (
	*label.Label, error) {
	return nil, nil
}

// 删除标签
func (i *impl) DeleteLabel(ctx context.Context, in *label.DeleteLabelRequest) (
	*label.Label, error) {
	return nil, nil
}
