package domain

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/pb/request"
	"github.com/rs/xid"
)

const (
	AppName = "domain"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewDefault todo
func NewDefault() *Domain {
	return &Domain{
		Spec: NewCreateDomainRequest(),
	}
}

// NewCreateDomainRequest todo
func NewCreateDomainRequest() *CreateDomainRequest {
	return &CreateDomainRequest{
		Profile:         &Profile{},
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

// Validate 校验请求是否合法
func (req *UpdateDomainRequest) Validate() error {
	return validate.Struct(req)
}

// NewDescribeDomainRequest 查询详情请求
func NewDescribeDomainRequest(id string) *DescribeDomainRequest {
	return &DescribeDomainRequest{
		DescribeBy: DESCRIBE_BY_ID,
		Id:         id,
	}
}

// NewPutDomainRequest todo
func NewPutDomainRequest(id string) *UpdateDomainRequest {
	return &UpdateDomainRequest{
		Id:         id,
		UpdateMode: request.UpdateMode_PUT,
		Spec:       &CreateDomainRequest{},
	}
}

// NewPatchDomainRequest todo
func NewPatchDomainRequest(id string) *UpdateDomainRequest {
	return &UpdateDomainRequest{
		Id:         id,
		UpdateMode: request.UpdateMode_PUT,
		Spec:       &CreateDomainRequest{},
	}
}

// LockedMiniteDuration todo
func (c *RetryLockConfig) LockedMiniteDuration() time.Duration {
	return time.Duration(c.LockedMinite) * time.Minute
}
