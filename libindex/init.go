package libindex

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/remind101/migrate"

	"github.com/Panzer1119/claircore/internal/indexer"
	"github.com/Panzer1119/claircore/internal/indexer/postgres"
	"github.com/Panzer1119/claircore/libindex/migrations"
)

// initialize a indexer.Store given libindex.Opts
func initStore(ctx context.Context, opts *Opts) (indexer.Store, error) {
	// we are going to use pgx for more control over connection pool and
	// and a cleaner api around bulk inserts
	cfg, err := pgxpool.ParseConfig(opts.ConnString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ConnString: %v", err)
	}
	cfg.MaxConns = 30
	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create ConnPool: %v", err)
	}

	db, err := sql.Open("pgx", opts.ConnString)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %v", err)
	}
	defer db.Close()

	// do migrations if requested
	if opts.Migrations {
		migrator := migrate.NewPostgresMigrator(db)
		migrator.Table = migrations.MigrationTable
		err := migrator.Exec(migrate.Up, migrations.Migrations...)
		if err != nil {
			return nil, fmt.Errorf("failed to perform migrations: %w", err)
		}
	}

	store := postgres.NewStore(pool)
	return store, nil
}
