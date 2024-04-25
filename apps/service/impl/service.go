package impl

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/service/provider/gitlab"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/pb/request"
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
				return nil, fmt.Errorf("list repo languages error, %s", err)
			}

			lan := languages.Primary()
			if lan != "" {
				lan, err := service.ParseLANGUAGEFromString(lan)
				if err != nil {
					return nil, fmt.Errorf("parse language from string error, %s", err)
				}
				repo.SetLanguage(lan)
			}
		}
		if repo.EnableHook {
			hookSetting, err := gitlab.ParseGitLabWebHookFromString(repo.HookConfig)
			if err != nil {
				return nil, fmt.Errorf("parse gitlab webhook from string error, %s", err)
			}
			hookSetting.Token = ins.Meta.Id
			hookSetting.Url = fmt.Sprintf("%s?service_id=%s", hookSetting.Url, ins.Meta.Id)

			addHookReq := gitlab.NewAddProjectHookRequest(repo.ProjectId, hookSetting)
			i.log.Debug().Msgf("add hook req: %s", addHookReq.ToJSON())
			resp, err := v4.Project().AddProjectHook(ctx, addHookReq)
			if err != nil {
				return nil, fmt.Errorf("add project hook error, %s", err)
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
	query := newQueryRequest(req)
	i.log.Debug().Msgf("find filter: %s", query.FindFilter())
	resp, err := i.col.Find(ctx, query.FindFilter(), query.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find book error, error is %s", err)
	}

	ServiceSet := service.NewServiceSet()
	// 循环
	for resp.Next(ctx) {
		ins := service.NewDefaultService()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode book error, error is %s", err)
		}

		ServiceSet.Add(ins.Desense())
	}

	// count
	count, err := i.col.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get Service count error, error is %s", err)
	}
	ServiceSet.Total = count
	return ServiceSet, nil
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

	return svcs.Sort(), nil
}

func (i *impl) DescribeService(ctx context.Context, req *service.DescribeServiceRequest) (
	*service.Service, error) {
	filter := bson.M{}

	switch req.DescribeBy {
	case service.DescribeBy_SERVICE_ID:
		filter["_id"] = req.Id
	case service.DescribeBy_SERVICE_CLIENT_ID:
		filter["credential.client_id"] = req.ClientId
	case service.DescribeBy_SERVICE_NAME:
		filter["name"] = req.Name
	}

	ins := service.NewDefaultService()
	if err := i.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("Service %s not found", req)
		}

		return nil, exception.NewInternalServerError("find Service %s error, %s", req, err)
	}

	return ins, nil

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
		if err == nil {
			v4 := gitlab.NewGitlabV4(gc)
			removeHookReq := gitlab.NewDeleteProjectHookReqeust(repo.ProjectId, repo.HookId)
			err = v4.Project().DeleteProjectHook(ctx, removeHookReq)
			if err != nil {
				i.log.Error().Msgf("delete project hook error, %s", err)
			}
		} else {
			i.log.Error().Msgf("new gitlab client error, %s", err)
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
