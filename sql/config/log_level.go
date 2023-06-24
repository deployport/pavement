package config

import (
	"errors"

	"github.com/jackc/pgx/v5/tracelog"
)

// LogLevel is the pgx logging level. See LogLevel* constants for
type LogLevel tracelog.LogLevel

// MarshalText marshals the Level to text.
func (l *LogLevel) MarshalText() ([]byte, error) {
	if l == nil {
		return nil, errUnmarshalNilLevel
	}
	return []byte(tracelog.LogLevel(*l).String()), nil
}

// UnmarshalText unmarshals text to a level. Like MarshalText, UnmarshalText
// expects the text representation of a Level to drop the -Level suffix (see
// example).
//
// In particular, this makes it easy to configure logging levels using YAML,
// TOML, or JSON files.
func (l *LogLevel) UnmarshalText(text []byte) error {
	if l == nil {
		return errUnmarshalNilLevel
	}
	lv, err := tracelog.LogLevelFromString(string(text))
	if err != nil {
		return err
	}
	*l = LogLevel(lv)
	return nil
}

// TraceLog returns the tracelog.LogLevel equivalent of the Level
func (l *LogLevel) TraceLog() tracelog.LogLevel {
	if l == nil {
		return tracelog.LogLevelNone
	}
	return tracelog.LogLevel(*l)
}

var errUnmarshalNilLevel = errors.New("can't unmarshal a nil *Level")
