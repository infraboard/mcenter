package middleware

import (
	_ "github.com/infraboard/mcenter/clients/rpc/middleware/auth/gorestful"
	_ "github.com/infraboard/mcenter/clients/rpc/middleware/auth/grpc"
	_ "github.com/infraboard/mcenter/clients/rpc/middleware/registry/endpoint"
	_ "github.com/infraboard/mcenter/clients/rpc/middleware/registry/service"
)
