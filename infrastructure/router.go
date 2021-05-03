package infrastructure

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	// NOTE: Cors()を先に設定すること。Auth()で401の場合は処理を中断させるため
	router.Use(Cors())
	router.Use(Auth())
	return router
}
