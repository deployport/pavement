package maindb

import (
	"context"

	"github.com/deployport/pavement/logging"
	sqlconfig "github.com/deployport/pavement/sql/config"
	"github.com/deployport/pavement/sql/migrations"
	"github.com/spf13/cobra"
)

func buildDownCommand(
	ctx context.Context,
	dbconfig *sqlconfig.Connection,
	logger *logging.Logger,
	newCatalog migrations.PreparedBuilder,
) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "down",
		RunE: func(cmd *cobra.Command, args []string) error {
			catalog, err := newCatalog(logger.Logger, *dbconfig)
			if err != nil {
				return err
			}
			return catalog.Down(ctx)
		},
	}
	return cmd
}
