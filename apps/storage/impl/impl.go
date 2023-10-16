package impl

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcenter/apps/storage"
	ioc_mongo "github.com/infraboard/mcube/ioc/config/mongo"
)

func init() {
	ioc.Controller().Registry(&service{})
}

type service struct {
	log logger.Logger
	db  *mongo.Database
	ioc.ObjectImpl
}

func (s *service) Init() error {
	s.db = ioc_mongo.DB()
	s.log = zap.L().Named("Storage")
	return nil
}

func (s *service) Name() string {
	return storage.AppName
}
