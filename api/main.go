package main

import (
	"github.com/finder-app/finder-backend/api/finder-protocol-buffers/pb"
	"github.com/finder-app/finder-backend/api/infrastructure"
	"github.com/finder-app/finder-backend/api/infrastructure/aws"
	"github.com/finder-app/finder-backend/api/infrastructure/logger"
	"github.com/finder-app/finder-backend/api/interface/controller"
)

func main() {
	// initialize etc...
	logger.NewLogger()
	s3uploader := aws.NewS3uplodaer()

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
	profileController := controller.NewProfileController(profileClient, s3uploader)
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
