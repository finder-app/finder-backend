package infrastructure

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer() *grpc.Server {
	zapLogger, _ := zap.NewDevelopment()

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(zapLogger),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(zapLogger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	// NOTE: grpcurl用 development環境でのみ実行する
	reflection.Register(grpcServer)

	return grpcServer
}
