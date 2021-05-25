package infrastructure

import (
	"finder/interface/controller"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
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

func (r *Router) GraphQL(server *handler.Server, playGroundHandler http.Handler) {
	r.Engine.POST("/query", func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
	})
	r.Engine.GET("/", func(c *gin.Context) {
		playGroundHandler.ServeHTTP(c.Writer, c.Request)
	})
}

// NOTE: routingのテストをするため、router配下に書くこと
func (r *Router) Users(userController *controller.UserController) {
	// GraphQL化
	// r.Engine.GET("/users", userController.Index)
	r.Engine.POST("/users", userController.Create)
	r.Engine.GET("/users/:uid", userController.Show)
}

func (r *Router) Profile(profileController *controller.ProfileController) {
	r.Engine.GET("/profile", profileController.Index)
	r.Engine.PUT("/profile", profileController.Update)
}
