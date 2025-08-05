package grpc

import (
	"context"
	"fmt"
	"github.com/phuchnd/eeaao/services/go/common/observability/tracing"
	"time"

	"github.com/avast/retry-go"
	"github.com/phuchnd/eeaao/services/go/common/observability/logging"
	"google.golang.org/grpc"
)

func propagateAndObservationUnaryClientInterceptor(cfg *Config) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx := tracing.PropagateRequestIDToContext(ctx)
		var err error

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
