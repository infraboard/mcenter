# traefik适配

## cmdb 和 keyauth services 配置
traefik/http/services/cmdb/loadBalancer/servers/0/url	http://127.0.0.1:8060
traefik/http/services/cmdb/loadBalancer/servers/0/url	h2c://127.0.0.1:18060
traefik/http/services/keyauth/loadBalancer/servers/0/url	http://127.0.0.1:8050
traefik/http/services/keyauth/loadBalancer/servers/0/url	h2c://127.0.0.1:18050

## cmdb 和 keyauth router配置
traefik/http/routers/cmdb-api/entryPoints/0	web
traefik/http/routers/cmdb-api/rule	PathPrefix(`/cmdb/api/v1`)
traefik/http/routers/cmdb-api/service cmdb-api

traefik/http/routers/cmdb-grpc/entryPoints/0 grpc
traefik/http/routers/cmdb-grpc/rule PathPrefix(`infraboard.cmdb`)
traefik/http/routers/cmdb-grpc/service cmdb-grpc

traefik/http/routers/keyauth-api/entryPoints/0 web
traefik/http/routers/keyauth-api/rule	PathPrefix(`/keyauth/api/v1`)
traefik/http/routers/keyauth-api/service keyauth-api

traefik/http/routers/keyauth-grpc/entryPoints/0 grpc
traefik/http/routers/keyauth-grpc/rule PathPrefix(`infraboard.keyauth`)
traefik/http/routers/keyauth-grpc/service keyauth-grpc

```yaml
## Dynamic configuration
http:
  services:
    app:
      weighted:
        healthCheck: {}
        services:
        - name: appv1
          weight: 3
        - name: appv2
          weight: 1

    appv1:
      loadBalancer:
        healthCheck:
          path: /status
          interval: 10s
          timeout: 3s
        servers:
        - url: "http://private-ip-server-1/"

    appv2:
      loadBalancer:
        healthCheck:
          path: /status
          interval: 10s
          timeout: 3s
        servers:
        - url: "http://private-ip-server-2/"
```