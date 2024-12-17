package postgres

import (
	"context"
	"fmt"

	"github.com/404th/value-holder/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgres(cfg *config.Config) (db *pgxpool.Pool, err error) {
	// postgres connection string
	postgresConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&pool_max_conns=%d",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
		cfg.PostgresSSLMode,
		cfg.PostgresMaxConnections,
	)

	// creating a pool
	var pool = new(pgxpool.Pool)
	pool, err = pgxpool.Connect(context.Background(), postgresConnectionString)
	if err != nil {
		return pool, err
	}

	// checking if connected successfully
	if err = pool.Ping(context.Background()); err != nil {
		return pool, err
	}

	return pool, err
}
