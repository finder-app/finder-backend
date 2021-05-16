package controller

import (
	"github.com/gin-gonic/gin"
)

func ErrorResponse(ctx *gin.Context, statusCode int, err error) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"errorMessage": string(err.Error()),
	})
}
