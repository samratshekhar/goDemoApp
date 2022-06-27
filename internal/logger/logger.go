package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

const (
	nonProd logEnv = iota
	prod    logEnv = iota
)

var logger *zap.SugaredLogger
var loggingSync sync.Once

type logEnv int

type options struct {
	env      logEnv //local or prod
	encoding string
	loglevel zapcore.Level
}

type Option func(*options)

// NonProdEnv Option which would set log for local configurations
func NonProdEnv() Option {
	return func(i *options) {
		i.env = nonProd
	}
}

// ProdEnv Option to set for prod - this is default
func ProdEnv() Option {
	return func(i *options) {
		i.env = prod
	}
}

// Encoding Option to set encoding - console is default
//Other encoding supported is json
func Encoding(encoding string) Option {
	return func(i *options) {
		i.encoding = encoding
	}
}

// LogLevel Option to move loglevel to given level
//default is info
func LogLevel(lvl string) Option {
	return func(i *options) {
		var level zapcore.Level
		err := level.UnmarshalText([]byte(lvl))
		if err != nil {
			i.loglevel = zapcore.InfoLevel
		} else {
			i.loglevel = level
		}
	}
}

func GetLogger(setters ...Option) *zap.SugaredLogger {
	loggingSync.Do(func() {
		var err error
		logger, err = NewLogger(setters...)
		if err != nil {
			panic("Error initialising logger " + err.Error())
		}
	})
	return logger
}

//NewLogger by default provides production grade configuration.
// For verbosity and local usage use Option to configure.
func NewLogger(setters ...Option) (*zap.SugaredLogger, error) {
	//default values - prod and console
	ops := &options{env: prod, encoding: "json", loglevel: zapcore.InfoLevel}
	for _, setter := range setters {
		setter(ops)
	}
	var cfg zap.Config
	if ops.env == nonProd {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.Encoding = ops.encoding
	cfg.Level.SetLevel(ops.loglevel)
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(logger)
	return logger.Sugar(), nil
}
