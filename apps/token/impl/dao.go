package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/exception"
)

func (s *service) save(ctx context.Context, tk *token.Token) error {
	if _, err := s.col.InsertOne(ctx, tk); err != nil {
		return exception.NewInternalServerError("inserted token(%s) document error, %s",
			tk.AccessToken, err)
	}

	return nil
}
