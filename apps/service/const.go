package service

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc/metadata"
)

const (
	GRPC_CLIENT_SERVICE_JSON = "service_json"
)

func GetServiceFromCtx(ctx context.Context) (*Service, error) {
	ins := NewDefaultService()

	// 重上下文中获取认证信息, 如果没有返回默认服务信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ins, nil
	}

	serviceJSON := GetValueFromMetaData(md, GRPC_CLIENT_SERVICE_JSON)
	if serviceJSON != "" {
		err := json.Unmarshal([]byte(serviceJSON), ins)
		if err != nil {
			return nil, err
		}
	}

	return ins, nil
}

func GetValueFromMetaData(md metadata.MD, key string) string {
	v := md.Get(key)
	if len(v) == 0 {
		return ""
	}

	return v[0]
}

// 从上下文中获取认证信息
func GetMetaData(ctx context.Context, key string) string {
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
