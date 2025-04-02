package redisx

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redis/go-redis/v9"
	"net"
	"strconv"
	"time"
)

// PrometheusHook redis接入prometheus
type PrometheusHook struct {
	summaryVec *prometheus.SummaryVec
}

func NewPrometheusHook(summaryOpts prometheus.SummaryOpts) *PrometheusHook {
	summaryVec := prometheus.NewSummaryVec(summaryOpts, []string{"cmd", "keyExists"})
	prometheus.MustRegister(summaryVec)
	return &PrometheusHook{
		summaryVec: summaryVec,
	}
}

func (p *PrometheusHook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}

func (p *PrometheusHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		now := time.Now()
		var keyExists = true
		defer func() {
			duration := time.Since(now)
			p.summaryVec.WithLabelValues(cmd.Name(), strconv.FormatBool(keyExists)).Observe(float64(duration.Milliseconds()))
		}()
		err := next(ctx, cmd)
		keyExists = err != redis.Nil
		return err
	}
}

func (p *PrometheusHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		return next(ctx, cmds)
	}
}
