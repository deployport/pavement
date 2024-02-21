package config

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.deployport.com/pavement/config"
)

func buildInitCommand[T any](ctx context.Context, config *config.Backed[T], configFileName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "initializes a copy of a default configuration file",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := "mainprocess-config.yaml"

			if err := config.WriteFile(configFileName); err != nil {
				return err
			}
			fmt.Fprintf(os.Stderr, "initialized config file %s\n", configFileName)
			return nil
		},
	}
	return cmd
}
