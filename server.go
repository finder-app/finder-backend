package main

import (
	"finder/infrastructure"
	"finder/interface/repository"
	"finder/usecase/interactor"
)

func main() {
	// validationもinfrastructureに作れ
	db := infrastructure.NewGormConnect()
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userInteractor := interactor.NewUserInteractor(userRepository)
	// userController := controller.NewUserController(userInteractor)
	// userController := controllers.NewUserController(db)

	router := infrastructure.NewRouter()
	router.Run(":8080")
}
