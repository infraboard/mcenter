package traefik

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/instance"
)

// 参考文档: https://doc.traefik.io/traefik/routing/services/
func (o *EtcdOperator) AddServiceInstance(ctx context.Context, ins *instance.Instance) error {
	err := o.setServiceURL(ctx, ins.ServiceName, 0, ins.RegistryInfo.Address)
	if err != nil {
		return err
	}
	return nil
}

func (o *EtcdOperator) GetServiceInstance(ctx context.Context, serviceName string) (*instance.InstanceSet, error) {
	resp, err := o.client.Get(ctx, o.ServiceInstancePrefixKey(serviceName))
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return nil, nil
}

// 例如: traefik/http/services/keyauth/loadBalancer/servers/0/url	h2c://127.0.0.1:18050
func (o *EtcdOperator) setServiceURL(ctx context.Context, serviceName string, index int, Url string) error {
	key := fmt.Sprintf("%s/loadBalancer/servers/%d/url", o.ServiceInstancePrefixKey(serviceName), index)
	_, err := o.client.Put(ctx, key, Url)
	if err != nil {
		return err
	}
	return nil
}

func (o *EtcdOperator) ServiceInstancePrefixKey(serviceName string) string {
	return fmt.Sprintf("%s/http/services/%s", o.rootKey, serviceName)
}
