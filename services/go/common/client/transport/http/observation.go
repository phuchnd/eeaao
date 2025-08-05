package http

import (
	"context"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/phuchnd/eeaao/services/go/common/observability/logging"
	"net/http"
	"time"
)

// doFunc is an executable function which will return http status code and the error
type doFunc func(ctx context.Context) (int, error)

func (t *httpClientImpl) retryAndObserve(ctx context.Context, endpoint, method string, doFunc doFunc) error {
	start := time.Now()
	var responseCode int
	defer func() {
		code := http.StatusText(responseCode)
		opencensus.SendExternalServiceMetric(ctx, start, t.cfg.ServiceName, t.cfg.ExternalServiceName, endpoint, method, code)
	}()

	err := retry.Do(func() error {
		var err error
		responseCode, err = doFunc(ctx)
		if err != nil {
			logger := logging.FromContext(ctx)
			logger.Warnw(fmt.Sprintf("[%s] %s: inner attempt failed", method, endpoint), "error", err)
		}
		return err
	},
		retry.Attempts(uint(t.cfg.MaxRetries)),
		retry.Delay(time.Duration(t.cfg.BackoffDelaysMs)*time.Millisecond),
		retry.LastErrorOnly(true),
	)
	return err
}
