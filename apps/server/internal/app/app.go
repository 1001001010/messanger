package app

import (
	"log/slog"

	"google.golang.org/grpc"
)

func ShutdownGRPC(grpcServer *grpc.Server, loggerInstance *slog.Logger) {
	// Корректно останавливаем gRPC сервер
	grpcServer.GracefulStop()
	loggerInstance.Info("gRPC server stopped")
}
