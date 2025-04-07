package ratelimit

import (
	"context"
	_ "embed"
	"github.com/redis/go-redis/v9"
	"time"
)

//go:embed redis_slide_window.lua
var luaScript string

// RedisSlideWindowRateLimiter 基于redis滑动窗口的限流器实现
type RedisSlideWindowRateLimiter struct {
	client redis.Cmdable
}

func NewRedisSliceWindowRateLimiter(cmd redis.Cmdable) RateLimiter {
	return &RedisSlideWindowRateLimiter{
		client: cmd,
	}
}

func (r *RedisSlideWindowRateLimiter) Limit(ctx context.Context, key string, interval time.Duration, rate int) (bool, error) {
	return r.client.Eval(ctx,
		luaScript, []string{key},
		interval.Milliseconds(), rate, time.Now().UnixMilli(),
	).Bool()
}
