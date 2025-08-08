package grpc

type Config struct {
	ServiceName         string
	ExternalServiceName string
	Host                string
	Port                int
	MaxRetries          int
	BackoffDelaysMs     int
}
