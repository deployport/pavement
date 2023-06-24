package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Root is the root config
type Root[T any] struct {
	C                 *T
	defaultConfigYAML []byte
}

// NewRoot returns a new root config
func NewRoot[T any](defaultConfigYAML []byte) *Root[T] {
	r := &Root[T]{
		defaultConfigYAML: defaultConfigYAML,
	}
	c := r.DefaultCopy()
	r.C = &c

	return r
}

// DefaultCopy returns a copy of the default config
func (root *Root[T]) DefaultCopy() T {
	var config T

	// unmarshal yaml defaultBytes into config
	if err := yaml.Unmarshal(root.defaultConfigYAML, &config); err != nil {
		panic(fmt.Errorf("failed to decode default config copy, %w", err))
	}

	return config
}

// WriteFile writes the config to the given filename
func (root *Root[T]) WriteFile(filename string) error {
	buf, err := yaml.Marshal(root.C)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filename, buf, 0644); err != nil {
		return err
	}
	return nil
}

// OverrideFromFile overrides the config from a file
func (root *Root[T]) OverrideFromFile(filename string) error {
	if err := root.overrideFromFile(filename); err != nil {
		return fmt.Errorf("failed to override config from file %s, %w", filename, err)
	}
	return nil
}

func (root *Root[T]) overrideFromFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(content, root); err != nil {
		return err
	}
	return nil
}
