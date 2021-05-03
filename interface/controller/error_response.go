package controller

func ErrorResponse(ctx Context, statusCode int, err error) {
	type errorResponse map[string]interface{}
	ctx.AbortWithStatusJSON(statusCode, errorResponse{
		"errorMessage": string(err.Error()),
	})
}
