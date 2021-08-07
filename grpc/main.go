package main

import (
	"grpc/finder-protocol-buffers/pb"
	"grpc/infrastructure"
	"grpc/infrastructure/logger"
	"grpc/interface/controller"
	"grpc/repository"
	"grpc/usecase"
)

func main() {
	// setup db & logger
	db := infrastructure.NewGormConnect()
	logger.NewLogger(db)
	defer db.Close()

	// initialize repository
	userRepository := repository.NewUserRepository(db)
	footPrintRepository := repository.NewFootPrintRepository(db)
	likeRepository := repository.NewLikeRepository(db)
	roomRepository := repository.NewRoomRepository(db)
	roomUserRepository := repository.NewRoomUserRepository(db)
	messageRepository := repository.NewMessageRepository(db)

	// initialize usecase
	userUsecase := usecase.NewUserUsecase(
		userRepository,
		footPrintRepository,
		likeRepository,
	)
	footPrintUsecase := usecase.NewFootPrintUsecase(
		footPrintRepository,
	)
	profileUsecase := usecase.NewProfileUsecase(
		userRepository,
	)
	likeUsecase := usecase.NewLikeUsecase(
		likeRepository,
		roomRepository,
		roomUserRepository,
	)
	roomUsecase := usecase.NewRoomUsecase(
		roomRepository,
	)
	messageUsecase := usecase.NewMessageUsecase(
		messageRepository,
		roomUserRepository,
	)

	// initiliaze controller
	userController := controller.NewUserController(userUsecase)
	footPrintController := controller.NewFootPrintController(footPrintUsecase)
	profileController := controller.NewProfileController(profileUsecase)
	likeController := controller.NewLikeController(likeUsecase)
	roomController := controller.NewRoomController(roomUsecase)
	messageController := controller.NewMessageController(messageUsecase)

	// register grpc server
	grpcServer := infrastructure.NewGrpcServer()
	pb.RegisterUserServiceServer(grpcServer, userController)
	pb.RegisterFootPrintServiceServer(grpcServer, footPrintController)
	pb.RegisterProfileServiceServer(grpcServer, profileController)
	pb.RegisterLikeServiceServer(grpcServer, likeController)
	pb.RegisterRoomServiceServer(grpcServer, roomController)
	pb.RegisterMessageServiceServer(grpcServer, messageController)

	infrastructure.GrpcServe(grpcServer)
}
