package token

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/pb/resource"
	"go.mongodb.org/mongo-driver/bson"
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

func MakeMongoFilter(m bson.M, scope *resource.Scope) {
	if scope == nil {
		return
	}
	if scope.Domain != "" {
		m["domain"] = scope.Domain
	}
	if scope.Namespace != "" {
		m["namespace"] = scope.Namespace
	}
}

func (ip *IPLocation) IsPublic() bool {
	return ip.City != "" && ip.City != "内网IP"
}
