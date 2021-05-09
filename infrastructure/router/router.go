package router

import (
	"finder/infrastructure"

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
