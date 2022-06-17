package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/security"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcenter/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	token.UnimplementedRPCServer
	log logger.Logger

	checker security.Checker
	user    user.Service
	domain  domain.Service
}

func (s *service) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	dc := db.Collection(s.Name())
	indexs := []mongo.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "refresh_token", Value: bsonx.Int32(-1)}},
			Options: options.Index().SetUnique(true),
		},
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
	s.domain = app.GetInternalApp(domain.AppName).(domain.Service)

	s.checker, err = security.NewChecker()
	if err != nil {
		return fmt.Errorf("new checker error, %s", err)
	}
	return nil
}

func (s *service) Name() string {
	return token.AppName
}

func (s *service) Registry(server *grpc.Server) {
	token.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}
