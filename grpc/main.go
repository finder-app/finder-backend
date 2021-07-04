package main

import (
	"grpc/infrastructure"
	"grpc/infrastructure/logger"
	"grpc/pb"
	"grpc/repository"
	"grpc/usecase"
	"log"
	"net"
	"os"
)

func main() {
	db := infrastructure.NewGormConnect()
	logger.NewLogger(db)
	grpcServer := infrastructure.NewGrpcServer()

	userRepository := repository.NewUserRepository(db)
	footPrintRepository := repository.NewFootPrintRepository(db)

	userUsecase := usecase.NewUserUseuserUsecase(userRepository, footPrintRepository)
	footPrintUsecase := usecase.NewFootPrintUsecase(footPrintRepository)
	profileUsecase := usecase.NewProfileUsecase(userRepository)

	pb.RegisterUserServiceServer(grpcServer, userUsecase)
	pb.RegisterFootPrintServiceServer(grpcServer, footPrintUsecase)
	pb.RegisterProfileServiceServer(grpcServer, profileUsecase)

	listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_SERVER_PORT")) // [::]:50051
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	log.Print("grpcServer serve")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("serve err", err)
	}
}
