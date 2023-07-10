package impl

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/service/provider/gitlab"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
)

func (i *impl) ValidateCredential(ctx context.Context, req *service.ValidateCredentialRequest) (
	*service.Service, error) {
	svr, err := i.DescribeService(ctx, service.NewDescribeServiceRequestByClientId(req.ClientId))
	if err != nil {
		return nil, err
	}

	if err := svr.Credential.Validate(req.ClientSecret); err != nil {
		return nil, err
	}

	return svr, nil
}

func (i *impl) CreateService(ctx context.Context, req *service.CreateServiceRequest) (
	*service.Service, error) {
	ins, err := service.NewService(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create book error, %s", err)
	}

	// 设置WebHook
	switch req.Type {
	case service.Type_SOURCE_CODE:
		repo := ins.Spec.CodeRepository
		if err := repo.Validate(); err != nil {
			return nil, exception.NewBadRequest("Repository validate failed, %s", err)
		}
		gc, err := repo.MakeGitlabConfig()
		if err != nil {
			return nil, err
		}
		v4 := gitlab.NewGitlabV4(gc)
		if repo.Language == nil {
			languages, err := v4.Project().ListProjectLanguage(ctx, repo.ProjectId)
			if err != nil {
				return nil, err
			}
			lan, err := service.ParseLANGUAGEFromString(languages.Primary())
			if err != nil {
				return nil, err
			}
			repo.SetLanguage(lan)
		}
		if repo.EnableHook {
			hookSetting, err := gitlab.ParseGitLabWebHookFromString(repo.HookConfig)
			if err != nil {
				return nil, err
			}
			hookSetting.Token = ins.Meta.Id
			addHookReq := gitlab.NewAddProjectHookRequest(repo.ProjectId, hookSetting)
			resp, err := v4.Project().AddProjectHook(ctx, addHookReq)
			if err != nil {
				return nil, err
			}
			repo.HookId = resp.IDToString()
		}
	}

	if err := i.save(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) UpdateService(ctx context.Context, req *service.UpdateServiceRequest) (
	*service.Service, error) {
	ins, err := i.DescribeService(ctx, service.NewDescribeServiceRequest(req.Id))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case request.UpdateMode_PUT:
		ins.Update(req)
	case request.UpdateMode_PATCH:
		err := ins.Patch(req)
		if err != nil {
			return nil, err
		}
	}

	// 校验更新后数据合法性
	if err := ins.Spec.Validate(); err != nil {
		return nil, err
	}

	if err := i.update(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) QueryService(ctx context.Context, req *service.QueryServiceRequest) (
	*service.ServiceSet, error) {
	query := newQueryRequest(req, i.log)
	return i.query(ctx, query)
}

func (i *impl) QueryGitlabProject(ctx context.Context, in *service.QueryGitlabProjectRequest) (
	*service.ServiceSet, error) {
	v4 := gitlab.NewGitlabV4(in.MakeConfig())
	pReq := gitlab.NewListProjectRequest()
	set, err := v4.Project().ListProjects(ctx, pReq)
	if err != nil {
		return nil, err
	}

	svcs := service.NewServiceSet()
	if set.Len() > 0 {
		gitSshUrls := set.GitSshUrls()
		query := service.NewQueryServiceRequest()
		query.Scope = in.Scope
		query.RepositorySshUrls = gitSshUrls
		query.Page.PageSize = uint64(len(gitSshUrls))
		svcs, err = i.QueryService(ctx, query)
		if err != nil {
			return nil, err
		}

		for i := range set.Items {
			p := set.Items[i]
			svcs.UpdateFromGitProject(p, in.Token)
		}
	}

	return svcs, nil
}

func (i *impl) DescribeService(ctx context.Context, req *service.DescribeServiceRequest) (
	*service.Service, error) {
	return i.get(ctx, req)
}

func (i *impl) DeleteService(ctx context.Context, req *service.DeleteServiceRequest) (
	*service.Service, error) {
	ins, err := i.DescribeService(ctx, service.NewDescribeServiceRequest(req.Id))
	if err != nil {
		return nil, err
	}

	// 如果开启了Hook需要移除Hook设置
	repo := ins.Spec.CodeRepository
	if repo.EnableHook {
		gc, err := repo.MakeGitlabConfig()
		if err != nil {
			return nil, err
		}
		v4 := gitlab.NewGitlabV4(gc)
		removeHookReq := gitlab.NewDeleteProjectHookReqeust(repo.ProjectId, repo.HookId)
		err = v4.Project().DeleteProjectHook(ctx, removeHookReq)
		if err != nil {
			return nil, err
		}
	}

	if err := i.delete(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) RefreshCredential(ctx context.Context, req *service.DescribeServiceRequest) (
	*service.Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshCredential not implemented")
}
