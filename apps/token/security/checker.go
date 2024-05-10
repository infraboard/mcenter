package security

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/cache"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/domain"
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
		domain: ioc.Controller().Get(domain.AppName).(domain.Service),
		user:   ioc.Controller().Get(user.AppName).(user.Service),
		token:  ioc.Controller().Get(token.AppName).(token.Service),
		cache:  c,
		log:    log.Sub("Login Security"),
	}, nil
}

type checker struct {
	domain domain.Service
	user   user.Service
	token  token.Service
	cache  cache.Cache
	log    *zerolog.Logger
}

func (c *checker) MaxFailedRetryCheck(ctx context.Context, req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySettingWithUser(ctx, req.Username)
	if !ss.RetryLock {
		c.log.Debug().Msgf("retry lock check disabled, don't check")
		return nil
	}
	c.log.Debug().Msgf("max failed retry lock check enabled, checking ...")

	var count uint32
	err := c.cache.Get(ctx, req.AbnormalUserCheckKey(), &count)
	if err != nil {
		c.log.Error().Msgf("get key %s from cache error, %s", req.AbnormalUserCheckKey(), err)
	}

	rc := ss.RetryLockConfig
	c.log.Debug().Msgf("retry times: %d, retry limite: %d", count, rc.RetryLimite)
	if count+1 >= rc.RetryLimite {
		return fmt.Errorf("登录失败次数过多, 请%d分钟后重试", rc.LockedMinite)
	}

	return nil
}

func (c *checker) UpdateFailedRetry(ctx context.Context, req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySettingWithUser(ctx, req.Username)
	if !ss.RetryLock {
		c.log.Debug().Msgf("retry lock check disabled, don't check")
		return nil
	}

	c.log.Debug().Msgf("update failed retry count, check key: %s", req.AbnormalUserCheckKey())

	var count int
	if err := c.cache.Get(ctx, req.AbnormalUserCheckKey(), &count); err == nil {
		// 之前已经登陆失败过
		_, err := c.cache.IncrBy(ctx, req.AbnormalUserCheckKey(), 1)
		if err != nil {
			c.log.Error().Msgf("set key %s to cache error, %s", req.AbnormalUserCheckKey(), err)
		}
	} else {
		// 首次登陆失败
		err := c.cache.Set(
			ctx,
			req.AbnormalUserCheckKey(),
			count+1,
			cache.WithExpiration(int64(ss.RetryLockConfig.LockedMinite)*60),
		)
		if err != nil {
			c.log.Error().Msgf("set key %s to cache error, %s", req.AbnormalUserCheckKey(), err)
		}
	}
	return nil
}

func (c *checker) OtherPlaceLoggedInChecK(ctx context.Context, tk *token.Token) error {
	ss := c.getOrDefaultSecuritySettingWithDomain(ctx, tk.Domain)
	if !ss.ExceptionLock {
		c.log.Debug().Msgf("exception check disabled, don't check")
		return nil
	}

	if !ss.ExceptionLockConfig.OtherPlaceLogin {
		c.log.Debug().Msgf("other place login check disabled, don't check")
		return nil
	}

	c.log.Debug().Msgf("other place login check enabled, checking ...")

	// 查询出用户上次登陆的地域
	queryReq := token.NewQueryUserWebLastToken(tk.UserId)
	lastTKSet, err := c.token.QueryToken(ctx, queryReq)
	if err != nil {
		return err
	}
	if lastTKSet.Length() == 0 {
		c.log.Debug().Msgf("last login session no ip info found, skip OtherPlaceLoggedInChecK")
		return nil
	}
	location := lastTKSet.Items[0].Location.IpLocation

	// 不错异地登录校验
	if !tk.Location.IpLocation.IsPublic() {
		c.log.Warn().Msgf("内网IP skip OtherPlaceLoggedInChecK")
		return nil
	}

	c.log.Debug().Msgf("user last login city: %s", location.City)
	if tk.Location.IpLocation.City != location.City {
		return fmt.Errorf("异地登录, 请输入验证码后再次提交")
	}

	return nil
}

func (c *checker) NotLoginDaysChecK(ctx context.Context, tk *token.Token) error {
	ss := c.getOrDefaultSecuritySettingWithUser(ctx, tk.Username)
	if !ss.ExceptionLock {
		c.log.Debug().Msgf("exception check disabled, don't check")
		return nil
	}
	c.log.Debug().Msgf("not login days check enabled, checking ...")

	// 查询出用户上次登陆的地域
	queryReq := token.NewQueryUserWebLastToken(tk.UserId)
	lastTKSet, err := c.token.QueryToken(ctx, queryReq)
	if err != nil {
		return err
	}
	if lastTKSet.Length() == 0 {
		c.log.Debug().Msgf("last login session no ip info found, skip OtherPlaceLoggedInChecK")
		return nil
	}
	ltk := lastTKSet.Items[0]

	days := uint32(time.Since(time.UnixMilli(ltk.IssueAt)).Hours() / 24)
	c.log.Debug().Msgf("user %d days not login", days)
	maxDays := ss.ExceptionLockConfig.NotLoginDays
	if days > maxDays {
		return fmt.Errorf("user not login days %d", days)
	}
	c.log.Debug().Msgf("not login days check passed, days: %d, max days: %d", days, maxDays)

	return nil
}

func (c *checker) IPProtectCheck(ctx context.Context, req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySettingWithUser(ctx, req.Username)
	if !ss.IpLimite {
		c.log.Debug().Msgf("ip limite check disabled, don't check")
		return nil
	}

	c.log.Debug().Msgf("ip limite check enabled, checking ...")

	return nil
}

func (c *checker) getOrDefaultSecuritySettingWithUser(ctx context.Context, username string) *domain.LoginSecurity {
	ss := domain.NewDefaultLoginSecurity()
	u, err := c.user.DescribeUser(ctx, user.NewDescriptUserRequestByName(username))
	if err != nil {
		c.log.Error().Msgf("get user error, %s, use default setting to check", err)
		return ss
	}

	return c.getOrDefaultSecuritySettingWithDomain(ctx, u.Spec.Domain)
}

func (c *checker) getOrDefaultSecuritySettingWithDomain(ctx context.Context, domainName string) *domain.LoginSecurity {
	ss := domain.NewDefaultLoginSecurity()
	d, err := c.domain.DescribeDomain(ctx, domain.NewDescribeDomainRequestWithName(domainName))
	if err != nil {
		c.log.Error().Msgf("get domain error, %s, use default setting to check", err)
		return ss
	}

	return d.Spec.LoginSecurity
}
