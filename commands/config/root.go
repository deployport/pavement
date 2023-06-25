package config

import (
	"context"

	"github.com/deployport/pavement/config"
	"github.com/spf13/cobra"
)

// RootParams are the parameters for the root command
type RootParams[T any] struct {
	BackedConfig         *config.Backed[T]
	InitFilename string
}

// Root creates the root command
func Root[T any](ctx context.Context, params RootParams[T]) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use: "config",
	}
	rootCmd.AddCommand(buildInitCommand(ctx, params.BackedConfig, params.InitFilename))
	return rootCmd
}
