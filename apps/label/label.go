package label

import (
	"encoding/json"
	"fmt"

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
