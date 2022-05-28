package resolver

import (
	"context"
	"time"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc/resolver"
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

var (
	exampleServiceName = "test"
)

func (*McenterResolverBuilder) Build(
	target resolver.Target,
	cc resolver.ClientConn,
	opts resolver.BuildOptions) (
	resolver.Resolver, error) {
	r := &mcenterResolver{
		target:             target,
		cc:                 cc,
		queryTimeoutSecond: 3 * time.Second,
		log:                zap.L().Named("Mcenter Resolver"),
	}
	return r, nil
}

func (*McenterResolverBuilder) Scheme() string {
	return "mcenter"
}

// exampleResolver is a
// Resolver(https://godoc.org/google.golang.org/grpc/resolver#Resolver).
type mcenterResolver struct {
	target             resolver.Target
	cc                 resolver.ClientConn
	mcenter            instance.ServiceClient
	queryTimeoutSecond time.Duration
	log                logger.Logger
}

func (m *mcenterResolver) ResolveNow(o resolver.ResolveNowOptions) {
	// 设置查询的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), m.queryTimeoutSecond)
	defer cancel()

	// 从mcenter中查询该target对应的服务实例
	searchReq := instance.NewSearchRequest()
	addrStrs, err := m.mcenter.Search(ctx, searchReq)
	if err != nil {
		m.log.Errorf("search target %s error, %s", m.target, err)
		return
	}

	// 更新给client
	addrs := make([]resolver.Address, len(addrStrs.Items))
	for i, s := range addrStrs.Items {
		addrs[i] = resolver.Address{Addr: s.RegistryInfo.Address}
	}

	m.cc.UpdateState(resolver.State{Addresses: addrs})
}

func (m *mcenterResolver) Close() {

}

func init() {
	// Register the mcenter ResolverBuilder. This is usually done in a package's
	// init() function.
	resolver.Register(&McenterResolverBuilder{})
}
