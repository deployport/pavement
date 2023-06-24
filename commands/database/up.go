package maindb

import (
	"context"

	"github.com/deployport/pavement/logging"
	sqlconfig "github.com/deployport/pavement/sql/config"
	"github.com/deployport/pavement/sql/migrations"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
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
			logger.Logger.Warn("db", zap.String("url", dbconfig.URL))
			catalog, err := newCatalog(logger.Logger, *dbconfig)
			if err != nil {
				return err
			}
			return catalog.Up(ctx)
		},
	}
	return cmd
}
