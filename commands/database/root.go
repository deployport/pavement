package maindb

import (
	"context"

	"github.com/spf13/cobra"
	"go.deployport.com/pavement/logging"
	sqlconfig "go.deployport.com/pavement/sql/config"
	"go.deployport.com/pavement/sql/migrations"
)

// RootParams are the parameters for the root command
type RootParams struct {
	NewCatalog migrations.PreparedBuilder
	Connection *sqlconfig.Connection
	Logger     *logging.Logger
}

// Build creates the root command
func Build(
	ctx context.Context,
	params RootParams,
) []*cobra.Command {
	return []*cobra.Command{
		buildUpCommand(
			ctx,
			params.Connection,
			params.Logger,
			params.NewCatalog,
		),
		buildDownCommand(
			ctx,
			params.Connection,
			params.Logger,
			params.NewCatalog,
		),
		buildRollbackCommand(
			ctx,
			params.Connection,
			params.Logger,
			params.NewCatalog,
		),
	}
}
