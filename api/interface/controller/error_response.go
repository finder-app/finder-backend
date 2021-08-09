package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrorRecordNotFound = errors.New("rpc error: code = Unknown desc = record not found")
)

func IsRecordNotFoundError(err error) bool {
	return err.Error() == ErrorRecordNotFound.Error()
}

func ErrorResponse(ctx *gin.Context, statusCode int, err error) {
	// NOTE: なるべくginに頼りたくないので、gin.Hを使わないようにした
	ctx.AbortWithStatusJSON(statusCode, map[string]interface{}{
		"errorMessage": err.Error(),
	})
	// ctx.AbortWithStatusJSON(statusCode, gin.H{
	// 	"errorMessage": err.Error(),
	// })
}
