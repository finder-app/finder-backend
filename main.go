package main

import (
	"finder/infrastructure"
	"finder/infrastructure/logger"
	"finder/interface/controller"
	"finder/interface/repository"
	"finder/usecase/interactor"

	"github.com/gin-gonic/gin"
)

func main() {
	db := infrastructure.NewGormConnect()
	logger.NewLogger(db)
	validate := infrastructure.NewValidator()
	router := infrastructure.NewRouter()

	footPrintRepository := repository.NewFootPrintRepository(db, validate)
	footPrintInteractor := interactor.NewFootPrintInteractor(footPrintRepository)
	footPrintController := controller.NewFootPrintController(footPrintInteractor)

	userRepository := repository.NewUserRepository(db, validate)
	userInteractor := interactor.NewUserInteractor(userRepository, footPrintRepository)
	userController := controller.NewUserController(userInteractor)

	likeRepository := repository.NewLikeRepository(db, validate)
	likeInteractor := interactor.NewLikeInteractor(likeRepository)
	likeController := controller.NewLikeController(likeInteractor)

	router.GET("/foot_prints", func(c *gin.Context) { footPrintController.Index(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users/:uid", func(c *gin.Context) { userController.Show(c) })
	router.POST("/users/:uid/likes", func(c *gin.Context) { likeController.Create(c) })
	router.GET("/likes", func(c *gin.Context) { likeController.Index(c) })
	router.PUT("/likes/:sent_uesr_uid/next", func(c *gin.Context) { likeController.Next(c) })
	router.GET("/likes/recieved", func(c *gin.Context) { likeController.Create(c) })
	router.GET("/likes/sent", func(c *gin.Context) { likeController.Create(c) })

	router.Run(":8080")
}
