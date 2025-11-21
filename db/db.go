// Package repo provides a repository for transfers, ie a way to store and retrieve transfers.
package db

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // File source for migrations
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"

	"github.com/kimbohlovette/clando-backend/db/sqlc"
)

type AppStore interface {
	Begin(ctx context.Context) (sqlc.Querier, pgx.Tx, error)
	Do() sqlc.Querier
}

type Impl struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Impl {
	return &Impl{db: db}
}

func (u *Impl) Begin(ctx context.Context) (sqlc.Querier, pgx.Tx, error) {
	tx, err := u.db.Begin(ctx)
	if err != nil {
		return nil, nil, err
	}
	return sqlc.New(tx), tx, nil
}

func (u *Impl) Do() sqlc.Querier {
	return sqlc.New(u.db)
}

// Migrate function applies migrations to the database.
func Migrate(dbURL string, migrationsPath string, logger zerolog.Logger) error {
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return err
	}

	// Create a new migration instance with the absolute path
	m, err := migrate.New(
		"file://"+absPath,
		dbURL,
	)
	if err != nil {
		return err
	}
	defer m.Close()

	// Apply migrations
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	logger.Info().Msg("Migrations applied successfully")
	return nil
}

// MigrateDown function rolls back migrations from the database.
func MigrateDown(dbURL string, migrationsPath string, logger zerolog.Logger) error {
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return err
	}

	// Create a new migration instance with the absolute path
	m, err := migrate.New(
		"file://"+absPath,
		dbURL,
	)
	if err != nil {
		return err
	}
	defer m.Close()

	// Apply migrations
	if err = m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	logger.Info().Msg("Migrations applied successfully")
	return nil
}
