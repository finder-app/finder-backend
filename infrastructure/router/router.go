package router

import (
	"finder/infrastructure"
	"finder/interface/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	engine := gin.Default()
	// NOTE: Cors()を先に設定すること。Auth()で401の場合は処理を中断させるため
	engine.Use(infrastructure.Cors())
	engine.Use(infrastructure.Auth())
	return &Router{
		Engine: engine,
	}
}

func (r *Router) Users(userController *controller.UserController) {
	r.Engine.GET("/users", userController.Index)
	// r.Engine.GET("/users", func(c *gin.Context) {
	// 	userController.Index(c)
	// })
	r.Engine.POST("/users", func(c *gin.Context) {
		userController.Create(c)
	})
	r.Engine.GET("/users/:uid", func(c *gin.Context) {
		userController.Show(c)
	})
}

func (r *Router) Profile(profileController *controller.ProfileController) {
	r.Engine.GET("/profile", func(c *gin.Context) {
		profileController.Index(c)
	})
	r.Engine.PUT("/profile", func(c *gin.Context) {
		profileController.Update(c)
	})
}
