package tracing

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"
)

func PropagateRequestIDToContext(ctx context.Context) context.Context {
	requestMetadata := FromContext(ctx)
	if requestMetadata == nil {
		return ctx
	}
	return metadata.AppendToOutgoingContext(ctx, DefaultContextKeyRequestID, requestMetadata.RequestID)
}

func PropagateRequestIDToHeader(ctx context.Context, outGoingHeader *http.Header) {
	requestMetadata := FromContext(ctx)
	if requestMetadata == nil {
		return
	}
	outGoingHeader.Set(DefaultContextKeyRequestID, requestMetadata.RequestID)
}
