package maindb

import (
	"context"

	"github.com/deployport/pavement/logging"
	sqlconfig "github.com/deployport/pavement/sql/config"
	"github.com/deployport/pavement/sql/migrations"
	"github.com/spf13/cobra"
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
	}
}
