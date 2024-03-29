package migrations

import (
	"context"
	"errors"
	"fmt"
	"io/fs"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"go.deployport.com/pavement/sql/config"
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

// PreparedBuilder is a function that creates a Catalog instance
type PreparedBuilder func(logger *zap.Logger, config config.Connection) (*Catalog, error)

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

// Down rollbacks the all previous database migrations
func (m *Catalog) Down(ctx context.Context) error {
	logger := m.logger
	logger.Debug("down")
	return m.migrate.Down()
}

// Rollback rollbacks the last database migration
func (m *Catalog) Rollback(ctx context.Context) error {
	logger := m.logger
	logger.Debug("down")
	if err := m.migrate.Steps(-1); err != nil {
		return fmt.Errorf("failed to rollback database migration, %w", err)
	}
	return nil
}
