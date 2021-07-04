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
			// NOTE: GraphQLへのリクエストの時
			// if c.Request.URL.Path == "/query" {
			// 	ctx = context.WithValue(
			// 		ctx,
			// 		"AuthenticationError",
			// 		&gqlerror.Error{
			// 			Message: err.Error(),
			// 			Extensions: map[string]interface{}{
			// 				"status": http.StatusUnauthorized,
			// 			},
			// 		},
			// 	)
			// 	c.Request = c.Request.WithContext(ctx)
			// 	c.Next()
			// 	return
			// }
			// NOTE: REST APIのリクエストの時
			controller.ErrorResponse(c, http.StatusUnauthorized, err)
			return
		}
		// NOTE: c.Setする時にinterfaceにされるけど、型が分かるからキャストする
		currentUserUid := token.Claims["user_id"].(string)
		fmt.Println("-------------------------------------------")
		fmt.Printf("\ncurrentUserUid:%v\nemail:%v\n", currentUserUid, token.Claims["email"])
		// NOTE: REST API時はc.Setでok
		c.Set("currentUserUid", currentUserUid)

		// NOTE: GraphQLの時はserver.HTTPしてるから、c.RequestにWithContextする必要あり
		ctx = context.WithValue(ctx, "currentUserUid", currentUserUid)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
