package controller

import (
	"finder/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type footPrintController struct {
	footPrintClient pb.FootPrintServiceClient
}

func NewFootPrintController(footPrintClient pb.FootPrintServiceClient) *footPrintController {
	return &footPrintController{
		footPrintClient: footPrintClient,
	}
}

func (c *footPrintController) Index(ctx *gin.Context) {
	// こっちも実装しないと駄目？
	// currentUserUid := ctx.Value("currentUserUid").(string)
	// footPrints, err := c.footPrintClient.GetFootPrints(ctx, currentUserUid)
	// // footPrints, err := c.footPrintUsecase.GetFootPrintsByUid(currentUserUid)
	// if err != nil {
	// 	ErrorResponse(ctx, http.StatusInternalServerError, err)
	// 	return
	// }
	// ctx.JSON(http.StatusOK, footPrints)
}

func (c *footPrintController) UnreadCount(ctx *gin.Context) {
	req := &pb.GetUnreadCountReq{
		CurrentUserUid: ctx.Value("currentUserUid").(string),
	}
	unreadCount, err := c.footPrintClient.GetUnreadCount(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, unreadCount)
}
