package infrastructure

import (
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
		authHeader := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		ctx := c.Request.Context()
		token, err := client.VerifyIDToken(ctx, idToken)
		if err != nil {
			fmt.Printf("infrastructure/auth.go error: %v\n", err)
			// NOTE: tokenが確認場合は意図的に401エラーを返して処理を中断させる
			// returnがないと関数から抜け出せず、後続の処理が実行される
			controller.ErrorResponse(c, http.StatusUnauthorized, err)
			return
		}
		// NOTE: c.Setする時にinterfaceにされるけど、型が分かるからキャストする
		currentUserUid := token.Claims["user_id"].(string)
		fmt.Printf("%v debug start %v\n", "--------", "--------")
		fmt.Printf("currentUserUid:%v\nemail:%v\n", currentUserUid, token.Claims["email"])
		fmt.Printf("%v debug end %v\n", "--------", "--------")
		c.Set("currentUserUid", currentUserUid)
		c.Next()
	}
}
