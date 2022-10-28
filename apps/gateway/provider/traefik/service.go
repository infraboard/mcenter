package traefik

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/instance"
)

func (o *EtcdOperator) ServiceKeyValue(ins *instance.Instance) (key, value string) {
	fmt.Sprintf("%s", o.rootKey)
	return "", ""
}

func (o *EtcdOperator) ServiceGroupKey(group string) string {
	return fmt.Sprintf("%s/http/services/%s/loadBalancer/servers", o.rootKey, group)
}

// 参考文档: https://doc.traefik.io/traefik/routing/services/
// 例如: traefik/http/services/keyauth/loadBalancer/servers/0/url	h2c://127.0.0.1:18050
func (o *EtcdOperator) AddService(ctx context.Context, ins *instance.Instance) error {
	// o.client.Put(ctx, "", ins)
	// ins.RegistryInfo.Group
	// resp, err := o.client.Get(ctx, o.prefix+key, clientv3.WithKeysOnly())
	// if err != nil {
	// 	return nil, err
	// }
	// for i := range resp.Kvs {
	// 	keys = append(keys, string(resp.Kvs[i].Key))
	// }
	return nil
}

func (o *EtcdOperator) GetServiceGroup(ctx context.Context, group string) (*instance.InstanceSet, error) {
	resp, err := o.client.Get(ctx, o.ServiceGroupKey(group))
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return nil, nil
}
