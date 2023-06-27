package token

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/pb/resource"
)

func GetTokenFromRequest(r *restful.Request) *Token {
	tk := r.Attribute(TOKEN_ATTRIBUTE_NAME)
	if tk == nil {
		return nil
	}
	return tk.(*Token)
}

func (t *Token) GenScope() *resource.Scope {
	s := resource.NewScope()
	s.Domain = t.Domain
	s.Namespace = t.Namespace
	return s
}
