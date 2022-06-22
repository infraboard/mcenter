package all

import (
	// 注册所有内部服务模块, 无须对外暴露的服务, 用于内部依赖
	_ "github.com/infraboard/mcenter/apps/ip2region/impl"
	_ "github.com/infraboard/mcenter/apps/setting/impl"
	_ "github.com/infraboard/mcenter/apps/storage/impl"
)
