package rate_limiter

import (
	"fmt"
	"golang.org/x/time/rate"
	"learn-plus/proxy/middleware"
)

func RateLimiter() func(c *middleware.SliceRouterContext) {
	l := rate.NewLimiter(1, 2)
	return func(c *middleware.SliceRouterContext) {
		if !l.Allow() {
			c.Rw.Write([]byte(fmt.Sprintf("rate limit:%v,%v", l.Limit(), l.Burst())))
			c.Abort()
			return
		}
		c.Next()
	}
}