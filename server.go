package main

import (
	"finder/infrastructure"
	"finder/interface/controller"
	"finder/interface/repository"
	"finder/usecase/interactor"
)

func main() {
	// validationもinfrastructureに作れ
	db := infrastructure.NewGormConnect()

	userRepository := repository.NewUserRepository(db)
	userInteractor := interactor.NewUserInteractor(userRepository)
	userController := controller.NewUserController(userInteractor)
	// userController := controllers.NewUserController(db)

	// useCaseを抽象化する意味が分からん
	router := infrastructure.NewRouter()
	userRouter := router.Group("users")
	{
		userRouter.GET("/:id", userController.Show)
	}
	// router.GET("/users/:id", userController.Show)

	router.Run(":8080")
}
