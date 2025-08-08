package metrics

import (
	"context"
	"time"
)

//go:generate mockery --name=Metrics --case=snake --disable-version-string
type Metrics interface {
	SendExternalServiceMetric(ctx context.Context, start time.Time, serviceName, externalServiceName, reqURL, reqMethod, respStatus string)
}

type metricsImpl struct{}

func NewMetrics() Metrics {
	return &metricsImpl{}
}

func (m *metricsImpl) SendExternalServiceMetric(ctx context.Context, start time.Time, serviceName, externalServiceName, reqURL, reqMethod, respStatus string) {
	//elapsed := time.Since(start)
	//msec := float64(elapsed.Nanoseconds()) / float64(time.Millisecond)
	//
}
