package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	loggingconfig "go.deployport.com/pavement/logging/config"
)

// Logger is a wrapper around zap.Logger
// that can be configured given certan user preferences
type Logger struct {
	*zap.Logger
}

// NewLogger returns a new logger
func NewLogger() *Logger {
	return &Logger{}
}

// Configure configures the logger
func (logger *Logger) Configure(config loggingconfig.Config) {
	level := zap.NewAtomicLevelAt(config.Level)
	var encoderConfig zapcore.EncoderConfig
	if config.Development {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	stdout := zapcore.AddSync(os.Stdout)

	zl := zap.New(zapcore.NewCore(encoder, stdout, level))
	logger.Logger = zl
}
