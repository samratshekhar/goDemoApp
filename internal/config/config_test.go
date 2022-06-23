package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_loadConfig(t *testing.T) {
	t.Run("valid config file", func(t *testing.T) {
		config := config{}
		loadConfig(&config, "valid_env", ".")
		assert.Equal(t, "debug-test", config.Loglevel)
		assert.Equal(t, "dev-test", config.Environment)
		assert.Equal(t, "9999", config.HTTPServerConfig.IdleTimeoutSeconds)
		assert.Equal(t, "9999", config.HTTPServerConfig.Port)
	})

	t.Run("invalid config file path", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		config := config{}
		loadConfig(&config, "invalid_name", "/invalid_path")
	})

	t.Run("invalid config file values", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		config := config{}
		loadConfig(&config, "invalid_env", ".")
	})
}