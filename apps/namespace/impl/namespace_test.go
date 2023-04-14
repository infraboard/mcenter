package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
)

func TestCreateNamespace(t *testing.T) {
	req := namespace.NewCreateNamespaceRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Name = namespace.DEFAULT_NAMESPACE
	req.Owner = "admin"
	r, err := impl.CreateNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestQueryNamespace(t *testing.T) {
	req := namespace.NewQueryNamespaceRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	r, err := impl.QueryNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestDescribeNamespace(t *testing.T) {
	req := namespace.NewDescriptNamespaceRequest(domain.DEFAULT_DOMAIN, namespace.DEFAULT_NAMESPACE)
	r, err := impl.DescribeNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
