package label

const (
	AppName = "labels"
)

type Service interface {
	RPCServer
}

func NewCreateLabelRequest() *CreateLabelRequest {
	return &CreateLabelRequest{
		EnumOptions:    []*EnumOption{},
		HttpEnumConfig: NewHttpEnumConfig(),
	}
}

func NewHttpEnumConfig() *HttpEnumConfig {
	return &HttpEnumConfig{}
}
