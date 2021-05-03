package infrastructure

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		app := NewFirebaseApp(ctx)
		NewAuthClient(app, ctx)
		client := NewAuthClient(app, ctx)
		authHeader := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(ctx, idToken)
		if err != nil {
			fmt.Println(err)
			c.Next()
			return
		}
		fmt.Println(token)
		fmt.Println(*token)
		c.Next()
	}
}
