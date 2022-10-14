package impl

import (
	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/permission"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	permission.UnimplementedRPCServer

	log      logger.Logger
	policy   policy.Service
	role     role.Service
	endpoint endpoint.Service
}

func (s *service) Config() error {
	s.policy = app.GetInternalApp(policy.AppName).(policy.Service)
	s.role = app.GetInternalApp(role.AppName).(role.Service)
	s.endpoint = app.GetInternalApp(endpoint.AppName).(endpoint.Service)
	s.log = zap.L().Named(s.Name())
	return nil
}

func (s *service) Name() string {
	return permission.AppName
}

func (s *service) Registry(server *grpc.Server) {
	permission.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
	app.RegistryInternalApp(svr)
}
