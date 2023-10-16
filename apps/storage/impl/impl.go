package impl

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcenter/apps/storage"
	"github.com/infraboard/mcube/ioc/config/logger"
	ioc_mongo "github.com/infraboard/mcube/ioc/config/mongo"
)

func init() {
	ioc.Controller().Registry(&service{})
}

type service struct {
	log *zerolog.Logger
	db  *mongo.Database
	ioc.ObjectImpl
}

func (s *service) Init() error {
	s.db = ioc_mongo.DB()
	s.log = logger.Sub("Storage")
	return nil
}

func (s *service) Name() string {
	return storage.AppName
}
