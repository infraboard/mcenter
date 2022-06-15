package user

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	request "github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

const (
	AppName = "user"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// New 实例
func New(req *CreateUserRequest) (*User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	pass, err := NewHashedPassword(req.Password)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	u := &User{
		Id:            xid.New().String(),
		CreateAt:      time.Now().UnixMilli(),
		Spec:          req,
		Password:      pass,
		Profile:       &Profile{},
		IsInitialized: false,
		Status: &Status{
			Locked: false,
		},
	}

	return u, nil
}

// NewHashedPassword 生产hash后的密码对象
func NewHashedPassword(password string) (*Password, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	return &Password{
		Password: string(bytes),
		CreateAt: time.Now().UnixMilli(),
	}, nil
}

// Validate 校验请求是否合法
func (req *CreateUserRequest) Validate() error {
	return validate.Struct(req)
}

// SetNeedReset 需要被重置
func (p *Password) SetNeedReset(format string, a ...interface{}) {
	p.NeedReset = true
	p.ResetReason = fmt.Sprintf(format, a...)
}

// NewCreateUserRequest 创建请求
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{}
}

// NewQueryUserRequestFromHTTP todo
func NewQueryUserRequestFromHTTP(r *http.Request) *QueryUserRequest {
	query := NewQueryUserRequest()

	qs := r.URL.Query()
	query.Page = request.NewPageRequestFromHTTP(r)
	query.Keywords = qs.Get("keywords")
	query.SkipItems = qs.Get("skip_items") == "true"

	uids := qs.Get("user_ids")
	if uids != "" {
		query.UserIds = strings.Split(uids, ",")
	}
	return query
}

// NewQueryUserRequest 列表查询请求
func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{
		Page:      request.NewPageRequest(20, 1),
		SkipItems: false,
	}
}

// NewDescriptUserRequestWithId 查询详情请求
func NewDescriptUserRequestWithId(id string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy: DESCRIBE_BY_USER_ID,
		Id:         id,
	}
}

// NewPatchAccountRequest todo
func NewPutUserRequest(userId string) *UpdateUserRequest {
	return &UpdateUserRequest{
		UserId:     userId,
		UpdateMode: pb_request.UpdateMode_PUT,
		Profile:    NewProfile(),
	}
}

// NewPatchAccountRequest todo
func NewPatchUserRequest(userId string) *UpdateUserRequest {
	return &UpdateUserRequest{
		UserId:     userId,
		UpdateMode: pb_request.UpdateMode_PATCH,
		Profile:    NewProfile(),
	}
}

// NewProfile todo
func NewProfile() *Profile {
	return &Profile{}
}

func NewDeleteUserRequest() *DeleteUserRequest {
	return &DeleteUserRequest{
		UserIds: []string{},
	}
}

func NewResetPasswordRequest() *ResetPasswordRequest {
	return &ResetPasswordRequest{}
}

func NewUpdatePasswordRequest() *UpdatePasswordRequest {
	return &UpdatePasswordRequest{}
}
