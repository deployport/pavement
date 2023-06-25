package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Backed is a wrapper around a configuration struct that is backed by a files that loaded
// in order, override each value in the config struct
type Backed[T any] struct {
	C                 *T
	defaultConfigYAML []byte
}

// NewBacked returns a new backed config
func NewBacked[T any](defaultConfigYAML []byte) *Backed[T] {
	r := &Backed[T]{
		defaultConfigYAML: defaultConfigYAML,
	}
	c := r.DefaultCopy()
	r.C = &c

	return r
}

// DefaultCopy returns a copy of the default config
func (backed *Backed[T]) DefaultCopy() T {
	var config T

	// unmarshal yaml defaultBytes into config
	if err := yaml.Unmarshal(backed.defaultConfigYAML, &config); err != nil {
		panic(fmt.Errorf("failed to decode default config copy, %w", err))
	}

	return config
}

// WriteFile writes the config to the given filename
func (backed *Backed[T]) WriteFile(filename string) error {
	buf, err := yaml.Marshal(backed.C)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filename, buf, 0644); err != nil {
		return err
	}
	return nil
}

// OverrideFromFile overrides the config from a file
func (backed *Backed[T]) OverrideFromFile(filename string) error {
	if err := backed.overrideFromFile(filename); err != nil {
		return fmt.Errorf("failed to override config from file %s, %w", filename, err)
	}
	return nil
}

func (backed *Backed[T]) overrideFromFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(content, backed.C); err != nil {
		return err
	}
	return nil
}
