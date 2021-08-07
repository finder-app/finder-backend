package infrastructure

import (
	"log"
	"os"

	"google.golang.org/grpc"
)

func GrpcClientConn() *grpc.ClientConn {
	target := os.Getenv("GRPC_SERVER_NAME") + ":" + os.Getenv("GRPC_SERVER_PORT")
	grpcClientConn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	return grpcClientConn
}
