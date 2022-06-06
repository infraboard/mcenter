package rpc

import (
	"github.com/caarlos0/env/v6"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/infraboard/mcenter/apps/application"
	"github.com/infraboard/mcenter/apps/instance"
)

var (
	client *ClientSet
)

func C() *ClientSet {
	if client == nil {
		panic("mcenter client config not load")
	}
	return client
}

func LoadClientFromEnv() error {
	conf := NewDefaultConfig()
	err := env.Parse(conf)
	if err != nil {
		return err
	}

	client, err = NewClient(conf)
	if err != nil {
		return err
	}

	return nil
}

func LoadClientFromConfig(conf *Config) error {
	c, err := NewClient(conf)
	if err != nil {
		return err
	}
	client = c
	return nil
}

// NewClient todo
func NewClient(conf *Config) (*ClientSet, error) {
	zap.DevelopmentSetup()
	log := zap.L()

	conn, err := grpc.Dial(
		conf.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(conf.Credentials()),
	)
	if err != nil {
		return nil, err
	}

	return &ClientSet{
		conn: conn,
		log:  log,
	}, nil
}

// Client 客户端
type ClientSet struct {
	conn *grpc.ClientConn
	log  logger.Logger
}

// Instance服务的SDK
func (c *ClientSet) Instance() instance.ServiceClient {
	return instance.NewServiceClient(c.conn)
}

// Application服务的SDK
func (c *ClientSet) Application() application.ServiceClient {
	return application.NewServiceClient(c.conn)
}
