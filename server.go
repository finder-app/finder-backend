package main

import (
	"finder/infrastructure"
	"finder/interface/controller"
	"finder/interface/repository"
	"finder/usecase/interactor"

	"github.com/gin-gonic/gin"
)

func main() {
	// validationもinfrastructureに作れ
	db := infrastructure.NewGormConnect()

	userRepository := repository.NewUserRepository(db)
	userInteractor := interactor.NewUserInteractor(userRepository)
	userController := controller.NewUserController(userInteractor)

	router := infrastructure.NewRouter()
	userRouter := router.Group("users")
	{
		userRouter.GET("/:id", func(c *gin.Context) {
			userController.Show(c)
		})
	}

	router.Run(":8080")
}
