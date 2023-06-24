package config

import (
	"context"
	"fmt"
	"os"

	"github.com/deployport/pavement/config"
	"github.com/spf13/cobra"
)

func buildInitCommand[T any](ctx context.Context, root *config.Root[T], configFileName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "initializes a copy of a default configuration file",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := "mainprocess-config.yaml"

			if err := root.WriteFile(configFileName); err != nil {
				return err
			}
			fmt.Fprintf(os.Stderr, "initialized config file %s\n", configFileName)
			return nil
		},
	}
	return cmd
}
