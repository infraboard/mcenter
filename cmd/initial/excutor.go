package initial

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/role"
	meta "github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/app"
)

func newExcutor() *excutor {
	return &excutor{
		domain:    app.GetInternalApp(domain.AppName).(domain.Service),
		namespace: app.GetInternalApp(namespace.AppName).(namespace.Service),
		role:      app.GetInternalApp(role.AppName).(role.Service),
	}
}

type excutor struct {
	domain    domain.Service
	namespace namespace.Service
	role      role.Service
}

func (e *excutor) InitDomain(ctx context.Context) error {
	req := domain.NewCreateDomainRequest()
	req.Name = domain.DEFAULT_DOMAIN
	ins, err := e.domain.CreateDomain(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println(ins)
	return nil
}

func (e *excutor) InitNamespace(ctx context.Context) error {
	req := namespace.NewCreateNamespaceRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Name = namespace.DEFAULT_NAMESPACE
	req.Owner = user.SYSTEM_INITAL_USERNAME
	ns, err := e.namespace.CreateNamespace(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println(ns)
	return nil
}

func (e *excutor) InitRole(ctx context.Context) error {
	req := role.CreateAdminRoleRequest()
	r, err := e.role.CreateRole(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println(r)

	req = role.CreateVisitorRoleRequest()
	r, err = e.role.CreateRole(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

func (e *excutor) InitService(ctx context.Context) error {
	apps := NewInitApps()
	apps.Add("maudit", "审计中心")
	apps.Add("cmdb", "资源中心")

	impl := app.GetInternalApp(meta.AppName).(meta.MetaService)

	for _, req := range apps.items {
		app, err := impl.CreateService(context.Background(), req)
		if err != nil {
			return err
		}
		fmt.Printf("init app %s success, client_id: %s, client_secret: %s",
			req.Name,
			app.Credential.ClientId,
			app.Credential.ClientSecret,
		)
	}

	return nil
}
