package auth

import (
	"context"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewHttpAuther() *httpAuther {
	return &httpAuther{
		log:            zap.L().Named("auther.http"),
		tk:             app.GetInternalApp(token.AppName).(token.Service),
		code:           app.GetInternalApp(code.AppName).(code.Service),
		cache:          cache.C(),
		codeCheckSlice: 30 * time.Minute,
	}
}

type httpAuther struct {
	log            logger.Logger
	tk             token.Service
	code           code.Service
	cache          cache.Cache
	codeCheckSlice time.Duration
}

// 设置静默时长
func (a *httpAuther) SetCodeCheckSliceTime(t time.Duration) {
	a.codeCheckSlice = t
}

func (a *httpAuther) GoRestfulAuthFunc(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {
	// 请求拦截
	entry := endpoint.NewEntryFromRestRequest(req)

	if entry != nil && entry.AuthEnable {
		// 访问令牌校验
		tk, err := a.CheckAccessToken(req)
		if err != nil {
			response.Failed(resp, err)
			return
		}

		// 接口调用权限校验
		err = a.CheckPermission(req.Request.Context(), tk, entry)
		if err != nil {
			response.Failed(resp, err)
			return
		}

		// 验证码校验(双因子认证)
		if a.IsCodeCheckSlice(tk.Username) && entry.CodeEnable {
			_, err := a.CheckCode(req, tk)
			if err != nil {
				response.Failed(resp, err)
				return
			}
		}
	}

	next.ProcessFilter(req, resp)
}

func (a *httpAuther) CheckAccessToken(req *restful.Request) (*token.Token, error) {
	// 获取用户Token, Token放在Heander Authorization
	ak := token.GetAccessTokenFromHTTP(req.Request)

	if ak == "" {
		return nil, token.ErrUnauthorized
	}

	// 调用GRPC 校验用户Token合法性
	tk, err := a.tk.ValidateToken(req.Request.Context(), token.NewValidateTokenRequest(ak))
	if err != nil {
		return nil, err
	}

	// 是不是需要返回用户的认证信息: 那个人, 那个空间下面， token本身的信息
	req.SetAttribute(token.TOKEN_ATTRIBUTE_NAME, tk)
	return tk, nil
}

func (a *httpAuther) CheckPermission(ctx context.Context, tk *token.Token, e *endpoint.Entry) error {
	// 检查是否是允许的类型
	if len(e.Allow) > 0 {
		a.log.Debugf("[%s] start check permission to keyauth ...", tk.Username)
		if !e.IsAllow(tk.UserType) {
			return exception.NewPermissionDeny("no permission, allow: %s, but current: %s", e.Allow, tk.UserType)
		}
		a.log.Debugf("[%s] permission check passed", tk.Username)
	}
	return nil
}

func (a *httpAuther) CheckCode(req *restful.Request, tk *token.Token) (*code.Code, error) {
	// 获取用户Code, Code放在Heander X-MCENTER-CODE
	cdStr := code.GetCodeFromHTTP(req.Request)
	if cdStr == "" {
		return nil, code.ErrUnauthorized
	}

	// 调用GRPC 校验用户Code合法性
	cd, err := a.code.VerifyCode(req.Request.Context(), code.NewVerifyCodeRequest(tk.Username, cdStr))
	if err != nil {
		return nil, err
	}

	// 保存返回的Code信息
	req.SetAttribute(code.CODE_ATTRIBUTE_NAME, cd)
	// 加入静默池中
	a.SetCodeCheckSlice(cd)
	return nil, nil
}

func (a *httpAuther) SetCodeCheckSlice(c *code.Code) {
	err := a.cache.PutWithTTL(c.Key(), c.Code, a.codeCheckSlice)
	if err != nil {
		a.log.Errorf("set code slice to cache error, %s", err)
	}
}

func (a *httpAuther) IsCodeCheckSlice(username string) bool {
	return a.cache.IsExist(code.NewCodeKey(username))
}
