package http

type Config struct {
	Name            string
	MaxRetries      int
	BackoffDelaysMs int
}
