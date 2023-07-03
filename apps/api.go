package apps

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github.com/infraboard/mcenter/apps/code/api"
	_ "github.com/infraboard/mcenter/apps/domain/api"
	_ "github.com/infraboard/mcenter/apps/endpoint/api"
	_ "github.com/infraboard/mcenter/apps/instance/api"
	_ "github.com/infraboard/mcenter/apps/label/api"
	_ "github.com/infraboard/mcenter/apps/namespace/api"
	_ "github.com/infraboard/mcenter/apps/resource/api"
	_ "github.com/infraboard/mcenter/apps/role/api"
	_ "github.com/infraboard/mcenter/apps/service/api"
	_ "github.com/infraboard/mcenter/apps/setting/api"
	_ "github.com/infraboard/mcenter/apps/token/api"
	_ "github.com/infraboard/mcenter/apps/user/api"
)
