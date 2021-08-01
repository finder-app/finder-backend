package main

import (
	"api/infrastructure"
	"api/infrastructure/logger"
	"api/interface/controller"
	"api/pb"
	"log"
	"os"

	"google.golang.org/grpc"
)

func main() {
	logger.NewLogger()
	router := infrastructure.NewRouter()

	target := os.Getenv("GRPC_SERVER_NAME") + ":" + os.Getenv("GRPC_SERVER_PORT")
	grpcClientConn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer grpcClientConn.Close()

	userClient := pb.NewUserServiceClient(grpcClientConn)
	footPrintClient := pb.NewFootPrintServiceClient(grpcClientConn)
	profileClient := pb.NewProfileServiceClient(grpcClientConn)
	likeClient := pb.NewLikeServiceClient(grpcClientConn)
	roomClient := pb.NewRoomServiceClient(grpcClientConn)
	messageClient := pb.NewMessageServiceClient(grpcClientConn)

	userController := controller.NewUserController(userClient)
	footPrintController := controller.NewFootPrintController(footPrintClient)
	likeController := controller.NewLikeController(likeClient)
	profileController := controller.NewProfileController(profileClient)
	roomController := controller.NewRoomController(roomClient)
	messageController := controller.NewMessageController(messageClient)

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

	log.Print("http server start")
	router.Engine.Run(":" + os.Getenv("PORT"))
}
