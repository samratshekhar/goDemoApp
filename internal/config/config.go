package config

import (
	"github.com/spf13/viper"
	"strings"
	"sync"
)

var configuration config
var configSync sync.Once

type config struct {
	Loglevel         string              `mapstructure:"LOG_LEVEL"`
	Environment      string              `mapstructure:"ENVIRONMENT"`
	HTTPServerConfig httpServerConfig    `mapstructure:"HTTP_SERVER_CONFIG"`
	MySQLConfig      dbConfig            `mapstructure:"MYSQL_CONFIG"`
	Dynamo           DynamoConfiguration `mapstructure:"DYNAMO"`
}

type httpServerConfig struct {
	Port               string `mapstructure:"PORT"`
	IdleTimeoutSeconds string `mapstructure:"IDLE_TIMEOUT_SECONDS"`
}

type dbConfig struct {
	UserName string `mapstructure:"USER_NAME"`
	Password string `mapstructure:"PASSWORD"`
	URL      string `mapstructure:"URL"`
}

type DynamoConfiguration struct {
	DaxEndpoint      string `mapstructure:"DAX_END_POINT"`
	DaxRegion        string `mapstructure:"DAX_REGION"`
	DBEndpoint       string `mapstructure:"DB_END_POINT"`
	DBRegion         string `mapstructure:"DB_REGION"`
	DaxEnabled       bool   `mapstructure:"DAX_ENABLED"`
	WidgetTable      string `mapstructure:"WIDGET_TABLE"`
	WidgetAuditTable string `mapstructure:"WIDGET_AUDIT_TABLE"`
	PageLimit        int64  `mapstructure:"PAGE_LIMIT"`
}

func GetConfig() config {
	initConfig()
	return configuration
}

func initConfig() {
	configSync.Do(func() {
		loadConfig(&configuration, "default", "./deployments/config", "etc/demo")
	})
}

func loadConfig(config *config, configName string, paths ...string) {
	viper.SetConfigName(configName)
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
