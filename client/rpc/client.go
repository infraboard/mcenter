package rpc

import (
	"github.com/caarlos0/env/v6"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/apps/permission"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/token"
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
		conf: conf,
		conn: conn,
		log:  log,
	}, nil
}

// Client 客户端
type ClientSet struct {
	conf *Config
	conn *grpc.ClientConn
	log  logger.Logger
}

func (c *ClientSet) GetClientID() string {
	return c.conf.ClientID
}

// Instance服务的SDK
func (c *ClientSet) Instance() instance.RPCClient {
	return instance.NewRPCClient(c.conn)
}

// Service服务的SDK
func (c *ClientSet) Service() service.RPCClient {
	return service.NewRPCClient(c.conn)
}

// Token服务的SDK
func (c *ClientSet) Token() token.RPCClient {
	return token.NewRPCClient(c.conn)
}

// Permission服务的SDK
func (c *ClientSet) Permission() permission.RPCClient {
	return permission.NewRPCClient(c.conn)
}

// Service服务的SDK
func (c *ClientSet) Serivce() service.RPCClient {
	return service.NewRPCClient(c.conn)
}

// Service服务的SDK
func (c *ClientSet) Endpoint() endpoint.RPCClient {
	return endpoint.NewRPCClient(c.conn)
}
