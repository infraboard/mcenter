package rpc

import (
	"context"

	"github.com/infraboard/mcenter/apps/service"
	"google.golang.org/grpc/metadata"
)

// 从上下文中获取认证信息
func GetClientId(ctx context.Context) string {
	return getMetaData(ctx, service.ClientHeaderKey)
}

// 从上下文中获取认证信息
func getMetaData(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	values := md.Get(key)
	if len(values) > 0 {
		return values[0]
	}

	return ""
}
