package infrastructure

import (
	"context"
	"finder/interface/controller"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		app := NewFirebaseApp()
		client := NewAuthClient(app)
		fmt.Println(client)
		authHeader := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		if _, err := client.VerifyIDToken(context.Background(), idToken); err != nil {
			fmt.Println(err)
			controller.ErrorResponse(c, http.StatusUnauthorized, err)
			return
		}
		// c.Nextいらないんじゃね？
		c.Next()
	}
}
