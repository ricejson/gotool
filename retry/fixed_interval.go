package retry

import (
	"errors"
	"sync/atomic"
	"time"
)

// FixedIntervalStrategy 固定间隔重试策略
type FixedIntervalStrategy struct {
	maxRetries int32         // 最大重试次数（<=0表示无限重试）
	curRetries int32         // 当前重试次数
	interval   time.Duration // 重试间隔
}

func NewFixedIntervalStrategy(maxRetries int32, interval time.Duration) (*FixedIntervalStrategy, error) {
	if interval <= 0 {
		return nil, errors.New("invalid retry interval")
	}
	return &FixedIntervalStrategy{
		interval:   interval,
		maxRetries: maxRetries,
		curRetries: 0,
	}, nil
}

func (s *FixedIntervalStrategy) Next() (time.Duration, bool) {
	if s.maxRetries <= 0 {
		// 可以继续重试
		atomic.AddInt32(&s.curRetries, 1)
		return s.interval, true
	}
	// 判断是否超时最大时间间隔
	var curRetriesVal = atomic.LoadInt32(&s.curRetries)
	for curRetriesVal < s.maxRetries {
		if !atomic.CompareAndSwapInt32(&s.curRetries, curRetriesVal, curRetriesVal+1) {
			curRetriesVal = atomic.LoadInt32(&s.curRetries)
			continue
		}
		return s.interval, true
	}

	// 超出最大重试次数
	return 0, false
}
