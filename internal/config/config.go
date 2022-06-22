package config

import (
	"github.com/spf13/viper"
	"strings"
	"sync"
)

var configuration config
var configSync sync.Once

type config struct {
	Loglevel string `json:"loglevel"`
	Env      string `json:"env"`
}

func GetConfig() config {
	initConfig()
	return configuration
}

func initConfig() {
	configSync.Do(func() {
		viper.SetConfigName("default")
		viper.SetConfigType("yml")
		viper.AddConfigPath("./deployments/config")
		viper.AddConfigPath("/etc/demo")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			panic("Error reading config file " + err.Error())
		}
		err := viper.Unmarshal(&configuration)
		if err != nil {
			panic("Error reading config file " + err.Error())
		}
	})
}
