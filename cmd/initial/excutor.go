package initial

import (
	"context"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/setting"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/app"

	// 注册所有服务
	_ "github.com/infraboard/mcenter/apps"
)

// NewInitialerFromCLI 初始化
func NewExecutorFromCLI() (*excutor, error) {
	e := newExcutor()

	// if err := i.checkIsInit(context.Background()); err != nil {
	// 	return nil, err
	// }

	err := survey.AskOne(
		&survey.Input{
			Message: "请输入公司(组织)名称:",
			Default: "基础设施服务中心",
		},
		&e.domainDescribe,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		return nil, err
	}

	err = survey.AskOne(
		&survey.Input{
			Message: "请输入管理员用户名称:",
			Default: "admin",
		},
		&e.username,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		return nil, err
	}

	var repeatPass string
	err = survey.AskOne(
		&survey.Password{
			Message: "请输入管理员密码:",
		},
		&e.password,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		return nil, err
	}
	err = survey.AskOne(
		&survey.Password{
			Message: "再次输入管理员密码:",
		},
		&repeatPass,
		survey.WithValidator(survey.Required),
		survey.WithValidator(func(ans interface{}) error {
			if ans.(string) != e.password {
				return fmt.Errorf("两次输入的密码不一致")
			}
			return nil
		}),
	)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func newExcutor() *excutor {
	return &excutor{
		domain:    app.GetInternalApp(domain.AppName).(domain.Service),
		namespace: app.GetInternalApp(namespace.AppName).(namespace.Service),
		role:      app.GetInternalApp(role.AppName).(role.Service),
		user:      app.GetInternalApp(user.AppName).(user.Service),
		service:   app.GetInternalApp(service.AppName).(service.MetaService),
	}
}

type excutor struct {
	domainDescribe string
	username       string
	password       string

	domain    domain.Service
	namespace namespace.Service
	role      role.Service
	system    setting.Service
	user      user.Service
	service   service.MetaService
}

func (e *excutor) InitDomain(ctx context.Context) error {
	req := domain.NewCreateDomainRequest()
	req.Name = domain.DEFAULT_DOMAIN
	req.Description = e.domainDescribe
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

	for _, req := range apps.items {
		app, err := e.service.CreateService(ctx, req)
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

func (e *excutor) InitSystemSetting(ctx context.Context) error {
	sysConf := setting.NewDefaultSetting()
	st, err := e.system.UpdateSetting(ctx, sysConf)
	if err != nil {
		return err
	}
	fmt.Println(st)
	return nil
}

func (e *excutor) InitAdminUser(ctx context.Context) error {
	req := user.NewCreateUserRequest()
	req.Type = user.TYPE_SUPPER
	req.Username = strings.TrimSpace(e.username)
	req.Password = strings.TrimSpace(e.password)
	req.Domain = domain.DEFAULT_DOMAIN
	u, err := e.user.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println(u)
	return nil
}
