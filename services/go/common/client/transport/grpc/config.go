package grpc

type Config struct {
	Name            string
	Host            string
	Port            int
	MaxRetries      int
	BackoffDelaysMs int
}
