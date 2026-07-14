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
		return nil, fmt.Errorf("Отсутствует URL базы данных в конфигурации")
	}
	pool, err := pgxpool.New(ctx, cfg.DBURL)
	if err != nil {
		log.Error("Ошибка создания пула подключения", "error", err)
		return nil, fmt.Errorf("Ошибка создания пула подключения: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		log.Error("Ошибка подключения к базе данных", "error", err)
		return nil, fmt.Errorf("Ошибка подключения к базе данных: %w", err)
	}

	log.Info("БД успешно подключена")
	return pool, nil
}
