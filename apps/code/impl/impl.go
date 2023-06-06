package impl

import (
	"context"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/setting"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcenter/conf"
	"github.com/infraboard/mcube/ioc"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	log logger.Logger
	code.UnimplementedRPCServer
	ioc.IocObjectImpl

	user    user.Service
	token   token.Service
	setting setting.Service
	notify  notify.Service
}

func (s *service) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	dc := db.Collection(s.Name())
	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "issue_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err = dc.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = dc
	s.log = zap.L().Named(s.Name())
	s.user = ioc.GetController(user.AppName).(user.Service)
	s.token = ioc.GetController(token.AppName).(token.Service)
	s.setting = ioc.GetController(setting.AppName).(setting.Service)
	s.notify = ioc.GetController(notify.AppName).(notify.Service)
	return nil
}

func (s *service) Name() string {
	return code.AppName
}

func (s *service) Registry(server *grpc.Server) {
	code.RegisterRPCServer(server, svr)
}

func init() {
	ioc.RegistryController(svr)
}
