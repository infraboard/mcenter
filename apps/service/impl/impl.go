package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/conf"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	service.UnimplementedMetaServiceServer
}

func (s *impl) Config() error {

	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	s.col = db.Collection(s.Name())

	s.log = zap.L().Named(s.Name())
	return nil
}

func (s *impl) Name() string {
	return service.AppName
}

func (s *impl) Registry(server *grpc.Server) {
	service.RegisterMetaServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
