package infrastructure

import (
	"finder/interface/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	engine := gin.Default()
	// NOTE: Cors()を先に設定すること。Auth()で401の場合は処理を中断させるため
	engine.Use(Cors())
	engine.Use(Auth())
	return &Router{
		Engine: engine,
	}
}

func (r *Router) Users(userController *controller.UserController) {
	r.Engine.GET("/users", userController.Index)
	r.Engine.POST("/users", userController.Create)
	r.Engine.GET("/users/:uid", userController.Show)
}

func (r *Router) Profile(profileController *controller.ProfileController) {
	r.Engine.GET("/profile", profileController.Index)
	r.Engine.PUT("/profile", profileController.Update)
}
