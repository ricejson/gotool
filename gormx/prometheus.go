package gormx

import (
	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
	"time"
)

type GORMCallback struct {
	summaryVector *prometheus.SummaryVec
}

func NewGORMCallback(opts prometheus.SummaryOpts) *GORMCallback {
	summaryVector := prometheus.NewSummaryVec(opts, []string{"cmdType"})
	prometheus.MustRegister(summaryVector)
	return &GORMCallback{
		summaryVector: summaryVector,
	}
}

func (c *GORMCallback) Before(tx *gorm.DB) {
	// 存入当前时间
	now := time.Now()
	tx.Set("startTime", now)
}

func (c *GORMCallback) After(cmdType string, tx *gorm.DB) {
	// 取出当前时间
	val, _ := tx.Get("startTime")
	startTime := val.(time.Time)
	duration := time.Since(startTime)
	c.summaryVector.WithLabelValues(cmdType).Observe(float64(duration.Milliseconds()))
}
