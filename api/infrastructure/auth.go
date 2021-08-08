package infrastructure

import (
	"api/infrastructure/firebase"
	"api/interface/controller"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(firebaseClient firebase.FirebaseClient) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		authorization := ginContext.Request.Header.Get("Authorization")
		idToken := strings.Replace(authorization, "Bearer ", "", 1)
		ctx := ginContext.Request.Context()
		token, err := firebaseClient.VerifyIDToken(ctx, idToken)
		if err != nil {
			// NOTE: GraphQLへのリクエストの時
			// if ginContext.Request.URL.Path == "/query" {
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
			// 	ginContext.Request = ginContext.Request.WithContext(ctx)
			// 	ginContext.Next()
			// 	return
			// }
			// NOTE: REST APIのリクエストの時
			controller.ErrorResponse(ginContext, http.StatusUnauthorized, err)
			return
		}
		// NOTE: ginContext.Setする時にinterfaceにされるけど、型が分かるからキャストする
		currentUserUid := token.Claims["user_id"].(string)
		fmt.Println("-------------------------------------------")
		fmt.Printf("\ncurrentUserUid:%v\nemail:%v\n", currentUserUid, token.Claims["email"])
		// NOTE: REST API時はginContext.Setでok
		ginContext.Set("currentUserUid", currentUserUid)

		// NOTE: GraphQLの時はserver.HTTPしてるから、ginContext.RequestにWithContextする必要あり
		ctx = context.WithValue(ctx, "currentUserUid", currentUserUid)
		ginContext.Request = ginContext.Request.WithContext(ctx)
		ginContext.Next()
	}
}
