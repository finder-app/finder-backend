package main

import (
	"finder/infrastructure"
	"finder/infrastructure/logger"
	finderRouter "finder/infrastructure/router"
	"finder/interface/controller"
	"finder/interface/repository"
	"finder/usecase/interactor"

	"github.com/gin-gonic/gin"
)

func main() {
	db := infrastructure.NewGormConnect()
	logger.NewLogger(db)
	validate := infrastructure.NewValidator()
	router := finderRouter.NewRouter()

	footPrintRepository := repository.NewFootPrintRepository(db, validate)
	footPrintInteractor := interactor.NewFootPrintInteractor(footPrintRepository)
	footPrintController := controller.NewFootPrintController(footPrintInteractor)

	userRepository := repository.NewUserRepository(db, validate)
	userInteractor := interactor.NewUserInteractor(userRepository, footPrintRepository)
	userController := controller.NewUserController(userInteractor)

	likeRepository := repository.NewLikeRepository(db, validate)
	likeInteractor := interactor.NewLikeInteractor(likeRepository)
	likeController := controller.NewLikeController(likeInteractor)

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
