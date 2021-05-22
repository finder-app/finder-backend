package controller

import (
	"finder/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type footPrintController struct {
	footPrintUsecase usecase.FootPrintUsecase
}

func NewFootPrintController(fpu usecase.FootPrintUsecase) *footPrintController {
	return &footPrintController{
		footPrintUsecase: fpu,
	}
}

func (c *footPrintController) Index(ctx *gin.Context) {
	currentUserUid := ctx.Value("currentUserUid").(string)
	footPrints, err := c.footPrintUsecase.GetFootPrintsByUid(currentUserUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, footPrints)
}

func (c *footPrintController) UnreadCount(ctx *gin.Context) {
	currentUserUid := ctx.Value("currentUserUid").(string)
	unreadCount, err := c.footPrintUsecase.GetUnreadCount(currentUserUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, unreadCount)
}
