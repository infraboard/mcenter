package namespace

import context "context"

type Service interface {
	CreateNamespace(context.Context, *CreateNamespaceRequest) (*Namespace, error)
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*Namespace, error)
	RPCServer
}
