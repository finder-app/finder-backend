package main

import (
	"finder/infrastructure"
	"finder/interface/controller"
	"finder/interface/repository"
	"finder/usecase/interactor"

	"github.com/gin-gonic/gin"
)

func main() {
	db := infrastructure.NewGormConnect()
	validate := infrastructure.NewValidator()

	userRepository := repository.NewUserRepository(db, validate)
	userInteractor := interactor.NewUserInteractor(userRepository)
	userController := controller.NewUserController(userInteractor)

	router := infrastructure.NewRouter()
	userRouter := router.Group("users")
	{
		userRouter.GET("/index", func(c *gin.Context) { userController.Index(c) })
		userRouter.POST("/create", func(c *gin.Context) { userController.Create(c) })
		userRouter.GET("/:id", func(c *gin.Context) { userController.Show(c) })
	}

	router.Run(":8080")
}
