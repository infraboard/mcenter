package label

import (
	"encoding/json"
	"fmt"

	"github.com/infraboard/mcenter/common/validate"
	"github.com/infraboard/mcube/pb/resource"
	"github.com/infraboard/mcube/tools/hash"
	"github.com/infraboard/mcube/tools/pretty"
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
	if err := in.Validate(); err != nil {
		return nil, err
	}

	ins := NewDefaultLabel()
	ins.Spec = in
	ins.Meta.Id = hash.FnvHash(in.UniqueKey())
	return ins, nil
}

func (r *CreateLabelRequest) Validate() error {
	return validate.Validate(r)
}

func (r *CreateLabelRequest) UniqueKey() string {
	switch r.Scope {
	case SCOPE_STSTEM:
		return r.Key
	case SCOPE_DOMAIN:
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
