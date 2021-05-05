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
	// 足跡一覧を確認する機能をやる時に実装！
	footPrintController := controller.NewFootPrintController(footPrintInteractor)

	footPrintRouter := router.Group("foot_prints")
	{
		footPrintRouter.GET("", func(c *gin.Context) { footPrintController.Index(c) })
	}

	userRepository := repository.NewUserRepository(db, validate)
	userInteractor := interactor.NewUserInteractor(userRepository, footPrintRepository)
	userController := controller.NewUserController(userInteractor)

	userRouter := router.Group("users")
	{
		userRouter.GET("/index", func(c *gin.Context) { userController.Index(c) })
		userRouter.POST("/create", func(c *gin.Context) { userController.Create(c) })
		userRouter.GET("/:uid", func(c *gin.Context) { userController.Show(c) })
	}
	router.Run(":8080")
}
