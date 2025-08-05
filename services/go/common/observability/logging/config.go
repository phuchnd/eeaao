package logging

import (
	"github.com/phuchnd/eeaao/services/go/common/config"
	"github.com/phuchnd/eeaao/services/go/common/config/registry"
	"github.com/spf13/viper"
)

// ConfigName is the logging configuration name
const ConfigName = "logging"

type Config struct {
	IsDevelopment bool
	Level         string
}

func GetConfig(cp config.Provider) *Config {
	return cp.Get(ConfigName).(*Config)
}

func init() {
	registry.RegisterConfig(ConfigName, registry.NewConfig(func(v *viper.Viper) interface{} {
		return &Config{
			IsDevelopment: v.GetBool(ConfigName + ".is_development"),
			Level:         v.GetString(ConfigName + ".level"),
		}
	}, registry.WithSetDefault(func(v *viper.Viper) {
		v.SetDefault(ConfigName, map[string]interface{}{
			"is_development": true,
			"level":          "debug",
		})
	})))
}
