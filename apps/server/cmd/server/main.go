package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/1001001010/messanger/grpc"
	"github.com/1001001010/messanger/internal/app"
	"github.com/1001001010/messanger/internal/config"
	"github.com/1001001010/messanger/internal/database"
	"github.com/1001001010/messanger/internal/logger"
)

func main() {
	// Подгружаем конфиг
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	fmt.Printf("Config loaded: env=%s grpc_port=%d\n", cfg.Env, cfg.GRPC.Port)

	// Подгружаем логгер
	loggerInstance := logger.SetupLogger(cfg.Env)
	loggerInstance.Info("Logging initialized")

	// Коннектим БД
	dbPool, err := database.Connect(context.Background(), cfg, loggerInstance)
	if err != nil {
		loggerInstance.Error("Error connecting to the db", "error", err)
		log.Fatalf("Error connecting to the db: %v", err)
	}

	// Запускаем gRPC сервер
	grpcServer := grpc.StartGRPCServer(cfg, loggerInstance)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Ждем сигнала завершения
	<-ctx.Done()

	loggerInstance.Info("Shutdown signal received")

	// Останавливаем gRPC сервер
	app.ShutdownGRPC(grpcServer, loggerInstance)
	defer dbPool.Close()
	loggerInstance.Info("Database connection closed")
}
