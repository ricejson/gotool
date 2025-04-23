package retry

import (
	"errors"
	"sync/atomic"
	"time"
)

// MaxIntervalExp 最大重试间隔指数
const MaxIntervalExp = 10

// ExponentialIntervalStrategy 基于指数退避的重试策略实现
type ExponentialIntervalStrategy struct {
	initialInterval time.Duration // 起始重试时间间隔
	maxRetries      int32         // 最大重试次数（<=0表示无限重试）
	curRetries      int32         // 当前重试次数
}

func NewExponentialIntervalStrategy(initialInterval time.Duration, maxRetries int32) (*ExponentialIntervalStrategy, error) {
	if initialInterval <= 0 {
		return nil, errors.New("invalid retry initial interval")
	}
	return &ExponentialIntervalStrategy{
		initialInterval: initialInterval,
		maxRetries:      maxRetries,
		curRetries:      0,
	}, nil
}

// calculateCurrentInterval 计算当前时间间隔
func (s *ExponentialIntervalStrategy) calculateCurrentInterval() time.Duration {
	return s.initialInterval * (1 << min(MaxIntervalExp, atomic.LoadInt32(&s.curRetries)))
}

func (s *ExponentialIntervalStrategy) Next() (time.Duration, bool) {
	var curRetriesVal = atomic.LoadInt32(&s.curRetries)
	if s.maxRetries <= 0 {
		defer func() {
			// 重试次数+1
			atomic.AddInt32(&s.curRetries, 1)
		}()
		return s.calculateCurrentInterval(), true
	}
	for curRetriesVal < s.maxRetries {
		if !atomic.CompareAndSwapInt32(&s.curRetries, curRetriesVal, curRetriesVal+1) {
			curRetriesVal = atomic.LoadInt32(&s.curRetries)
			continue
		}
		return s.calculateCurrentInterval(), true
	}
	// 超出最大重试次数
	return 0, false
}
