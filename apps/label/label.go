package label

import (
	"encoding/json"
	"fmt"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/pb/resource"
	"github.com/infraboard/mcube/v2/tools/hash"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

func NewLabelSet() *LabelSet {
	return &LabelSet{
		Items: []*Label{},
	}
}

func (s *LabelSet) Add(items ...*Label) {
	s.Items = append(s.Items, items...)
}

func (s *LabelSet) Len() int {
	return len(s.Items)
}

func (s *LabelSet) ToJSON() string {
	return pretty.ToJSON(s)
}

func New(in *CreateLabelRequest) (*Label, error) {
	// 填充Value
	in.BuildValue()

	if err := in.Validate(); err != nil {
		return nil, err
	}

	ins := NewDefaultLabel()
	ins.Spec = in
	ins.Meta.Id = hash.FnvHash(in.UniqueKey())
	return ins, nil
}

func (r *CreateLabelRequest) Validate() error {
	return validator.Validate(r)
}

func (r *CreateLabelRequest) BuildValue() {
	for i := range r.EnumOptions {
		item := r.EnumOptions[i]
		item.Value = item.Input
		MakeEnumOption(item)
	}
}

func MakeEnumOption(o *EnumOption) {
	if !o.HasChildren() {
		return
	}

	for i := range o.Children {
		c := o.Children[i]
		c.Value = fmt.Sprintf("%s-%s", o.Input, c.Input)
		MakeEnumOption(c)
	}
}

func (r *CreateLabelRequest) UniqueKey() string {
	switch r.Visiable {
	case resource.VISIABLE_GLOBAL:
		return r.Key
	case resource.VISIABLE_DOMAIN:
		return fmt.Sprintf("%s.%s", r.Domain, r.Key)
	default:
		return fmt.Sprintf("%s.%s.%s", r.Domain, r.Namespace, r.Key)
	}
}

func NewDefaultLabel() *Label {
	return &Label{
		Meta: resource.NewMeta(),
		Spec: NewCreateLabelRequest(),
	}
}

func (s *Label) ToJSON() string {
	return pretty.ToJSON(s)
}

func (b *Label) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*resource.Meta
		*CreateLabelRequest
	}{b.Meta, b.Spec})
}

// 需要被继承的标签
var InheritanceLables = []string{ENV_KEY, RESOURCE_GROUP_KEY, USER_GROUP, DEPLOY_GROUP}

const (
	ENV_KEY            = "Env"
	RESOURCE_GROUP_KEY = "ResourceGroup"
	USER_GROUP         = "UserGroup"
	DEPLOY_GROUP       = "DeployGroup"
)

func MergeInheritanceLabel(src, target map[string]string) {
	for _, k := range InheritanceLables {
		// 如果存在 则不覆盖
		if _, ok := target[k]; ok {
			continue
		}
		// 如果目标元素有值
		if v, ok := src[k]; ok {
			target[k] = v
		}
	}
}

func BuildInLables() (items []*CreateLabelRequest) {
	env := NewCreateLabelRequest()
	env.Visiable = resource.VISIABLE_GLOBAL
	env.Key = ENV_KEY
	env.KeyDesc = "环境"
	env.ValueDesc = "资源所属环境"
	env.Domain = domain.DEFAULT_DOMAIN
	env.Namespace = namespace.DEFAULT_NAMESPACE
	env.Required = true
	env.ValueType = VALUE_TYPE_ENUM
	env.Visiable = resource.VISIABLE_GLOBAL
	env.DefaultValue = "开发"
	env.AddEnumOption(
		NewEnumOption("开发"),
		NewEnumOption("测试"),
		NewEnumOption("生产"),
	)
	items = append(items, env)

	rgroup := NewCreateLabelRequest()
	rgroup.Visiable = resource.VISIABLE_GLOBAL
	rgroup.Key = RESOURCE_GROUP_KEY
	rgroup.KeyDesc = "资源组"
	rgroup.ValueDesc = "资源所属组"
	rgroup.Domain = domain.DEFAULT_DOMAIN
	rgroup.Namespace = namespace.DEFAULT_NAMESPACE
	rgroup.Required = true
	rgroup.ValueType = VALUE_TYPE_ENUM
	rgroup.Visiable = resource.VISIABLE_GLOBAL
	rgroup.DefaultValue = "默认"
	rgroup.AddEnumOption(
		NewEnumOption("默认"),
	)
	items = append(items, rgroup)

	ugroup := NewCreateLabelRequest()
	ugroup.Visiable = resource.VISIABLE_GLOBAL
	ugroup.Key = USER_GROUP
	ugroup.KeyDesc = "用户组"
	ugroup.ValueDesc = "用户所属组"
	ugroup.Domain = domain.DEFAULT_DOMAIN
	ugroup.Namespace = namespace.DEFAULT_NAMESPACE
	ugroup.Required = false
	ugroup.ValueType = VALUE_TYPE_ENUM
	ugroup.Visiable = resource.VISIABLE_GLOBAL
	ugroup.DefaultValue = ""
	ugroup.AddEnumOption(
		NewEnumOption("研发部"),
		NewEnumOption("测试部"),
		NewEnumOption("运维部"),
	)
	items = append(items, ugroup)

	dgroup := NewCreateLabelRequest()
	dgroup.Visiable = resource.VISIABLE_GLOBAL
	dgroup.Key = DEPLOY_GROUP
	dgroup.KeyDesc = "部署组"
	dgroup.ValueDesc = "部署组"
	dgroup.Domain = domain.DEFAULT_DOMAIN
	dgroup.Namespace = namespace.DEFAULT_NAMESPACE
	dgroup.Required = true
	dgroup.ValueType = VALUE_TYPE_ENUM
	dgroup.Visiable = resource.VISIABLE_GLOBAL
	dgroup.DefaultValue = "default"
	// 色名 https://htmlcolorcodes.com/zh/yanse-ming/
	dgroup.AddEnumOption(
		NewEnumOption("green").SetLabel("绿组").SetColor("MediumSeaGreen"),
		NewEnumOption("blue").SetLabel("蓝组").SetColor("SteelBlue"),
	)
	items = append(items, dgroup)
	return
}
