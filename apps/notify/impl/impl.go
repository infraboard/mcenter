package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/v2/ioc/config/grpc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
)

func init() {
	ioc.Controller().Registry(&service{})
}

type service struct {
	col *mongo.Collection
	notify.UnimplementedRPCServer
	ioc.ObjectImpl

	user   user.Service
	log    *zerolog.Logger
	domain domain.Service
}

func (s *service) Init() error {
	s.col = ioc_mongo.DB().Collection(s.Name())
	s.log = log.Sub(s.Name())
	s.user = ioc.Controller().Get(user.AppName).(user.Service)
	s.domain = ioc.Controller().Get(domain.AppName).(domain.Service)

	notify.RegisterRPCServer(grpc.Get().Server(), s)
	return nil
}

func (s *service) Name() string {
	return notify.AppName
}
