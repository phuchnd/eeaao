package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/avast/retry-go"
	"github.com/phuchnd/eeaao/services/go/common/observability/logging"
)

// doFunc is an executable function which will return http status code and the error
type doFunc func(ctx context.Context) (int, error)

func (t *httpClientImpl) retryAndObserve(ctx context.Context, httpReq *http.Request, doFunc doFunc) error {
	start := time.Now()
	var responseCode int
	defer func() {
		code := http.StatusText(responseCode)
		t.metricsExporter.SendExternalServiceMetric(ctx, start, t.cfg.ServiceName, t.cfg.ExternalServiceName, httpReq.URL.Path, httpReq.Method, code)
	}()

	err := retry.Do(func() error {
		var err error
		responseCode, err = doFunc(ctx)
		if err != nil {
			logger := logging.FromContext(ctx)
			logger.Warnw(fmt.Sprintf("[%s] %s: inner attempt failed", httpReq.Method, httpReq.URL.Path), "error", err)
		}
		return err
	},
		retry.Attempts(uint(t.cfg.MaxRetries)),
		retry.Delay(time.Duration(t.cfg.BackoffDelaysMs)*time.Millisecond),
		retry.LastErrorOnly(true),
	)
	return err
}
