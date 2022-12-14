package code

import context "context"

type Service interface {
	RPCServer
	IssueCode(context.Context, *IssueCodeRequest) (*IssueCodeResponse, error)
}
