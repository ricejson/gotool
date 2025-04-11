package retry

import (
	"context"
	"errors"
	"time"
)

// Retry 重试方法
// 该方法的退出条件需要满足以下三种条件之一
// 1. context超时或者取消
// 2. 超出最大重试次数
// 3. bizFunc业务代码没有返回error
func Retry(ctx context.Context, strategy Strategy, bizFunc func() error) error {
	var ticker *time.Ticker
	defer func() {
		if ticker != nil {
			ticker.Stop()
		}
	}()
	for {
		err := bizFunc()
		if err == nil {
			// 业务正常运行
			return nil
		}
		// 计算下次重试间隔
		nextTime, ok := strategy.Next()
		if !ok {
			// 说明重试次数耗尽
			return errors.New("retry count has been exhausted")
		}
		if ticker == nil {
			ticker = time.NewTicker(nextTime)
		} else {
			ticker.Reset(0)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}
	}
}
