package all

import (
	_ "github.com/infraboard/mcenter/apps/token/provider/feishu"
	_ "github.com/infraboard/mcenter/apps/token/provider/ldap"
	_ "github.com/infraboard/mcenter/apps/token/provider/password"
	_ "github.com/infraboard/mcenter/apps/token/provider/private_token"
	_ "github.com/infraboard/mcenter/apps/token/provider/refresh"
)
