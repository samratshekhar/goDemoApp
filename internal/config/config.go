package config

import (
	"github.com/spf13/viper"
	"strings"
	"sync"
)

var configuration config
var configSync sync.Once

type config struct {
	Loglevel         string           `mapstructure:"LOG_LEVEL"`
	Env              string           `mapstructure:"ENVIRONMENT"`
	HttpServerConfig httpServerConfig `mapstructure:"HTTP_SERVER_CONFIG"`
}

type httpServerConfig struct {
	Port               string `mapstructure:"PORT"`
	IdleTimeoutSeconds string `mapstructure:"IDLE_TIMEOUT_SECONDS"`
}

func GetConfig() config {
	initConfig()
	return configuration
}

func initConfig() {
	configSync.Do(func() {
		loadConfig(&configuration, "./deployments/config", "etc/demo")
	})
}

func loadConfig(config *config, paths ...string) {
	viper.SetConfigName("default")
	viper.SetConfigType("yml")
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file " + err.Error())
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		panic("Error reading config file " + err.Error())
	}
}
