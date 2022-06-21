package impl

import (
	"context"

	"github.com/infraboard/mcube/app"
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
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	log logger.Logger
	code.UnimplementedRPCServer

	user    user.Service
	token   token.Service
	setting setting.Service
	notify  notify.Service
}

func (s *service) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	dc := db.Collection("domain")
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
	s.user = app.GetInternalApp(user.AppName).(user.Service)
	s.token = app.GetInternalApp(token.AppName).(token.Service)
	s.setting = app.GetInternalApp(setting.AppName).(setting.Service)
	s.notify = app.GetInternalApp(notify.AppName).(notify.Service)
	return nil
}

func (s *service) Name() string {
	return code.AppName
}

func (s *service) Registry(server *grpc.Server) {
	code.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}
