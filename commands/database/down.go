package maindb

import (
	"context"

	"github.com/spf13/cobra"
	"go.deployport.com/pavement/logging"
	sqlconfig "go.deployport.com/pavement/sql/config"
	"go.deployport.com/pavement/sql/migrations"
)

func buildDownCommand(
	ctx context.Context,
	dbconfig *sqlconfig.Connection,
	logger *logging.Logger,
	newCatalog migrations.PreparedBuilder,
) *cobra.Command {
	confirm := false
	var cmd = &cobra.Command{
		Use:   "down",
		Short: "Rollbacks all pending down migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			catalog, err := newCatalog(logger.Logger, *dbconfig)
			if err != nil {
				return err
			}
			return catalog.Down(ctx)
		},
	}
	cmd.Flags().BoolVar(&confirm, "confirm", false, "confirms the caller understand this command rollbacks all migrations effectively destroying the structure of the database")
	cmd.MarkFlagRequired("confirm")
	return cmd
}
