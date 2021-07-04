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

	userRepository := repository.NewUserRepository(db)
	likeRepository := repository.NewLikeRepository(db)
	roomRepository := repository.NewRoomRepository(db)
	roomUserRepository := repository.NewRoomUserRepository(db)

	likeUsecase := usecase.NewLikeUsecase(
		likeRepository,
		roomRepository,
		roomUserRepository,
	)
	profileUsecase := usecase.NewProfileUsecase(userRepository)

	target := os.Getenv("GRPC_SERVER_NAME") + ":" + os.Getenv("GRPC_SERVER_PORT")
	grpcClientConn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer grpcClientConn.Close()

	userClient := pb.NewUserServiceClient(grpcClientConn)
	footPrintClient := pb.NewFootPrintServiceClient(grpcClientConn)

	userController := controller.NewUserController(userClient)
	footPrintController := controller.NewFootPrintController(footPrintClient)
	likeController := controller.NewLikeController(likeUsecase)
	profileController := controller.NewProfileController(profileUsecase)

	router.Users(userController)
	router.Profile(profileController)

	router.Engine.GET("/foot_prints", footPrintController.Index)
	router.Engine.GET("/foot_prints/unread_count", footPrintController.UnreadCount)
	router.Engine.POST("/users/:uid/likes", likeController.Create)
	router.Engine.GET("/likes", likeController.Index)
	router.Engine.PUT("/likes/:sent_uesr_uid/consent", likeController.Consent)
	router.Engine.PUT("/likes/:sent_uesr_uid/next", likeController.Next)
	// router.Engine.GET("/likes/recieved", func(c *gin.Context) { likeController.Recieved(c) })
	// router.Engine.GET("/likes/sent", func(c *gin.Context) { likeController.Sent(c) })

	// resolver := graph.NewResolver(
	// 	userUsecase,
	// )
	// server := graph.NewGraphQLHandler(resolver)
	// playGroundHandler := graph.NewPlayGroundHandler()
	// router.GraphQL(server, playGroundHandler)

	router.Engine.Run(":" + os.Getenv("PORT"))
}
