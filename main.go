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

	userRepository := repository.NewUserRepository(db, validate)
	userInteractor := interactor.NewUserInteractor(userRepository)
	userController := controller.NewUserController(userInteractor)
	router.GET("/users/index", func(c *gin.Context) { userController.Index(c) })
	router.POST("/users/create", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	router.Run(":8080")
}
