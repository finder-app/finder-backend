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
		firebaseApp := NewFirebaseApp()
		authClient := NewAuthClient(firebaseApp)
		authorization := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authorization, "Bearer ", "", 1)
		ctx := c.Request.Context()
		token, err := authClient.VerifyIDToken(ctx, idToken)
		if err != nil {
			// NOTE: tokenが確認場合は意図的に401エラーを返して処理を中断させる
			// returnがないと関数から抜け出せず、後続の処理が実行される
			controller.ErrorResponse(c, http.StatusUnauthorized, err)
			return
		}
		// NOTE: c.Setする時にinterfaceにされるけど、型が分かるからキャストする
		currentUserUid := token.Claims["user_id"].(string)
		fmt.Printf("currentUserUid:%v\nemail:%v\n", currentUserUid, token.Claims["email"])
		// NOTE: controllerの時はc.Setでok
		c.Set("currentUserUid", currentUserUid)

		// NOTE: GraphQLの時はserver.HTTPしてるから、c.RequestにWithContextする必要あり
		ctx = context.WithValue(ctx, "currentUserUid", currentUserUid)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
