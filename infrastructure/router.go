package infrastructure

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(Cors())
	// authつけると動かないから一旦保留
	// router.Use(Auth())
	return router
}
