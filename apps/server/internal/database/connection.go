package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/1001001010/messanger/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, cfg *config.Config, log *slog.Logger) (*pgxpool.Pool, error) {
	if cfg.DBURL == "" {
		return nil, fmt.Errorf("There is no database URL in the configuration")
	}
	pool, err := pgxpool.New(ctx, cfg.DBURL)
	if err != nil {
		log.Error("Connection pool creation error", "error", err)
		return nil, fmt.Errorf("Error creating connection pool: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		log.Error("Error connecting to the database", "error", err)
		return nil, fmt.Errorf("Error connecting to the database: %w", err)
	}

	log.Info("Database connected successfully")
	return pool, nil
}
