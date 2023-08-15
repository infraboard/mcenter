package initial

import (
	"fmt"

	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/service"
)

func NewInitApps() *InitApps {
	return &InitApps{
		items: []*service.CreateServiceRequest{},
	}
}

type InitApps struct {
	items []*service.CreateServiceRequest
}

func (i *InitApps) Add(name, descrption string) {
	req := service.NewCreateServiceRequest()
	req.Name = name
	req.Namespace = namespace.SYSTEM_NAMESPACE
	req.Description = descrption
	req.Owner = "admin"
	req.Type = service.Type_CONTAINER_IMAGE
	req.ImageRepository.Address = fmt.Sprintf("registry.cn-hangzhou.aliyuncs.com/infraboard/%s", name)
	i.items = append(i.items, req)
}
