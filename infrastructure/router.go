package infrastructure

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	// TODO: gin.Loggerの挙動は要検証
	router.Use(gin.Logger())
	router.Use(Auth())
	router.Use(Cors())
	return router
}
