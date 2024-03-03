package maindb

import (
	"context"

	"github.com/spf13/cobra"
	"go.deployport.com/pavement/logging"
	sqlconfig "go.deployport.com/pavement/sql/config"
	"go.deployport.com/pavement/sql/migrations"
)

func buildRollbackCommand(
	ctx context.Context,
	dbconfig *sqlconfig.Connection,
	logger *logging.Logger,
	newCatalog migrations.PreparedBuilder,
) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "rollback",
		Short: "Rollbacks the last migration",
		RunE: func(cmd *cobra.Command, args []string) error {
			catalog, err := newCatalog(logger.Logger, *dbconfig)
			if err != nil {
				return err
			}
			return catalog.Rollback(ctx)
		},
	}
	return cmd
}
