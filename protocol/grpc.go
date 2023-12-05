package protocol

import (
	"net"

	"google.golang.org/grpc"

	"github.com/infraboard/mcube/v2/grpc/middleware/recovery"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/protocol/auth"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/infraboard/mcube/v2/ioc/config/logger"
)

// NewGRPCService todo
func NewGRPCService() *GRPCService {
	appImpl := ioc.GetController(service.AppName).(service.MetaService)
	rc := recovery.NewInterceptor(recovery.NewZeroLogRecoveryHandler())
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		rc.UnaryServerInterceptor(),
		otelgrpc.UnaryServerInterceptor(),
		auth.GrpcAuthUnaryServerInterceptor(appImpl),
	))

	return &GRPCService{
		svr: grpcServer,
		l:   logger.Sub("grpc"),
		c:   application.App().GRPC,
	}
}

// GRPCService grpc服务
type GRPCService struct {
	svr *grpc.Server
	l   *zerolog.Logger
	c   *application.Grpc
}

// Start 启动GRPC服务
func (s *GRPCService) Start() {
	// 装载所有GRPC服务
	ioc.LoadGrpcController(s.svr)

	// 启动HTTP服务
	lis, err := net.Listen("tcp", s.c.Addr())
	if err != nil {
		s.l.Error().Msgf("listen grpc tcp conn error, %s", err)
		return
	}

	s.l.Info().Msgf("GRPC 服务监听地址: %s", s.c.Addr())
	if err := s.svr.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Info().Msgf("service is stopped")
			return
		}
		s.l.Error().Msgf("start grpc service error, %s", err.Error())
	}
}

// Stop 启动GRPC服务
func (s *GRPCService) Stop() error {
	s.svr.GracefulStop()
	return nil
}
