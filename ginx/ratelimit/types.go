package ratelimit

import (
	"context"
	"time"
)

// RateLimiter 抽象限流器定义
type RateLimiter interface {
	Limit(ctx context.Context, key string, interval time.Duration, rate int) (bool, error)
}
