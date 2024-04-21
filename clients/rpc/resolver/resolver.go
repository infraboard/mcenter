package resolver

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mcube/v2/grpc/balancer/wrr"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"
)

func init() {
	// Register the mcenter ResolverBuilder. This is usually done in a package's
	// init() function.
	resolver.Register(&McenterResolverBuilder{})
}

const (
	Scheme = "mcenter"
)

// Following is an example name resolver. It includes a
// ResolverBuilder(https://godoc.org/google.golang.org/grpc/resolver#Builder)
// and a Resolver(https://godoc.org/google.golang.org/grpc/resolver#Resolver).
//
// A ResolverBuilder is registered for a scheme (in this example, "example" is
// the scheme). When a ClientConn is created for this scheme, the
// ResolverBuilder will be picked to build a Resolver. Note that a new Resolver
// is built for each ClientConn. The Resolver will watch the updates for the
// target, and send updates to the ClientConn.

// McenterResolverBuilder is a ResolverBuilder.
type McenterResolverBuilder struct{}

func (*McenterResolverBuilder) Build(
	target resolver.Target,
	cc resolver.ClientConn,
	opts resolver.BuildOptions) (
	resolver.Resolver, error) {

	r := &mcenterResolver{
		mcenter:            rpc.C().Instance(),
		target:             target,
		cc:                 cc,
		queryTimeoutSecond: 3 * time.Second,
		log:                log.Sub("mcenter.resolver"),
		refreshInterval:    30,
	}

	// 同步更新一次
	r.resolve()
	// 后台刷新
	go r.Watch()
	return r, nil
}

func (*McenterResolverBuilder) Scheme() string {
	return Scheme
}

// exampleResolver is a
// Resolver(https://godoc.org/google.golang.org/grpc/resolver#Resolver).
type mcenterResolver struct {
	mcenter            instance.RPCClient
	target             resolver.Target
	cc                 resolver.ClientConn
	queryTimeoutSecond time.Duration
	log                *zerolog.Logger
	addrHash           string
	refreshInterval    int32
}

func (m *mcenterResolver) refreshTime() time.Duration {
	return time.Duration(m.refreshInterval) * time.Second
}

func (m *mcenterResolver) ResolveNow(o resolver.ResolveNowOptions) {}

func (m *mcenterResolver) resolve() {
	// 从mcenter中查询该target对应的服务实例
	address, err := m.search()
	if err != nil {
		m.log.Debug().Msgf("update endpoints error, %s", err)
		return
	}

	// 更新给client
	err = m.cc.UpdateState(resolver.State{Addresses: address})
	if err != nil {
		m.log.Error().Msgf("update state error, %s", err)
		return
	}
}

// 查询名称对应的实例
func (m *mcenterResolver) search() ([]resolver.Address, error) {
	req := m.buildSerchReq()
	endpoints := []resolver.Address{}
	if req.ServiceName == "" {
		return nil, fmt.Errorf("application name required")
	}

	// 设置查询的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), m.queryTimeoutSecond)
	defer cancel()

	set, err := m.mcenter.Search(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("search target %s error, %s", m.target.URL.String(), err)
	}
	if set.Len() == 0 {
		return nil, fmt.Errorf("no service instance")
	}

	// 是否需要更新
	newHash := set.RegistryInfoHash()
	if m.addrHash == newHash {
		return nil, fmt.Errorf("address not changed")
	}
	m.addrHash = newHash

	// 优先获取对应匹配的组, 如果没匹配的，则使用最老的那个组
	items := set.GetGroupInstance(req.Group)
	if len(items) == 0 {
		items = set.GetOldestGroup()
	}

	addrString := []string{}
	for i := range items {
		item := items[i]
		attr := attributes.New("region", item.RegistryInfo.Region)
		attr.WithValue("environment", item.RegistryInfo.Environment)
		attr.WithValue("group", item.RegistryInfo.Group)
		attr.WithValue(wrr.WEIGHT_ATTRIBUTE_KEY, item.Config.Weight)

		endpoints = append(endpoints, resolver.Address{
			Addr:       item.RegistryInfo.Address,
			Attributes: attr,
		})
		addrString = append(addrString, item.RegistryInfo.Address)
	}

	m.log.Info().Msgf("search service params: %s,  address: %v", req.ToJSON(), addrString)
	return endpoints, nil
}

// 构建mcenter实例查询参数
func (m *mcenterResolver) buildSerchReq() *instance.SearchRequest {
	searchReq := instance.NewSearchRequest()
	searchReq.ServiceName = m.target.URL.Host

	qs := m.target.URL.Query()
	searchReq.Page.PageSize = 500
	searchReq.Region = qs.Get("region")
	searchReq.Environment = qs.Get("environment")
	searchReq.Cluster = qs.Get("cluster")
	searchReq.Labels = instance.ParseStrLable(qs.Get("labels"))

	return searchReq
}

func (m *mcenterResolver) Close() {

}
