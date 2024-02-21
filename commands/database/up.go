package maindb

import (
	"context"

	"github.com/spf13/cobra"
	"go.deployport.com/pavement/logging"
	sqlconfig "go.deployport.com/pavement/sql/config"
	"go.deployport.com/pavement/sql/migrations"
)

func buildUpCommand(
	ctx context.Context,
	dbconfig *sqlconfig.Connection,
	logger *logging.Logger,
	newCatalog migrations.PreparedBuilder,
) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "up",
		RunE: func(cmd *cobra.Command, args []string) error {
			catalog, err := newCatalog(logger.Logger, *dbconfig)
			if err != nil {
				return err
			}
			return catalog.Up(ctx)
		},
	}
	return cmd
}
