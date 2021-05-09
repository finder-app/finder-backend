package router

import (
	"finder/interface/controller"

	"github.com/gin-gonic/gin"
)

func (r *Router) Users(userController *controller.UserController) {
	r.Engine.GET("/users", func(c *gin.Context) { userController.Index(c) })
	r.Engine.POST("/users", func(c *gin.Context) { userController.Create(c) })
	r.Engine.GET("/users/:uid", func(c *gin.Context) { userController.Show(c) })
}
