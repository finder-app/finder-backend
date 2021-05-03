package infrastructure

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(Cors())
	return router
}
