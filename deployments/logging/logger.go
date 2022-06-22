package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goDemoApp/internal/config"
	"sync"
)

var logger *zap.SugaredLogger
var loggingSync sync.Once

func GetLogger() *zap.SugaredLogger {
	initLogger()
	return logger
}

func initLogger() {
	loggingSync.Do(func() {
		var logLevel zapcore.Level
		err := logLevel.UnmarshalText([]byte(config.GetConfig().Loglevel))
		if err == nil {
			panic("Error setting up logger " + err.Error())
		}
		configObj := zap.Config{
			Level:            zap.NewAtomicLevelAt(logLevel),
			Development:      false,
			Encoding:         "json",
			EncoderConfig:    zap.NewProductionEncoderConfig(),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stdout"},
		}
		configObj.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		log, err := configObj.Build()
		if err != nil {
			panic("Error setting up logger " + err.Error())
		}
		logger = log.Sugar()
	})
}
