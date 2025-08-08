package mysql

type Config struct {
	Host            string
	Port            int
	Username        string
	Password        string
	Database        string
	MaxIdleConns    int
	MaxOpenConns    int
	MaxRetries      int
	BackoffDelaysMs int
}
