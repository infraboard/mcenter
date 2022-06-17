package security

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/ip2region"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/user"
)

// NewChecker todo
func NewChecker() (Checker, error) {
	c := cache.C()
	if c == nil {
		return nil, fmt.Errorf("denpence cache service is nil")
	}

	return &checker{
		domain:    app.GetInternalApp(domain.AppName).(domain.Service),
		user:      app.GetInternalApp(user.AppName).(user.Service),
		token:     app.GetInternalApp(token.AppName).(token.Service),
		cache:     c,
		ip2Regoin: app.GetInternalApp(ip2region.AppName).(ip2region.Service),
		log:       zap.L().Named("Login Security"),
	}, nil
}

type checker struct {
	domain    domain.Service
	user      user.Service
	token     token.Service
	cache     cache.Cache
	ip2Regoin ip2region.Service
	log       logger.Logger
}

func (c *checker) MaxFailedRetryCheck(ctx context.Context, req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySettingWithUser(ctx, req.Username)
	if !ss.LoginSecurity.RetryLock {
		c.log.Debugf("retry lock check disabled, don't check")
		return nil
	}
	c.log.Debugf("max failed retry lock check enabled, checking ...")

	var count uint32
	err := c.cache.Get(req.AbnormalUserCheckKey(), &count)
	if err != nil {
		c.log.Errorf("get key %s from cache error, %s", req.AbnormalUserCheckKey(), err)
	}

	rc := ss.LoginSecurity.RetryLockConfig
	c.log.Debugf("retry times: %d, retry limite: %d", count, rc.RetryLimite)
	if count+1 >= rc.RetryLimite {
		return fmt.Errorf("登录失败次数过多, 请%d分钟后重试", rc.LockedMinite)
	}

	return nil
}

func (c *checker) UpdateFailedRetry(ctx context.Context, req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySettingWithUser(ctx, req.Username)
	if !ss.LoginSecurity.RetryLock {
		c.log.Debugf("retry lock check disabled, don't check")
		return nil
	}

	c.log.Debugf("update failed retry count, check key: %s", req.AbnormalUserCheckKey())

	var count int
	if err := c.cache.Get(req.AbnormalUserCheckKey(), &count); err == nil {
		// 之前已经登陆失败过
		err := c.cache.Put(req.AbnormalUserCheckKey(), count+1)
		if err != nil {
			c.log.Errorf("set key %s to cache error, %s", req.AbnormalUserCheckKey())
		}
	} else {
		// 首次登陆失败
		err := c.cache.PutWithTTL(
			req.AbnormalUserCheckKey(),
			count+1,
			ss.LoginSecurity.RetryLockConfig.LockedMiniteDuration(),
		)
		if err != nil {
			c.log.Errorf("set key %s to cache error, %s", req.AbnormalUserCheckKey())
		}
	}
	return nil
}

func (c *checker) OtherPlaceLoggedInChecK(ctx context.Context, tk *token.Token) error {
	ss := c.getOrDefaultSecuritySettingWithDomain(ctx, tk.Domain)
	if !ss.LoginSecurity.ExceptionLock {
		c.log.Debugf("exception check disabled, don't check")
		return nil
	}

	if !ss.LoginSecurity.ExceptionLockConfig.OtherPlaceLogin {
		c.log.Debugf("other place login check disabled, don't check")
		return nil
	}

	c.log.Debugf("other place login check enabled, checking ...")

	// 查询当前登陆IP地域
	rip := tk.Location.IpLocation.RemoteIp
	c.log.Debugf("query remote ip: %s location ...", rip)
	login, err := c.ip2Regoin.LookupIP(rip)
	if err != nil {
		c.log.Errorf("lookup ip %s error, %s, skip OtherPlaceLoggedInChecK", rip, err)
		return nil
	}

	// 查询出用户上次登陆的地域
	queryReq := token.NewQueryUserWebLastToken(tk.UserId)
	lastTKSet, err := c.token.QueryToken(ctx, queryReq)
	if err != nil {
		return err
	}
	if lastTKSet.Length() == 0 {
		c.log.Debugf("last login session no ip info found, skip OtherPlaceLoggedInChecK")
		return nil
	}
	location := lastTKSet.Items[0].Location.IpLocation

	// city为0 表示内网IP, 不错异地登录校验
	if login.CityID == 0 || location.CityId == 0 {
		c.log.Warnf("city id is 0, 内网IP skip OtherPlaceLoggedInChecK")
		return nil
	}

	c.log.Debugf("user last login city: %s (%d)", location.City, location.CityId)
	if login.CityID != location.CityId {
		return fmt.Errorf("异地登录, 请输入验证码后再次提交")
	}

	return nil
}

func (c *checker) NotLoginDaysChecK(ctx context.Context, tk *token.Token) error {
	ss := c.getOrDefaultSecuritySettingWithUser(ctx, tk.Username)
	if !ss.LoginSecurity.ExceptionLock {
		c.log.Debugf("exception check disabled, don't check")
		return nil
	}
	c.log.Debugf("not login days check enabled, checking ...")

	// 查询出用户上次登陆的地域
	queryReq := token.NewQueryUserWebLastToken(tk.UserId)
	lastTKSet, err := c.token.QueryToken(ctx, queryReq)
	if err != nil {
		return err
	}
	if lastTKSet.Length() == 0 {
		c.log.Debugf("last login session no ip info found, skip OtherPlaceLoggedInChecK")
		return nil
	}
	ltk := lastTKSet.Items[0]

	days := uint32(time.Since(time.UnixMilli(ltk.IssueAt)).Hours() / 24)
	c.log.Debugf("user %d days not login", days)
	maxDays := ss.LoginSecurity.ExceptionLockConfig.NotLoginDays
	if days > maxDays {
		return fmt.Errorf("user not login days %d", days)
	}
	c.log.Debugf("not login days check passed, days: %d, max days: %d", days, maxDays)

	return nil
}

func (c *checker) IPProtectCheck(ctx context.Context, req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySettingWithUser(ctx, req.Username)
	if !ss.LoginSecurity.IpLimite {
		c.log.Debugf("ip limite check disabled, don't check")
		return nil
	}

	c.log.Debugf("ip limite check enabled, checking ...")

	return nil
}

func (c *checker) getOrDefaultSecuritySettingWithUser(ctx context.Context, username string) *domain.SecuritySetting {
	ss := domain.NewDefaultSecuritySetting()
	u, err := c.user.DescribeUser(ctx, user.NewDescriptUserRequestWithName(username))
	if err != nil {
		c.log.Errorf("get user error, %s, use default setting to check", err)
		return ss
	}

	return c.getOrDefaultSecuritySettingWithDomain(ctx, u.Spec.Domain)
}

func (c *checker) getOrDefaultSecuritySettingWithDomain(ctx context.Context, domainName string) *domain.SecuritySetting {
	ss := domain.NewDefaultSecuritySetting()
	d, err := c.domain.DescribeDomain(ctx, domain.NewDescribeDomainRequestWithName(domainName))
	if err != nil {
		c.log.Errorf("get domain error, %s, use default setting to check", err)
		return ss
	}

	return d.Spec.SecuritySetting
}
