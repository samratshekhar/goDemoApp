package logger

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"reflect"
	"testing"
)

func TestNewLogger(t *testing.T) {
	t.Run("default logger is production grade", func(t *testing.T) {
		logger, err := NewLogger()
		assert.Nil(t, err)
		loggerValue := reflect.ValueOf(logger.Desugar())
		dev := loggerValue.Elem().FieldByName("development")
		assert.False(t, dev.Bool())
	})

	t.Run("non prod logger is configurable", func(t *testing.T) {
		logger, err := NewLogger(NonProdEnv())
		assert.Nil(t, err)
		loggerValue := reflect.ValueOf(logger.Desugar())
		dev := loggerValue.Elem().FieldByName("development")
		assert.True(t, dev.Bool())
	})
}

func TestNonProdEnv(t *testing.T) {
	t.Run("non prod env", func(t *testing.T) {
		o := NonProdEnv()
		ops := &options{}
		o(ops)
		assert.True(t, ops.env == nonProd)
	})
}

func TestProdEnv(t *testing.T) {
	t.Run("prod env", func(t *testing.T) {
		o := ProdEnv()
		ops := &options{}
		o(ops)
		assert.True(t, ops.env == prod)
	})
}

func TestEncoding(t *testing.T) {
	t.Run("encoding value", func(t *testing.T) {
		o := Encoding("testEncoding")
		ops := &options{}
		o(ops)
		assert.True(t, ops.encoding == "testEncoding")
	})
}

func TestLogLevel(t *testing.T) {
	tests := []struct {
		name string
		lvl string
		want options
	}{
		{
			name: "debug",
			lvl: "debug",
			want: options{loglevel: zapcore.DebugLevel},
		},
		{
			name: "DEBUG",
			lvl: "DEBUG",
			want: options{loglevel: zapcore.DebugLevel},
		},
		{
			name: "info",
			lvl: "info",
			want: options{loglevel: zapcore.InfoLevel},
		},
		{
			name: "INFO",
			lvl: "INFO",
			want: options{loglevel: zapcore.InfoLevel},
		},
		{
			name: "warn",
			lvl: "warn",
			want: options{loglevel: zapcore.WarnLevel},
		},
		{
			name: "WARN",
			lvl: "WARN",
			want: options{loglevel: zapcore.WarnLevel},
		},
		{
			name: "error",
			lvl: "error",
			want: options{loglevel: zapcore.ErrorLevel},
		},
		{
			name: "ERROR",
			lvl: "ERROR",
			want: options{loglevel: zapcore.ErrorLevel},
		},
		{
			name: "dpanic",
			lvl: "dpanic",
			want: options{loglevel: zapcore.DPanicLevel},
		},
		{
			name: "DPANIC",
			lvl: "DPANIC",
			want: options{loglevel: zapcore.DPanicLevel},
		},
		{
			name: "panic",
			lvl: "panic",
			want: options{loglevel: zapcore.PanicLevel},
		},
		{
			name: "PANIC",
			lvl: "PANIC",
			want: options{loglevel: zapcore.PanicLevel},
		},
		{
			name: "fatal",
			lvl: "fatal",
			want: options{loglevel: zapcore.FatalLevel},
		},
		{
			name: "FATAL",
			lvl: "FATAL",
			want: options{loglevel: zapcore.FatalLevel},
		},
		{
			name: "garbage",
			lvl: "garbage",
			want: options{loglevel: zapcore.InfoLevel},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lvl := LogLevel(tt.lvl)
			got := &options{}
			lvl(got)
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("LogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}