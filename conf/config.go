package conf

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/cache/memory"
	"github.com/infraboard/mcube/cache/redis"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/validator"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

func newConfig() *Config {
	return &Config{
		App:    newDefaultAPP(),
		Log:    newDefaultLog(),
		Cache:  newDefaultCache(),
		Mongo:  newDefaultMongoDB(),
		Jaeger: newJaeger(),
	}
}

// Config 应用配置
type Config struct {
	App    *app     `toml:"app"`
	Log    *log     `toml:"log"`
	Mongo  *mongodb `toml:"mongodb"`
	Cache  *_cache  `toml:"cache"`
	Jaeger *jaeger  `toml:"jaeger"`
}

// InitGlobal 注入全局变量
func (c *Config) InitGlobal() error {
	// 加载全局缓存
	if err := c.Cache.LoadCache(); err != nil {
		return fmt.Errorf("load cache error, %s", err)
	}

	if err := validator.Init(); err != nil {
		return nil
	}

	return nil
}

type app struct {
	Name       string `toml:"name" env:"APP_NAME"`
	EncryptKey string `toml:"encrypt_key" env:"APP_ENCRYPT_KEY"`
	HTTP       *http  `toml:"http"`
	GRPC       *grpc  `toml:"grpc"`
}

func newDefaultAPP() *app {
	return &app{
		Name:       "cmdb",
		EncryptKey: "defualt app encrypt key",
		HTTP:       newDefaultHTTP(),
		GRPC:       newDefaultGRPC(),
	}
}

type http struct {
	Host      string `toml:"host" env:"HTTP_HOST"`
	Port      string `toml:"port" env:"HTTP_PORT"`
	EnableSSL bool   `toml:"enable_ssl" env:"HTTP_ENABLE_SSL"`
	CertFile  string `toml:"cert_file" env:"HTTP_CERT_FILE"`
	KeyFile   string `toml:"key_file" env:"HTTP_KEY_FILE"`
}

func (a *http) Addr() string {
	return a.Host + ":" + a.Port
}

func newDefaultHTTP() *http {
	return &http{
		Host: "127.0.0.1",
		Port: "8010",
	}
}

type grpc struct {
	Host      string `toml:"host" env:"GRPC_HOST"`
	Port      string `toml:"port" env:"GRPC_PORT"`
	EnableSSL bool   `toml:"enable_ssl" env:"GRPC_ENABLE_SSL"`
	CertFile  string `toml:"cert_file" env:"GRPC_CERT_FILE"`
	KeyFile   string `toml:"key_file" env:"GRPC_KEY_FILE"`
}

func (a *grpc) Addr() string {
	return a.Host + ":" + a.Port
}

func newDefaultGRPC() *grpc {
	return &grpc{
		Host: "127.0.0.1",
		Port: "18010",
	}
}

type log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

func newDefaultLog() *log {
	return &log{
		Level:   "debug",
		PathDir: "logs",
		Format:  "text",
		To:      "stdout",
	}
}

func newDefaultMongoDB() *mongodb {
	m := &mongodb{
		UserName:       "mcenter",
		Password:       "123456",
		Database:       "mcenter",
		AuthDB:         "",
		Endpoints:      []string{"127.0.0.1:27017"},
		K8sServiceName: "MONGODB",
	}
	m.LoadK8sEnv()
	return m
}

type mongodb struct {
	Endpoints      []string `toml:"endpoints" env:"MONGO_ENDPOINTS" envSeparator:","`
	UserName       string   `toml:"username" env:"MONGO_USERNAME"`
	Password       string   `toml:"password" env:"MONGO_PASSWORD"`
	Database       string   `toml:"database" env:"MONGO_DATABASE"`
	AuthDB         string   `toml:"auth_db" env:"MONGO_AUTH_DB"`
	K8sServiceName string   `toml:"k8s_service_name" env:"K8S_SERVICE_NAME"`

	client *mongo.Client
	lock   sync.Mutex
}

// 当 Pod 运行在 Node 上，kubelet 会为每个活跃的 Service 添加一组环境变量。
// kubelet 为 Pod 添加环境变量 {SVCNAME}_SERVICE_HOST 和 {SVCNAME}_SERVICE_PORT。
// 这里 Service 的名称需大写，横线被转换成下划线
// 具体请参考: https://kubernetes.io/zh-cn/docs/concepts/services-networking/service/#environment-variables
func (m *mongodb) LoadK8sEnv() {
	host := os.Getenv(fmt.Sprintf("%s_SERVICE_HOST", m.K8sServiceName))
	port := os.Getenv(fmt.Sprintf("%s_SERVICE_PORT", m.K8sServiceName))
	addr := fmt.Sprintf("%s:%s", host, port)
	if host != "" && port != "" {
		m.Endpoints = []string{addr}
	}
}

func (m *mongodb) GetAuthDB() string {
	if m.AuthDB != "" {
		return m.AuthDB
	}

	return m.Database
}

func (m *mongodb) GetDB() (*mongo.Database, error) {
	conn, err := m.Client()
	if err != nil {
		return nil, err
	}
	return conn.Database(m.Database), nil
}

// 关闭数据库连接
func (m *mongodb) Close(ctx context.Context) error {
	if m.client == nil {
		return nil
	}

	return m.client.Disconnect(ctx)
}

// Client 获取一个全局的mongodb客户端连接
func (m *mongodb) Client() (*mongo.Client, error) {
	// 加载全局数据量单例
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.client == nil {
		conn, err := m.getClient()
		if err != nil {
			return nil, err
		}
		m.client = conn
	}

	return m.client, nil
}

func (m *mongodb) getClient() (*mongo.Client, error) {
	opts := options.Client()

	if m.UserName != "" && m.Password != "" {
		cred := options.Credential{
			AuthSource: m.GetAuthDB(),
		}

		cred.Username = m.UserName
		cred.Password = m.Password
		cred.PasswordSet = true
		opts.SetAuth(cred)
	}
	opts.SetHosts(m.Endpoints)
	opts.SetConnectTimeout(5 * time.Second)
	opts.Monitor = otelmongo.NewMonitor(
		otelmongo.WithCommandAttributeDisabled(true),
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("new mongodb client error, %s", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping mongodb server(%s) error, %s", m.Endpoints, err)
	}

	return client, nil
}

func newDefaultCache() *_cache {
	return &_cache{
		Type:   "memory",
		Memory: memory.NewDefaultConfig(),
		Redis:  redis.NewDefaultConfig(),
	}
}

type _cache struct {
	Type   string         `toml:"type" json:"type" yaml:"type" env:"MCENTER_CACHE_TYPE"`
	Memory *memory.Config `toml:"memory" json:"memory" yaml:"memory"`
	Redis  *redis.Config  `toml:"redis" json:"redis" yaml:"redis"`
}

func (c *_cache) LoadCache() error {
	// 设置全局缓存
	switch c.Type {
	case "memory", "":
		ins := memory.NewCache(c.Memory)
		cache.SetGlobal(ins)
		zap.L().Info("use cache in local memory")
	case "redis":
		ins := redis.NewCache(c.Redis)
		cache.SetGlobal(ins)
		zap.L().Info("use redis to cache")
	default:
		return fmt.Errorf("unknown cache type: %s", c.Type)
	}

	return nil
}

func newJaeger() *jaeger {
	return &jaeger{}
}

type jaeger struct {
	Endpoint string `toml:"endpoint" json:"endpoint" yaml:"endpoint" env:"JAEGER_ENDPOINT"`
}
