package metrics

import (
	"context"
	"time"
)

type Metrics interface {
	SendExternalServiceMetric(ctx context.Context, start time.Time, serviceName, externalServiceName, reqURL, reqMethod, respStatus string)
}

func SendExternalServiceMetric(ctx context.Context, start time.Time, serviceName, externalServiceName, reqURL, reqMethod, respStatus string) {
	//elapsed := time.Since(start)
	//msec := float64(elapsed.Nanoseconds()) / float64(time.Millisecond)
	//
}
