package infrastructure

import (
	"errors"
	"finder/interface/controller"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// NOTE: フロントから認証済みのuidが送られてくる
		currentUserUid := c.Request.Header.Get("currentUserUid")
		if utf8.RuneCountInString(currentUserUid) == 0 {
			// NOTE: tokenが確認場合は意図的に401エラーを返して処理を中断させる
			// returnがないと関数から抜け出せず、後続の処理が実行される
			err := errors.New("infrastructure/auth.go error")
			controller.ErrorResponse(c, http.StatusUnauthorized, err)
			return
		}
		c.Next()
	}
}
