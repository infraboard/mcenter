package all

import (
	// 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序
	_ "github.com/infraboard/mcenter/apps/code/impl"
	_ "github.com/infraboard/mcenter/apps/domain/impl"
	_ "github.com/infraboard/mcenter/apps/endpoint/impl"
	_ "github.com/infraboard/mcenter/apps/instance/impl"
	_ "github.com/infraboard/mcenter/apps/notify/impl"
	_ "github.com/infraboard/mcenter/apps/resource/impl"
	_ "github.com/infraboard/mcenter/apps/service/impl"
	_ "github.com/infraboard/mcenter/apps/token/impl"
	_ "github.com/infraboard/mcenter/apps/user/impl"
)
