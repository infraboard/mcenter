package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/user"
	ioc_mongo "github.com/infraboard/mcube/ioc/config/mongo"
)

func init() {
	ioc.Controller().Registry(&service{})
}

type service struct {
	col *mongo.Collection
	notify.UnimplementedRPCServer
	ioc.ObjectImpl

	user   user.Service
	log    logger.Logger
	domain domain.Service
}

func (s *service) Init() error {
	s.col = ioc_mongo.DB().Collection(s.Name())
	s.log = zap.L().Named(s.Name())
	s.user = ioc.GetController(user.AppName).(user.Service)
	s.domain = ioc.GetController(domain.AppName).(domain.Service)
	return nil
}

func (s *service) Name() string {
	return notify.AppName
}

func (s *service) Registry(server *grpc.Server) {
	notify.RegisterRPCServer(server, s)
}
