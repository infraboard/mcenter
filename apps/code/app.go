package code

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger/zap"
)

const (
	AppName = "code"
)

var (
	validate = validator.New()
)

// NewIssueCodeRequest todo
func NewIssueCodeRequest() *IssueCodeRequest {
	return &IssueCodeRequest{}
}

func (req *IssueCodeRequest) Validate() error {
	return validate.Struct(req)
}

// NewCode todo
func NewCode(req *IssueCodeRequest) (*Code, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate issue code request error, %s", err)
	}

	c := &Code{
		Code:          GenRandomCode(6),
		Username:      req.Username,
		IssueAt:       time.Now().UnixMilli(),
		ExpiredMinite: 10,
	}

	c.Id = HashID(c.Username, c.Code)
	return c, nil
}

// IsExpired todo
func (c *Code) IsExpired() bool {
	return time.Since(time.UnixMilli(c.IssueAt)).Minutes() > float64(c.ExpiredMinite)
}

// ExpiredMiniteString todo
func (c *Code) ExpiredMiniteString() string {
	return fmt.Sprintf("%d", c.ExpiredMinite)
}

func (c *Code) Key() string {
	return NewCodeKey(c.Username)
}

func NewCodeKey(username string) string {
	return fmt.Sprintf("%s.%s", CODE_ATTRIBUTE_NAME, username)
}

// GenRandomCode todo
func GenRandomCode(length uint) string {
	numbers := []string{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < int(length); i++ {
		c := rand.Intn(9)
		// 第一位不能为0
		if c == 0 {
			c = 1
		}

		numbers = append(numbers, strconv.Itoa(c))
	}

	return strings.Join(numbers, "")
}

// HashID todo
func HashID(username, code string) string {
	hash := fnv.New32a()
	hash.Write([]byte(username))
	hash.Write([]byte(code))
	return fmt.Sprintf("%x", hash.Sum32())
}

// NewIssueCodeResponse todo
func NewIssueCodeResponse(message string) *IssueCodeResponse {
	return &IssueCodeResponse{Message: message}
}

// Validate todo
func (req *VerifyCodeRequest) Validate() error {
	return validate.Struct(req)
}

// HashID todo
func (req *VerifyCodeRequest) HashID() string {
	hash := fnv.New32a()
	hash.Write([]byte(req.Username))
	hash.Write([]byte(req.Code))
	return fmt.Sprintf("%x", hash.Sum32())
}

// NewDefaultCode todo
func NewDefaultCode() *Code {
	return &Code{}
}

func NewVerifyCodeRequest(username, code string) *VerifyCodeRequest {
	return &VerifyCodeRequest{
		Username: username,
		Code:     code,
	}
}

// 优先从认证头中获取, 如果头没有从Cookie中获取
func GetCodeFromHTTP(r *http.Request) string {
	code := r.Header.Get(CODE_HEADER_KEY)
	if code != "" {
		return code
	}

	ck, err := r.Cookie(CODE_COOKIE_KEY)
	if err != nil {
		zap.L().Warnf("get code from cookie: %s error, %s", CODE_COOKIE_KEY, err)
		return ""
	}

	return ck.Value
}
