package infrastructure

import (
	"log"

	"github.com/finder-app/finder-backend/api/infrastructure/env"

	"google.golang.org/grpc"
)

func GrpcClientConn() *grpc.ClientConn {
	target := env.GRPC_SERVER_NAME + ":" + env.GRPC_SERVER_PORT
	grpcClientConn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	return grpcClientConn
}
