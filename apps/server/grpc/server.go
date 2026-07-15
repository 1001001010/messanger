package grpc

import (
	"fmt"
	"log"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/1001001010/messanger/internal/config"
)

// Создаем gRPC сервер
func StartGRPCServer(cfg *config.Config, loggerInstance *slog.Logger) *grpc.Server {
	grpcServer := grpc.NewServer()

	// Регистрируем health check сервис
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	loggerInstance.Info("gRPC server initialized", "grpc_port", cfg.GRPC.Port)

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

	return grpcServer
}
