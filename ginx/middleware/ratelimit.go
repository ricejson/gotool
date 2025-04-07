package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ricejson/gotool/ginx/ratelimit"
	"github.com/ricejson/gotool/stringx"
	"net/http"
	"time"
)

// RateLimiterBuilder 限流中间件构造器
type RateLimiterBuilder struct {
	limiter  ratelimit.RateLimiter
	prefix   string
	key      string
	interval time.Duration
	rate     int
}

func NewRateLimiterBuilder(limiter ratelimit.RateLimiter, interval time.Duration, rate int) *RateLimiterBuilder {
	return &RateLimiterBuilder{
		limiter:  limiter,
		interval: interval,
		rate:     rate,
		prefix:   "ip-limiter", // 针对ip进行限流
	}
}

func (builder *RateLimiterBuilder) Prefix(prefix string) *RateLimiterBuilder {
	builder.prefix = prefix
	return builder
}

func (builder *RateLimiterBuilder) Key(key string) *RateLimiterBuilder {
	builder.key = key
	return builder
}

func (builder *RateLimiterBuilder) generateKey(suffix string) string {
	// 构造key
	key := fmt.Sprintf("%s:%s", builder.prefix, suffix)
	return key
}

func (builder *RateLimiterBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var key string
		if stringx.IsEmpty(builder.key) {
			ip := ctx.ClientIP()
			// 使用默认ip作为后缀
			key = builder.generateKey(ip)
		} else {
			key = builder.generateKey(builder.key)
		}
		// 进行限流
		limit, err := builder.limiter.Limit(ctx, key, builder.interval, builder.rate)
		if err != nil {
			// 说明redis发生不可名状错误
			// 这里直接进行限流
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
			// 你也可以激进一点，继续执行业务代码（但要有兜底措施）
		}
		if limit {
			// 需要进行限流
			ctx.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		// 放行
		ctx.Next()
	}
}
