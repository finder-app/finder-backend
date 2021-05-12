package main

import (
	"finder/infrastructure"
	"finder/infrastructure/logger"
	finderRouter "finder/infrastructure/router"
	"finder/interface/controller"
	"finder/interface/repository"
	"finder/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	db := infrastructure.NewGormConnect()
	logger.NewLogger(db)
	validate := infrastructure.NewValidator()
	router := finderRouter.NewRouter()

	footPrintRepository := repository.NewFootPrintRepository(db, validate)
	footPrintUsecase := usecase.NewFootPrintUsecase(footPrintRepository)
	footPrintController := controller.NewFootPrintController(footPrintUsecase)

	userRepository := repository.NewUserRepository(db, validate)
	userUsecase := usecase.NewUserUsecase(userRepository, footPrintRepository)
	userController := controller.NewUserController(userUsecase)

	likeRepository := repository.NewLikeRepository(db, validate)
	likeUsecase := usecase.NewLikeUsecase(likeRepository)
	likeController := controller.NewLikeController(likeUsecase)

	router.Engine.GET("/foot_prints", func(c *gin.Context) { footPrintController.Index(c) })
	router.Users(userController)
	router.Engine.POST("/users/:uid/likes", func(c *gin.Context) { likeController.Create(c) })
	router.Engine.GET("/likes", func(c *gin.Context) { likeController.Index(c) })
	router.Engine.PUT("/likes/:sent_uesr_uid", func(c *gin.Context) { likeController.Update(c) })
	router.Engine.PUT("/likes/:sent_uesr_uid/next", func(c *gin.Context) { likeController.Next(c) })
	// router.Engine.GET("/likes/recieved", func(c *gin.Context) { likeController.Recieved(c) })
	// router.Engine.GET("/likes/sent", func(c *gin.Context) { likeController.Sent(c) })

	router.Engine.Run(":8080")
}
