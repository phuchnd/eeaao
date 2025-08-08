package tracing

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

// DefaultContextKeyRequestID is an envoy-specific header that is used to consistently sample logs and traces.
const DefaultContextKeyRequestID = "X-Request-Id"

type contextKey struct{}

type RequestTracing struct {
	// More params like TraceID, SpanID
	RequestID string
}

func NewMetadataFromGeneralContext(ctx context.Context) *RequestTracing {
	reqID := ""
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if mdRequestID := md.Get(DefaultContextKeyRequestID); len(mdRequestID) > 0 {
			reqID = mdRequestID[0]
		}
	}
	if reqID == "" {
		reqID = uuid.New().String()
	}
	return &RequestTracing{
		RequestID: reqID,
	}
}

// NewContext returns a new Context with given metadata.
func NewContext(ctx context.Context, meta *RequestTracing) context.Context {
	return context.WithValue(ctx, contextKey{}, meta)
}

// FromContext returns the tracing metadata associated with a context.
// It returns the nil if no RequestTracing exists.
func FromContext(ctx context.Context) *RequestTracing {
	if logger, ok := ctx.Value(contextKey{}).(*RequestTracing); ok {
		return logger
	}
	return nil
}

// FromGinContext returns the tracing metadata associated with a gin context.
func FromGinContext(c *gin.Context) *RequestTracing {
	reqID := c.GetHeader(DefaultContextKeyRequestID)
	// More value from header here
	if reqID == "" {
		reqID = uuid.New().String()
	}
	return &RequestTracing{
		RequestID: reqID,
	}
}
