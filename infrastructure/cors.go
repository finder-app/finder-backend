package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// NOTE: 一旦*で運用する
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, PATCH, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			// NOTE: c.Statusだとpreflight request時にresponseを返せないので、c.JSONに変更
			c.JSON(http.StatusNoContent, nil)
			// c.Status(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
