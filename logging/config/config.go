package config

import "go.uber.org/zap/zapcore"

// Config is the configuration for the program
type Config struct {
	Level       zapcore.Level `yaml:"level"`
	Development bool          `yaml:"development"`
}
