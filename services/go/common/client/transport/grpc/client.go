package grpc

import (
	"fmt"

	"github.com/phuchnd/eeaao/services/go/common/observability/metrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRCPClientConn(cfg *Config, metricsExporter metrics.Metrics) (conn *grpc.ClientConn, err error) {
	return grpc.NewClient(
		fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			// propagate header, retry and sending external metrics count
			propagateAndObservationUnaryClientInterceptor(cfg, metricsExporter),
		),
	)
}
