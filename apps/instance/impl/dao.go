package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcube/exception"
)

func (i *impl) save(ctx context.Context, ins *instance.Instance) error {
	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted book(%s) document error, %s",
			ins.FullName(), err)
	}
	return nil
}
