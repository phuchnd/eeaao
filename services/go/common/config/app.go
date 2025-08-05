package config

type AppType string

const (
	AppTypeHTTP AppType = "http"
	AppTypeGRPC AppType = "grpc"
)

type AppEnv string

const (
	AppEnvLocal AppEnv  = "local"
	AppEnvStg   AppEnv  = "stg"
	AppEnvPrd   AppType = "prod"
)

type AppConfig struct {
	Type AppType
	Port int
	Name string
	Env  AppEnv
}
