package user

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
	"github.com/infraboard/mcenter/apps/domain"
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
		CreateAt:      time.Now().UnixMilli(),
		Spec:          req,
		Password:      pass,
		Profile:       &Profile{},
		IsInitialized: false,
		Status: &Status{
			Locked: false,
		},
	}

	if req.UseFullNamedUid {
		u.MakeFullNamedUid()
	} else {
		u.Id = xid.New().String()
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
		Password:      string(bytes),
		CreateAt:      time.Now().UnixMilli(),
		UpdateAt:      time.Now().UnixMilli(),
		ExpiredDays:   90,
		ExpiredRemind: 30,
	}, nil
}

// Validate 校验请求是否合法
func (req *CreateUserRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateUserRequest) GetFeishuUserId() string {
	if req.Feishu == nil {
		return ""
	}

	return req.Feishu.UserId
}

// SetNeedReset 需要被重置
func (p *Password) SetNeedReset(format string, a ...interface{}) {
	p.NeedReset = true
	p.ResetReason = fmt.Sprintf(format, a...)
}

// NewCreateUserRequest 创建请求
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		UseFullNamedUid: true,
	}
}

func NewLDAPCreateUserRequest(domain, username, password, descriptoin string) *CreateUserRequest {
	return &CreateUserRequest{
		Provider:    PROVIDER_LDAP,
		Type:        TYPE_SUB,
		CreateBy:    CREATE_BY_ADMIN,
		Domain:      domain,
		Username:    username,
		Password:    password,
		Description: descriptoin,
	}
}

func NewFeishuCreateUserRequest(domain, username, password, descriptoin string) *CreateUserRequest {
	return &CreateUserRequest{
		Provider:    PROVIDER_FEISHU,
		Type:        TYPE_SUB,
		CreateBy:    CREATE_BY_ADMIN,
		Domain:      domain,
		Username:    username,
		Password:    password,
		Description: descriptoin,
	}
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

// NewDescriptUserRequestWithId 查询详情请求
func NewDescriptUserRequestWithName(username string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy: DESCRIBE_BY_USER_NAME,
		Username:   username,
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
		UserId:      userId,
		UpdateMode:  pb_request.UpdateMode_PATCH,
		Profile:     NewProfile(),
		FeishuToken: NewFeishuAccessToken(),
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

// CheckPassword 判断password 是否正确
func (p *Password) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
	if err != nil {
		return exception.NewUnauthorized("user or password not connrect")
	}

	return nil
}

// CheckPasswordExpired 检测password是否已经过期
// remindDays 提前多少天提醒用户修改密码
// expiredDays 多少天后密码过期
func (p *Password) CheckPasswordExpired(remindDays, expiredDays uint) error {
	// 永不过期
	if expiredDays == 0 {
		return nil
	}

	now := time.Now()
	expiredAt := time.UnixMilli(p.UpdateAt).Add(time.Duration(expiredDays) * time.Hour * 24)

	ex := now.Sub(expiredAt).Hours() / 24
	if ex > 0 {
		return exception.NewPasswordExired("password expired %f days", ex)
	} else if ex >= -float64(remindDays) {
		p.SetNeedReset("密码%f天后过期, 请重置密码", -ex)
	}

	return nil
}

func NewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}

func (s *UserSet) Add(item *User) {
	s.Items = append(s.Items, item)
}

func (s *UserSet) HasUser(userId string) bool {
	for i := range s.Items {
		if s.Items[i].Id == userId {
			return true
		}
	}

	return false
}

func (s *UserSet) UserIds() (uids []string) {
	for i := range s.Items {
		uids = append(uids, s.Items[i].Id)
	}

	return
}

func NewDefaultUser() *User {
	return &User{}
}

// Desensitize 关键数据脱敏
func (u *User) Desensitize() {
	if u.Password != nil {
		u.Password.Password = ""
		u.Password.History = []string{}
	}
}

// Desensitize 关键数据脱敏
func (u *User) MakeFullNamedUid() {
	u.Id = fmt.Sprintf("%s@%s", u.Spec.Username, u.Spec.Domain)
}

func (i *User) Update(req *UpdateUserRequest) {
	i.UpdateAt = time.Now().UnixMicro()
	i.Profile = req.Profile
	i.Spec.Description = req.Description
	i.FeishuToken = req.FeishuToken
}

func (i *User) Patch(req *UpdateUserRequest) error {
	i.UpdateAt = time.Now().UnixMicro()
	err := mergo.MergeWithOverwrite(i.Profile, req.Profile)
	if err != nil {
		return err
	}

	if i.FeishuToken == nil {
		i.FeishuToken = NewFeishuAccessToken()
	}
	return mergo.MergeWithOverwrite(i.FeishuToken, req.FeishuToken)
}

func SpliteUserAndDomain(username string) (string, string) {
	kvs := strings.Split(username, "@")
	if len(kvs) > 1 {
		dom := strings.Join(kvs[1:], "")
		if dom == "" {
			dom = domain.DEFAULT_DOMAIN
		}
		return kvs[0], dom
	}

	return username, domain.DEFAULT_DOMAIN
}

// NewProfile todo
func NewFeishu() *Feishu {
	return &Feishu{}
}

func NewFeishuAccessToken() *FeishuAccessToken {
	return &FeishuAccessToken{
		IssueAt: time.Now().Unix(),
	}
}

func (t *FeishuAccessToken) IsExpired() bool {
	if t.AccessToken == "" {
		return true
	}

	// 为了避免误差, 再加30秒
	delta := time.Since(time.Unix(t.IssueAt, 0)).Seconds() + 30
	return delta > float64(t.ExpiresIn)
}

func NewDingDingAccessToken() *DingDingAccessToken {
	return &DingDingAccessToken{
		IssueAt: time.Now().Unix(),
	}
}
