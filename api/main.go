package main

import (
	"api/infrastructure"
	"api/infrastructure/logger"
	"api/interface/controller"
	"api/pb"
	"context"
	"log"
)

func main() {
	logger.NewLogger()

	// initialize grpc client
	grpcClientConn := infrastructure.GrpcClientConn()
	defer grpcClientConn.Close()
	userClient := pb.NewUserServiceClient(grpcClientConn)
	footPrintClient := pb.NewFootPrintServiceClient(grpcClientConn)
	profileClient := pb.NewProfileServiceClient(grpcClientConn)
	likeClient := pb.NewLikeServiceClient(grpcClientConn)
	roomClient := pb.NewRoomServiceClient(grpcClientConn)
	messageClient := pb.NewMessageServiceClient(grpcClientConn)

	// initialize controller
	userController := controller.NewUserController(userClient)
	footPrintController := controller.NewFootPrintController(footPrintClient)
	likeController := controller.NewLikeController(likeClient)
	stream, err := profileClient.TestImage(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	profileController := controller.NewProfileController(profileClient, stream)
	roomController := controller.NewRoomController(roomClient)
	messageController := controller.NewMessageController(messageClient)

	// set router
	router := infrastructure.NewRouter()
	router.Users(userController)
	router.Profile(profileController)
	router.FootPrints(footPrintController)
	router.Likes(likeController)
	router.Rooms(roomController)
	router.Messages(messageController)

	// NOTE: GraphQLの導入は保留中なのでコメントアウト
	// resolver := graph.NewResolver(
	// 	userUsecase,
	// )
	// server := graph.NewGraphQLHandler(resolver)
	// playGroundHandler := graph.NewPlayGroundHandler()
	// router.GraphQL(server, playGroundHandler)

	router.Run()
}
