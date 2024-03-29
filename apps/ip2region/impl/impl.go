package impl

import (
	"sync"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/apps/oss"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/ip2region"
)

func init() {
	ioc.Controller().Registry(&service{
		bucketName: "ip2region",
		dbFileName: "ip2region.db",
	})
}

type service struct {
	oss        oss.Service
	log        *zerolog.Logger
	bucketName string
	dbFileName string
	dbReader   *ip2region.IPReader
	sync.Mutex
	ioc.ObjectImpl
}

func (s *service) Init() error {
	s.oss = ioc.Controller().Get(oss.AppName).(oss.Service)

	s.log = log.Sub("IP2Region")
	return nil
}

func (s *service) Name() string {
	return ip2region.AppName
}
