package sql

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	pgxzap "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/jackc/pgx/v5/stdlib" // register pgx driver
	"github.com/jackc/pgx/v5/tracelog"
	sqlconfig "go.deployport.com/pavement/sql/config"
	"go.deployport.com/pavement/sql/migrations"
	"go.uber.org/zap"
)

// NewClient opens a connection to the database.
// This instance can be shared safely across requests
// as connection pooling is handled by the driver.
func NewClient[TTx EntTransaction, TClient EntClient[TTx]](
	ctx context.Context,
	logger *zap.Logger,
	config sqlconfig.Connection,
	entCreator func(driver *entsql.Driver) TClient,
	newMigration migrations.PreparedBuilder,
) (*Client[TTx, TClient], error) {
	logger = logger.Named("maindb")
	databaseURL := config.URL
	logger.Info("auto migrate", zap.Bool("enabled", config.AutoMigrate))
	if config.AutoMigrate {
		catalog, err := newMigration(logger, config)
		if err != nil {
			return nil, fmt.Errorf("failed to create migration catalog, %w", err)
		}
		if err := catalog.Up(ctx); err != nil {
			return nil, fmt.Errorf("failed to run auto-migrations, %w", err)
		}
	}
	connConfig, err := pgx.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database url, %w", err)
	}
	dbLogger := pgxzap.NewLogger(logger)
	connConfig.Tracer = &tracelog.TraceLog{
		Logger:   dbLogger,
		LogLevel: config.Logging.Level.TraceLog(),
	}
	connStr := stdlib.RegisterConnConfig(connConfig)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database, %w", err)
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect database, %w", err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return &Client[TTx, TClient]{
		Client: entCreator(drv),
	}, nil
}

// EntTransaction is an interface for ent.Tx
type EntTransaction interface {
	Rollback() error
	Commit() error
}

// EntClient is an interface for ent.Client
type EntClient[TTx EntTransaction] interface {
	Tx(ctx context.Context) (TTx, error)
}

// Client entdata is a connected client instance
type Client[TTx EntTransaction, TClient EntClient[TTx]] struct {
	Client TClient
}

// WithTx executes a function in the context of a transaction
func (client *Client[TTx, TClient]) WithTx(ctx context.Context, fn func(tx TTx) error) error {
	tx, err := client.Client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
