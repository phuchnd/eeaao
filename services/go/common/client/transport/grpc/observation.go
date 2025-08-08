package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/avast/retry-go"
	"github.com/phuchnd/eeaao/services/go/common/observability/logging"
	"github.com/phuchnd/eeaao/services/go/common/observability/metrics"
	"github.com/phuchnd/eeaao/services/go/common/observability/tracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func propagateAndObservationUnaryClientInterceptor(cfg *Config, metricsExporter metrics.Metrics) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx := tracing.PropagateRequestIDToContext(ctx)
		var err error
		start := time.Now()
		defer func() {
			code := codes.OK
			if err != nil {
				code = status.Code(err)
			}
			metricsExporter.SendExternalServiceMetric(newCtx, start, cfg.ServiceName, cfg.ExternalServiceName, method, "", code.String())
		}()

		err = retry.Do(func() error {
			err := invoker(newCtx, method, req, reply, cc, opts...)
			if err != nil {
				logger := logging.FromContext(newCtx)
				logger.With("error", err).Warn(fmt.Sprintf("%s: inner attempt failed", method))
			}
			return err
		},
			retry.Attempts(uint(cfg.MaxRetries)),
			retry.Delay(time.Duration(cfg.BackoffDelaysMs)*time.Millisecond),
			retry.LastErrorOnly(true),
		)
		return err
	}
}
