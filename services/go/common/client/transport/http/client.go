package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	commonerrs "github.com/phuchnd/eeaao/services/go/common/errors"
	"github.com/phuchnd/eeaao/services/go/common/observability/logging"
	"github.com/phuchnd/eeaao/services/go/common/observability/metrics"
	"github.com/phuchnd/eeaao/services/go/common/observability/tracing"
)

//go:generate mockery --name=Client --case=snake --disable-version-string
type Client interface {
	Do(ctx context.Context, method, reqURL string, req []byte, resp interface{}, headers map[string]string) (int, error)
	GET(ctx context.Context, reqURL string, resp interface{}, headers map[string]string) (int, error)
	POST(ctx context.Context, reqURL string, req []byte, resp interface{}, headers map[string]string) (int, error)
}

type httpClientImpl struct {
	cfg             *Config
	client          *http.Client
	metricsExporter metrics.Metrics
}

func NewHTTPClient(cfg *Config, metricsExporter metrics.Metrics) Client {
	return &httpClientImpl{
		cfg:             cfg,
		client:          &http.Client{},
		metricsExporter: metricsExporter,
	}
}

func (t *httpClientImpl) Do(ctx context.Context, method string, reqURL string, req []byte, resp interface{}, headers map[string]string) (responseCode int, err error) {
	requestBody := strings.NewReader(string(req))
	var httpReq *http.Request
	httpReq, err = http.NewRequest(method, reqURL, requestBody)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if err := t.retryAndObserve(ctx, httpReq, func(ctx context.Context) (int, error) {
		responseCode, err = t.sendRequest(ctx, httpReq, &resp, headers)
		return responseCode, err
	}); err != nil {
		logger := logging.FromContext(ctx)
		logger.Errorw(fmt.Sprintf("[%s] %s failed with status code %d", method, httpReq.URL.Path, responseCode), "err", err)
		return responseCode, err
	}
	return responseCode, err
}

func (t *httpClientImpl) GET(ctx context.Context, reqURL string, resp interface{}, headers map[string]string) (responseCode int, err error) {
	var httpReq *http.Request
	httpReq, err = http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if err := t.retryAndObserve(ctx, httpReq, func(ctx context.Context) (int, error) {
		responseCode, err = t.sendRequest(ctx, httpReq, &resp, headers)
		return responseCode, err
	}); err != nil {
		logger := logging.FromContext(ctx)
		logger.Errorw(fmt.Sprintf("[GET] %s failed with status code %d", httpReq.URL.Path, responseCode), "err", err)
		return responseCode, err
	}
	return responseCode, nil
}

func (t *httpClientImpl) POST(ctx context.Context, reqURL string, req []byte, resp interface{}, headers map[string]string) (responseCode int, err error) {
	requestBody := strings.NewReader(string(req))
	var httpReq *http.Request
	httpReq, err = http.NewRequest(http.MethodPost, reqURL, requestBody)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if err := t.retryAndObserve(ctx, httpReq, func(ctx context.Context) (int, error) {
		responseCode, err = t.sendRequest(ctx, httpReq, &resp, headers)
		return responseCode, err
	}); err != nil {
		logger := logging.FromContext(ctx)
		logger.Errorw(fmt.Sprintf("[POST] %s failed with status code %d", httpReq.URL.Path, responseCode), "err", err)
		return responseCode, err
	}
	return responseCode, nil
}

func (t *httpClientImpl) sendRequest(ctx context.Context, req *http.Request, resp interface{}, headers map[string]string) (httpRespCode int, err error) {
	// Append request_id to the out going header
	tracing.PropagateRequestIDToHeader(ctx, &req.Header)

	// Set default header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Set custom header
	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}

	req = req.WithContext(ctx)

	logger := logging.FromContext(ctx)

	httpResp, err := t.client.Do(req)
	if err != nil {
		logger.Errorw(fmt.Sprintf("[%s] %s failed", req.Method, req.URL.Path), "err", err)
		return http.StatusInternalServerError, errors.Join(
			fmt.Errorf("[%s] %s failed", req.Method, req.URL.Path),
			err,
		)
	}
	defer func() {
		_ = httpResp.Body.Close()
	}()

	httpRespCode = httpResp.StatusCode
	httpRespBody, _ := io.ReadAll(httpResp.Body)

	if httpRespCode < http.StatusOK || httpRespCode >= http.StatusBadRequest {
		var errRes error
		var errBody *errorResponse

		if err = json.Unmarshal(httpRespBody, &errBody); err != nil {
			errRes = errors.Join(
				fmt.Errorf("[%s] %s got unexpected error code %d", req.Method, req.URL.Path, httpRespCode),
				errBody,
			)
		} else {
			errRes = errors.Join(
				commonerrs.ErrUnknown,
				fmt.Errorf("[%s] %s got unexpected error code %d", req.Method, req.URL.Path, httpRespCode),
			)
		}
		logger.Errorw(fmt.Sprintf("[%s] %s got unexpected error code", req.Method, req.URL.Path), "err", errRes, "request_url", httpResp.Request.URL, "response_code", httpRespCode, "httpRespBody", string(httpRespBody))
		return httpRespCode, errRes
	}

	// Not parsing body in-case request have no content
	if httpRespCode != http.StatusNoContent && len(httpRespBody) != 0 {
		if err = json.Unmarshal(httpRespBody, &resp); err != nil {
			logger.Errorw(fmt.Sprintf("[%s] %s failed, unable to decode repose body", req.Method, req.URL.Path), "err", err, "request_url", httpResp.Request.URL, "response_code", httpRespCode)
			return httpRespCode, errors.Join(
				fmt.Errorf("[%s] %s failed, unable to decode repose body", req.Method, req.URL.Path),
				err,
			)
		}
	}

	logger.Infow(fmt.Sprintf("[%s] %s success", req.Method, req.URL.Path), "request_url", httpResp.Request.URL, "response_code", httpRespCode)
	return httpRespCode, nil
}

// errorResponse is a general error struct to parsing to get more error on the http response body
type errorResponse struct {
	Message string `json:"message"`
	Err     string `json:"error"`
}

func (e *errorResponse) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%s| %s", e.Err, e.Message)
}
