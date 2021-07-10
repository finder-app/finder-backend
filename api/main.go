package main

import (
	"finder/infrastructure"
	"finder/infrastructure/logger"
	"finder/infrastructure/repository"
	"finder/interface/controller"
	"finder/pb"
	"finder/usecase"
	"os"

	"google.golang.org/grpc"
)

func main() {
	db := infrastructure.NewGormConnect()
	logger.NewLogger(db)
	router := infrastructure.NewRouter()

	likeRepository := repository.NewLikeRepository(db)
	roomRepository := repository.NewRoomRepository(db)
	roomUserRepository := repository.NewRoomUserRepository(db)

	likeUsecase := usecase.NewLikeUsecase(
		likeRepository,
		roomRepository,
		roomUserRepository,
	)

	target := os.Getenv("GRPC_SERVER_NAME") + ":" + os.Getenv("GRPC_SERVER_PORT")
	grpcClientConn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer grpcClientConn.Close()

	userClient := pb.NewUserServiceClient(grpcClientConn)
	footPrintClient := pb.NewFootPrintServiceClient(grpcClientConn)
	profileClint := pb.NewProfileServiceClient(grpcClientConn)

	userController := controller.NewUserController(userClient)
	footPrintController := controller.NewFootPrintController(footPrintClient)
	likeController := controller.NewLikeController(likeUsecase)
	profileController := controller.NewProfileController(profileClint)

	router.Users(userController)
	router.Profile(profileController)
	router.FootPrints(footPrintController)
	router.Likes(likeController)

	// NOTE: GraphQLの導入は保留中なのでコメントアウト
	// resolver := graph.NewResolver(
	// 	userUsecase,
	// )
	// server := graph.NewGraphQLHandler(resolver)
	// playGroundHandler := graph.NewPlayGroundHandler()
	// router.GraphQL(server, playGroundHandler)

	router.Engine.Run(":" + os.Getenv("PORT"))
}
