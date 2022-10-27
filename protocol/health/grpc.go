package health

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

func RegisterGrpcHealthServer(s grpc.ServiceRegistrar) {
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(s, healthcheck)
}

func NewHealthCheckRequest() *healthgrpc.HealthCheckRequest {
	return &healthgrpc.HealthCheckRequest{}
}
