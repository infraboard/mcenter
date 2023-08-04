package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcenter/apps/token/security"
	"github.com/infraboard/mcenter/conf"

	_ "github.com/infraboard/mcenter/apps/token/provider/all"
)

func init() {
	ioc.RegistryController(&service{})
}

type service struct {
	col *mongo.Collection
	token.UnimplementedRPCServer
	ioc.IocObjectImpl
	log logger.Logger

	policy  policy.Service
	ns      namespace.Service
	checker security.Checker
	domain  domain.Service
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
	s.ns = ioc.GetController(namespace.AppName).(namespace.Service)
	s.policy = ioc.GetController(policy.AppName).(policy.Service)
	s.domain = ioc.GetController(domain.AppName).(domain.Service)
	s.notify = ioc.GetController(notify.AppName).(notify.Service)

	s.checker, err = security.NewChecker()
	if err != nil {
		return fmt.Errorf("new checker error, %s", err)
	}

	// 初始化所有的auth provider
	if err := provider.Init(); err != nil {
		return err
	}

	return nil
}

func (s *service) Name() string {
	return token.AppName
}

func (s *service) Registry(server *grpc.Server) {
	token.RegisterRPCServer(server, s)
}
