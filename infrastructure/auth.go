package infrastructure

import (
	"finder/interface/controller"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	// NOTE: 毎回firebaseを呼び出してほしくないので、アプリ初期化時に宣言しておく
	app := NewFirebaseApp()
	client := NewAuthClient(app)
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		authHeader := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(ctx, idToken)
		if err != nil {
			controller.ErrorResponse(c, http.StatusUnauthorized, err)
			return
		}
		fmt.Println(token)
		c.Next()
	}
}
