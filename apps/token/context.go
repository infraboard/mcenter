package token

import (
	"context"

	"github.com/emicklei/go-restful/v3"
)

var (
	CONTEXT_KEY = struct{}{}
)

func WithTokenCtx(r *restful.Request) context.Context {
	return context.WithValue(
		r.Request.Context(),
		CONTEXT_KEY,
		GetTokenFromRequest(r),
	)
}

func GetTokenFromCtx(ctx context.Context) *Token {
	tk := ctx.Value(CONTEXT_KEY)
	if tk == nil {
		return nil
	}
	return tk.(*Token)
}
