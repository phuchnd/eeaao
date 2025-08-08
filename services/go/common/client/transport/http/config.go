package http

type Config struct {
	ServiceName         string
	ExternalServiceName string
	MaxRetries          int
	BackoffDelaysMs     int
}
