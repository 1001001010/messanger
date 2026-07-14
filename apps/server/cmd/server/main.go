package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/1001001010/messanger/internal/config"
	"github.com/1001001010/messanger/internal/database"
	"github.com/1001001010/messanger/internal/logger"
)

func main() {
	// Подгружаем конфиг
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфига: %v", err)
	}
	fmt.Printf("Конфиг загружен: env=%s grpc_port=%d\n", cfg.Env, cfg.GRPC.Port)

	// Подгружаем логгер
	loggerInstance := logger.SetupLogger(cfg.Env)
	loggerInstance.Info("Логирование инициализировано")

	// Коннектим БД
	dbPool, err := database.Connect(context.Background(), cfg, loggerInstance)
	if err != nil {
		loggerInstance.Error("Ошибка подключения к бд", "error", err)
		log.Fatalf("Ошибка подключения к бд: %v", err)
	}
	defer dbPool.Close()

	loggerInstance.Info("Сервер запущен", "grpc_port", cfg.GRPC.Port)

	// Создаем gRPC сервер
	grpcServer := grpc.NewServer()

	// Регистрируем health check сервис
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	loggerInstance.Info("gRPC сервер запущен", "grpc_port", cfg.GRPC.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		loggerInstance.Error("failed to listen", "error", err)
		log.Fatalf("failed to listen: %v", err)
	}

	// Запускаем сервер в отдельной goroutine
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			loggerInstance.Error("grpc server error", "error", err)
		}
	}()

	loggerInstance.Info("gRPC server started", "port", cfg.GRPC.Port)

	// Ловим сигналы завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Блокируемся до получения сигнала
	sig := <-sigChan
	loggerInstance.Info("shutdown signal received", "signal", sig.String())

	// Безопасно останавливаем gRPC сервер
	grpcServer.GracefulStop()
	loggerInstance.Info("gRPC server stopped")
}
