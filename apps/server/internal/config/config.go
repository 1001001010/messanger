package config

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env      string
	DBURL    string
	LogLevel string

	GRPC GRPCConfig

	TokenTTL time.Duration
}

type GRPCConfig struct {
	Port    int
	Timeout time.Duration
}

// Чтение настроек приложения из .env и переменных окружения.
func Load() (*Config, error) {
	envPath, err := findEnvFile()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	if envPath != "" {
		if err := godotenv.Load(envPath); err != nil {
			return nil, err
		}
	}

	cfg := &Config{
		Env:      getEnv("APP_ENV", "local"),
		DBURL:    getEnv("DB_URL", ""),
		LogLevel: getEnv("LOG_LEVEL", "info"),
		TokenTTL: mustDuration(getEnv("TOKEN_TTL", "1h")),
		GRPC: GRPCConfig{
			Port:    getEnvAsInt("GRPC_PORT", 50051),
			Timeout: mustDuration(getEnv("GRPC_TIMEOUT", "5s")),
		},
	}

	return cfg, nil
}

// Поиск .env в текущей директории и родительских папках.
func findEnvFile() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		path := filepath.Join(dir, ".env")
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", os.ErrNotExist
		}
		dir = parent
	}
}

// Возвращает значение переменной окружения или fallback, если переменная не задана.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// Возвращает целочисленное значение переменной окружения или fallback.
func getEnvAsInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
	}
	return fallback
}

// Парсит строку вида "1h", "5s" и возвращает time.Duration.
func mustDuration(value string) time.Duration {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return 0
	}
	return duration
}
