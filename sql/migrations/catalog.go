package migrations

import (
	"context"
	"errors"
	"fmt"
	"io/fs"

	"github.com/deployport/pavement/sql/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"go.uber.org/zap"
)

// Catalog contains the database migrations
type Catalog struct {
	migrate *migrate.Migrate
	logger  *zap.Logger
}

// NewCatalog creates a new Catalog instance
func NewCatalog(
	logger *zap.Logger,
	config config.Connection,
	fs fs.FS,
) (*Catalog, error) {
	logger = logger.Named("migration-catalog")
	d, err := iofs.New(fs, "sql")
	if err != nil {
		return nil, fmt.Errorf("failed to load database migrations, %w", err)
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, config.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrations source instance, %w", err)
	}
	return &Catalog{
		migrate: m,
		logger:  logger,
	}, nil
}

// Up runs all pending migrations
func (m *Catalog) Up(ctx context.Context) error {
	logger := m.logger
	logger.Debug("up")
	err := m.migrate.Up()
	v, _, _ := m.migrate.Version()
	if errors.Is(err, migrate.ErrNoChange) {
		logger.Info("up-to-date", zap.Uint("version", v))
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to run database migrations, %w", err)
	}
	logger.Info("migrated", zap.Uint("version", v))
	return nil
}

// Down rollbacks the last database migration
func (m *Catalog) Down(ctx context.Context) error {
	logger := m.logger
	logger.Debug("down")
	return m.migrate.Down()
}
