package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/exception"
)

func (s *service) save(ctx context.Context, u *user.User) error {
	if _, err := s.col.InsertOne(context.TODO(), u); err != nil {
		return exception.NewInternalServerError("inserted user(%s) document error, %s",
			u.Id, err)
	}

	return nil
}
