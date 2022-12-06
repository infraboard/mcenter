package domain

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"

	request "github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
)

const (
	AppName = "domain"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

func NewDefaultDomain() *Domain {
	return &Domain{
		Spec: NewCreateDomainRequest(),
	}
}

func (d *Domain) Desensitize() {

}

// NewCreateDomainRequest todo
func NewCreateDomainRequest() *CreateDomainRequest {
	return &CreateDomainRequest{
		SecuritySetting: NewDefaultSecuritySetting(),
	}
}

// NewDefaultSecuritySetting todo
func NewDefaultSecuritySetting() *SecuritySetting {
	return &SecuritySetting{
		PasswordSecurity: NewDefaulPasswordSecurity(),
		LoginSecurity:    NewDefaultLoginSecurity(),
	}
}

// NewDefaulPasswordSecurity todo
func NewDefaulPasswordSecurity() *PasswordSecurity {
	return &PasswordSecurity{
		Length:                  8,
		IncludeNumber:           true,
		IncludeLowerLetter:      true,
		IncludeUpperLetter:      false,
		IncludeSymbols:          false,
		RepeateLimite:           1,
		PasswordExpiredDays:     90,
		BeforeExpiredRemindDays: 10,
	}
}

// NewDefaultLoginSecurity todo
func NewDefaultLoginSecurity() *LoginSecurity {
	return &LoginSecurity{
		ExceptionLock: false,
		ExceptionLockConfig: &ExceptionLockConfig{
			OtherPlaceLogin: true,
			NotLoginDays:    30,
		},
		RetryLock: true,
		RetryLockConfig: &RetryLockConfig{
			RetryLimite:  5,
			LockedMinite: 30,
		},
		IpLimite: false,
		IpLimiteConfig: &IPLimiteConfig{
			Ip: []string{},
		},
	}
}

// New 新建一个domain
func New(req *CreateDomainRequest) (*Domain, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	d := &Domain{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
	}

	return d, nil
}

// Validate 校验请求是否合法
func (req *CreateDomainRequest) Validate() error {
	return validate.Struct(req)
}

func NewDescribeDomainRequestWithName(name string) *DescribeDomainRequest {
	return &DescribeDomainRequest{
		DescribeBy: DESCRIBE_BY_NAME,
		Name:       name,
	}
}

// Validate 校验请求是否合法
func (req *DescribeDomainRequest) Validate() error {
	switch req.DescribeBy {
	case DESCRIBE_BY_ID:
		if req.Id == "" {
			return fmt.Errorf("id required")
		}
	case DESCRIBE_BY_NAME:
		if req.Name == "" {
			return fmt.Errorf("name required")
		}
	}

	return validate.Struct(req)
}

func NewPatchPomainRequest(id string, req *CreateDomainRequest) *UpdateDomainRequest {
	return &UpdateDomainRequest{
		Id:   id,
		Spec: req,
	}
}

// Validate 校验请求是否合法
func (req *UpdateDomainRequest) Validate() error {
	if req.Id == "" && req.Name == "" {
		return fmt.Errorf("id or name required")
	}
	if req.UpdateMode.Equal(pb_request.UpdateMode_PUT) {
		return validate.Struct(req)
	}

	return nil
}

// NewDescribeDomainRequest 查询详情请求
func NewDescribeDomainRequestById(id string) *DescribeDomainRequest {
	return &DescribeDomainRequest{
		DescribeBy: DESCRIBE_BY_ID,
		Id:         id,
	}
}

// NewDescribeDomainRequest 查询详情请求
func NewDescribeDomainRequestByName(name string) *DescribeDomainRequest {
	return &DescribeDomainRequest{
		DescribeBy: DESCRIBE_BY_NAME,
		Name:       name,
	}
}

// NewPutDomainRequest todo
func NewPutDomainRequest(id string) *UpdateDomainRequest {
	return &UpdateDomainRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		Spec:       &CreateDomainRequest{},
	}
}

// NewPatchDomainRequest todo
func NewPatchDomainRequestById(id string) *UpdateDomainRequest {
	return &UpdateDomainRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		Spec:       &CreateDomainRequest{},
	}
}

// NewPatchDomainRequest todo
func NewPatchDomainRequestByName(name string) *UpdateDomainRequest {
	req := NewPatchDomainRequestById("")
	req.Name = name
	return req
}

// LockedMiniteDuration todo
func (c *RetryLockConfig) LockedMiniteDuration() time.Duration {
	return time.Duration(c.LockedMinite) * time.Minute
}

func NewDomainSet() *DomainSet {
	return &DomainSet{
		Items: []*Domain{},
	}
}

func (s *DomainSet) Add(item *Domain) {
	s.Items = append(s.Items, item)
}

func NewQueryDomainRequest() *QueryDomainRequest {
	return &QueryDomainRequest{
		Page:  request.NewPageRequest(20, 1),
		Ids:   []string{},
		Names: []string{},
	}
}

// NewQueryDomainRequestFromHTTP todo
func NewQueryDomainRequestFromHTTP(r *http.Request) *QueryDomainRequest {
	query := NewQueryDomainRequest()

	qs := r.URL.Query()
	query.Page = request.NewPageRequestFromHTTP(r)

	uids := qs.Get("domain_ids")
	if uids != "" {
		query.Ids = strings.Split(uids, ",")
	}

	dn := qs.Get("domain_names")
	if uids != "" {
		query.Names = strings.Split(dn, ",")
	}
	return query
}

func NewFeishuAccessToken() *FeishuAccessToken {
	return &FeishuAccessToken{
		IssueAt: time.Now().Unix(),
	}
}

func (c *FeishuConfig) MakeGetTokenFormRequest(code string) string {
	form := make(url.Values)
	form.Add("grant_type", "authorization_code")
	form.Add("client_id", c.AppId)
	form.Add("client_secret", c.AppSecret)
	form.Add("code", code)
	form.Add("redirect_uri", c.RedirectUri)
	return form.Encode()
}

func (t *FeishuAccessToken) IsExpired() bool {
	if t.AccessToken == "" {
		return true
	}

	// 为了避免误差, 再加30秒
	delta := time.Since(time.Unix(t.IssueAt, 0)).Seconds() + 30
	return delta > float64(t.ExpiresIn)
}
