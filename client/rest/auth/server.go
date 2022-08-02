package auth

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	httpb "github.com/infraboard/mcube/pb/http"
	"github.com/rs/xid"

	"github.com/infraboard/mcenter/apps/permission"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcenter/client/rpc"
)

const (
	OAuthTokenHeader = "X-OAUTH-TOKEN"
	RequestIdHeader  = "X-REQUEST-ID"
)

type PermissionCheckMode int

const (
	// PRBAC_MODE 基于策略的权限校验
	PRBAC_MODE PermissionCheckMode = 1
	// ACL_MODE 基于用户类型的权限校验
	ACL_MODE = 2
)

// NewInternalAuther 内部使用的auther
func NewHTTPAuther(c *rpc.ClientSet) *HTTPAuther {
	return &HTTPAuther{
		client: c,
		l:      zap.L().Named("Http Interceptor"),
		mode:   PRBAC_MODE,
		allows: []string{},
	}
}

// internal todo
type HTTPAuther struct {
	l      logger.Logger
	client *rpc.ClientSet
	mode   PermissionCheckMode
	svr    *service.Service
	lock   sync.Mutex
	allows []string
}

func (a *HTTPAuther) SetPermissionCheckMode(m PermissionCheckMode) {
	a.mode = m
}

func (a *HTTPAuther) SetAllows(allows ...fmt.Stringer) {
	for _, v := range allows {
		a.allows = append(a.allows, v.String())
	}
}

func (a *HTTPAuther) Auth(r *http.Request, entry httpb.Entry) (
	authInfo interface{}, err error) {
	var tk *token.Token

	// 从请求中获取access token
	acessToken := r.Header.Get(OAuthTokenHeader)

	if entry.AuthEnable {
		// 校验身份
		tk, err = a.ValidateIdentity(r.Context(), acessToken)
		if err != nil {
			return nil, err
		}

		// namesapce检查
		if entry.RequiredNamespace && tk.Namespace == "" {
			return nil, exception.NewBadRequest("namespace required!")
		}

		// 权限检查
		if entry.PermissionEnable {
			err = a.CheckPermission(r.Context(), tk, entry)
			if err != nil {
				return nil, err
			}
		}
	}

	// 设置RequestID
	if r.Header.Get(RequestIdHeader) == "" {
		r.Header.Set(RequestIdHeader, xid.New().String())
	}

	return tk, nil
}

// Gin Auth Middleware
func (a *HTTPAuther) GinAuthHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取access token
		acessToken := c.GetHeader(OAuthTokenHeader)

		// 校验身份
		tk, err := a.ValidateIdentity(c.Request.Context(), acessToken)
		if err != nil {
			response.Failed(c.Writer, err)
			return
		}
		c.Set("token", tk)
		c.Next()
	}
}

// Gin Perm Middleware
func (a *HTTPAuther) GinPermHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := c.MustGet("token")

		tk, ok := obj.(*token.Token)
		if !ok {
			response.Failed(c.Writer, fmt.Errorf("auth middleware first"))
			return
		}

		e := httpb.Entry{
			Method: c.Request.Method,
			Path:   c.FullPath(),
			Allow:  a.allows,
		}

		// 权限检查
		err := a.CheckPermission(c.Request.Context(), tk, e)
		if err != nil {
			response.Failed(c.Writer, err)
			return
		}
		c.Next()
	}
}

func (a *HTTPAuther) ValidateIdentity(ctx context.Context, accessToken string) (*token.Token, error) {
	a.l.Debug("start token identity check ...")

	if accessToken == "" {
		return nil, exception.NewBadRequest("token required")
	}

	req := token.NewValidateTokenRequestByAccessToken(accessToken)
	tk, err := a.client.Token().ValidateToken(ctx, req)
	if err != nil {
		return nil, err
	}

	a.l.Debugf("token check ok, username: %s", tk.Username)
	return tk, nil
}

func (a *HTTPAuther) CheckPermission(ctx context.Context, tk *token.Token, e httpb.Entry) error {
	if tk == nil {
		return exception.NewUnauthorized("validate permission need token")
	}

	// 如果是超级管理员不做权限校验, 直接放行
	if tk.UserType.IsIn(user.TYPE_SUPPER) {
		a.l.Debugf("[%s] supper admin skip permission check!", tk.Username)
		return nil
	}

	switch a.mode {
	case ACL_MODE:
		return a.ValidatePermissionByACL(ctx, tk, e)
	case PRBAC_MODE:
		return a.ValidatePermissionByPRBAC(ctx, tk, e)
	default:
		return fmt.Errorf("only support acl and prbac")
	}
}

func (a *HTTPAuther) ValidatePermissionByACL(ctx context.Context, tk *token.Token, e httpb.Entry) error {
	// 检查是否是允许的类型
	if len(e.Allow) > 0 {
		a.l.Debugf("[%s] start check permission to keyauth ...", tk.Username)
		if !e.IsAllow(tk.UserType) {
			return exception.NewPermissionDeny("no permission, allow: %s, but current: %s", e.Allow, tk.UserType)
		}
		a.l.Debugf("[%s] permission check passed", tk.Username)
	}

	return nil
}

func (a *HTTPAuther) ValidatePermissionByPRBAC(ctx context.Context, tk *token.Token, e httpb.Entry) error {
	svr, err := a.GetClientService(ctx)
	if err != nil {
		return err
	}

	req := permission.NewCheckPermissionRequest()
	req.Username = tk.Username
	req.Namespace = tk.Namespace
	req.ServiceId = svr.Id
	req.Path = e.UniquePath()
	_, err = a.client.Permission().CheckPermission(ctx, req)
	if err != nil {
		return exception.NewPermissionDeny(err.Error())
	}
	a.l.Debugf("[%s] permission check passed", tk.Username)
	return nil
}

func (a *HTTPAuther) GetClientService(ctx context.Context) (*service.Service, error) {
	if a.svr != nil {
		return a.svr, nil
	}
	a.lock.Lock()
	defer a.lock.Unlock()

	req := service.NewDescribeServiceRequestByClientId(a.client.GetClientID())
	ins, err := a.client.Service().DescribeService(ctx, req)
	if err != nil {
		return nil, err
	}
	a.svr = ins
	return ins, nil
}

// SetLogger todo
func (a *HTTPAuther) SetLogger(l logger.Logger) {
	a.l = l
}
