package retry

import "time"

// Strategy 重试策略接口定义
type Strategy interface {
	// Next 返回下一次重试的时间间隔
	// 如果返回false表示无法继续重试
	Next() (time.Duration, bool)
}
