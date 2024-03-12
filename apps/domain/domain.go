package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"

	notify "github.com/infraboard/mcenter/apps/notify"
	request "github.com/infraboard/mcube/v2/http/request"
	pb_request "github.com/infraboard/mcube/v2/pb/request"
	"github.com/infraboard/mcube/v2/tools/hash"
	"github.com/infraboard/mcube/v2/tools/pretty"
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

// New 新建一个domain
func New(req *CreateDomainRequest) (*Domain, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	d := &Domain{
		Meta: NewMeta(),
		Spec: req,
	}

	// 采用公司名称Hash作为Id
	d.Meta.Id = hash.FnvHash(req.Name)
	return d, nil
}

func NewMeta() *Meta {
	return &Meta{
		Id:       xid.New().String(),
		CreateAt: time.Now().Unix(),
	}
}

func (d *Domain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*Meta
		*CreateDomainRequest
	}{d.Meta, d.Spec})
}

func (d *Domain) ToJson() string {
	return pretty.ToJSON(d)
}

func (d *Domain) Desense() {
	if d.Spec.FeishuSetting != nil {
		d.Spec.FeishuSetting.DeSense()
	}
	if d.Spec.LdapSetting != nil {
		d.Spec.LdapSetting.DeSense()
	}
}

// NewCreateDomainRequest todo
func NewCreateDomainRequest() *CreateDomainRequest {
	return &CreateDomainRequest{
		PasswordConfig: NewDefaulPasswordSecurity(),
		LoginSecurity:  NewDefaultLoginSecurity(),
		CodeConfig:     NewDefaultCodeSetting(),
		NotifyConfig:   notify.NewNotifySetting(),
		FeishuSetting:  NewDefaultFeishuConfig(),
		LdapSetting:    NewDefaultLDAPConfig(),
	}
}

// NewDefaulPasswordSecurity todo
func NewDefaulPasswordSecurity() *PasswordConfig {
	return &PasswordConfig{
		Enabled:                 true,
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
			return fmt.Errorf("domain id required")
		}
	case DESCRIBE_BY_NAME:
		if req.Name == "" {
			return fmt.Errorf("domain name required")
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

// NewPutDomainRequestByName todo
func NewPutDomainRequestByName(name string) *UpdateDomainRequest {
	return &UpdateDomainRequest{
		Name:       name,
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

func (s *DomainSet) ToJson() string {
	return pretty.ToJSON(s)
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

func (c *FeishuConfig) MakeGetTokenFormRequest(code string) string {
	form := make(url.Values)
	form.Add("grant_type", "authorization_code")
	form.Add("client_id", c.AppId)
	form.Add("client_secret", c.AppSecret)
	form.Add("code", code)
	form.Add("redirect_uri", c.RedirectUri)
	return form.Encode()
}

// NewDefaultCodeSetting todo
func NewDefaultCodeSetting() *CodeSetting {
	return &CodeSetting{
		NotifyType:    notify.NOTIFY_TYPE_MAIL,
		ExpireMinutes: 10,
		MailTemplate:  "您的动态验证码为：{1}，{2}分钟内有效！，如非本人操作，请忽略本邮件！",
	}
}

// RenderMailCentent todo
func (c *CodeSetting) RenderMailCentent(code string, expireMinite uint32) string {
	// 如果为0 则使用默认值
	if expireMinite == 0 {
		expireMinite = c.ExpireMinutes
	}

	t1 := strings.ReplaceAll(c.MailTemplate, "{1}", code)
	return strings.ReplaceAll(t1, "{2}", fmt.Sprintf("%d", expireMinite))
}
