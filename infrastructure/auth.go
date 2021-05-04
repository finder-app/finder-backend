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

	// tokenのリセットで他でclientを使い回すかもだから、要検討
	return func(c *gin.Context) {
		app := NewFirebaseApp()
		client := NewAuthClient(app)
		ctx := c.Request.Context()
		authHeader := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		fmt.Println(idToken)
		token, err := client.VerifyIDToken(ctx, idToken)
		if err != nil {
			fmt.Println(err)
			// NOTE: tokenが確認場合は意図的に401エラーを返して処理を中断させる
			// returnがないと関数から抜け出せず、後続の処理が実行される
			controller.ErrorResponse(c, http.StatusUnauthorized, err)
			return
		}
		email := token.Firebase.Identities["email"]
		c.Set("email", email)
		c.Next()
	}
}
