package config

import (
	"testing"

	"github.com/jackc/pgx/v5/tracelog"
)

func TestLogLevel(t *testing.T) {
	var logLevel LogLevel
	logLevel.UnmarshalText([]byte("debug"))
	if logLevel.TraceLog().String() != "debug" {
		t.Errorf("expected debug, got %s", logLevel.TraceLog().String())
	}
	logLevel = LogLevel(tracelog.LogLevelInfo)
	// test MarshalText
	text, err := logLevel.MarshalText()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if string(text) != "info" {
		t.Errorf("expected debug, got %s", string(text))
	}
}
