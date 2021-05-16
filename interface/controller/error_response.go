package controller

import (
	"github.com/gin-gonic/gin"
)

func ErrorResponse(ctx *gin.Context, statusCode int, err error) {
	type errorResponse map[string]interface{}
	ctx.AbortWithStatusJSON(statusCode, errorResponse{
		"errorMessage": string(err.Error()),
	})
}
