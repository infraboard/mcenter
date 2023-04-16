package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/permission"
)

func TestCheckPermission(t *testing.T) {
	req := permission.NewCheckPermissionRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.Username = "test"

	// 检查test用户在默认空间下是否有访问mpaas服务的构建配置功能
	req.ServiceId = "cd08fc9c"
	req.Path = "POST./mpaas/api/v1/builds"
	r, err := impl.CheckPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())

	// 检查是否有创建Pipeline权限
	req.Path = "POST./mpaas/api/v1/pipelines"
	r, err = impl.CheckPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}
