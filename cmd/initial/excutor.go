package initial

import (
	"context"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/label"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/v2/ioc"

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
		domain:    ioc.Controller().Get(domain.AppName).(domain.Service),
		namespace: ioc.Controller().Get(namespace.AppName).(namespace.Service),
		role:      ioc.Controller().Get(role.AppName).(role.Service),
		user:      ioc.Controller().Get(user.AppName).(user.Service),
		service:   ioc.Controller().Get(service.AppName).(service.MetaService),
		label:     ioc.Controller().Get(service.AppName).(label.Service),
	}
}

type excutor struct {
	domainDescribe string
	username       string
	password       string

	domain    domain.Service
	namespace namespace.Service
	role      role.Service
	user      user.Service
	service   service.MetaService
	label     label.Service
}

func (e *excutor) InitDomain(ctx context.Context) error {
	req := domain.NewCreateDomainRequest()
	req.Name = domain.DEFAULT_DOMAIN
	req.Description = e.domainDescribe
	ins, err := e.domain.CreateDomain(ctx, req)
	if err != nil {
		return err
	}

	fmt.Printf("初始化域: %17s [成功]", ins.Spec.Name)
	fmt.Println()
	return nil
}

func (e *excutor) InitNamespace(ctx context.Context) error {
	req := namespace.NewCreateNamespaceRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Name = namespace.DEFAULT_NAMESPACE
	req.Visible = namespace.VISIBLE_PUBLIC
	req.Description = "默认空间"
	req.Owner = e.username
	ns, err := e.namespace.CreateNamespace(ctx, req)
	if err != nil {
		return fmt.Errorf("初始化空间失败: %s", err)
	}

	fmt.Printf("初始化空间: %15s [成功]", ns.Spec.Name)
	fmt.Println()

	req.Name = namespace.SYSTEM_NAMESPACE
	req.Visible = namespace.VISIBLE_PRIVATE
	req.Description = "系统空间"
	ns, err = e.namespace.CreateNamespace(ctx, req)
	if err != nil {
		return fmt.Errorf("初始化空间失败: %s", err)
	}
	fmt.Printf("初始化空间: %15s [成功]", ns.Spec.Name)
	fmt.Println()
	return nil
}

func (e *excutor) InitRole(ctx context.Context) error {
	req := role.CreateAdminRoleRequest(e.username)
	r, err := e.role.CreateRole(ctx, req)
	if err != nil {
		return fmt.Errorf("初始化角色失败: %s", err)
	}
	fmt.Printf("初始化角色: %15s [成功]", r.Spec.Name)
	fmt.Println()

	req = role.CreateVisitorRoleRequest(e.username)
	r, err = e.role.CreateRole(ctx, req)
	if err != nil {
		return fmt.Errorf("初始化角色失败: %s", err)
	}
	fmt.Printf("初始化角色: %15s [成功]", r.Spec.Name)
	fmt.Println()
	return nil
}

func (e *excutor) InitService(ctx context.Context) error {
	apps := NewInitApps()
	apps.Add("maudit", "审计中心")
	apps.Add("mflow", "流水线服务")
	apps.Add("moperator", "平台k8s Operator状态同步器")
	apps.Add("mpaas", "发布中心")
	apps.Add("cmdb", "资源中心")

	for _, req := range apps.items {
		svc, err := e.service.CreateService(ctx, req)
		if err != nil {
			return err
		}
		fmt.Printf("初始化服务: %15s [成功]", svc.Spec.Name)
		fmt.Println()
	}

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
	fmt.Printf("初始化系统管理员: %9s [成功]", u.Spec.Username)
	fmt.Println()
	return nil
}

func (e *excutor) InitLabel(ctx context.Context) error {
	lables := label.BuildInLables()
	for i := range lables {
		ins, err := e.label.CreateLabel(ctx, lables[i])
		if err != nil {
			return fmt.Errorf("初始化标签失败: %s", err)
		}

		fmt.Printf("初始化标签: %15s [成功]", ins.Spec.Key)
	}
	fmt.Println()
	return nil
}
