package deploy

func NewDeploySet() *DeploySet {
	return &DeploySet{
		Items: []*Deploy{},
	}
}

func (s *DeploySet) Add(item *Deploy) {
	s.Items = append(s.Items, item)
}

func NewDefaultDeploy() *Deploy {
	return &Deploy{
		Spec: NewCreateDeployRequest(),
	}
}
