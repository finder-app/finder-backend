package controller

import (
	"net/http"

	"github.com/finder-app/finder-backend/api/finder-protocol-buffers/pb"

	"github.com/gin-gonic/gin"
)

type FootPrintController struct {
	footPrintClient pb.FootPrintServiceClient
}

func NewFootPrintController(footPrintClient pb.FootPrintServiceClient) *FootPrintController {
	return &FootPrintController{
		footPrintClient: footPrintClient,
	}
}

func (c *FootPrintController) Index(ctx *gin.Context) {
	req := &pb.GetFootPrintsReq{
		CurrentUserUid: ctx.Value("currentUserUid").(string),
	}
	footPrints, err := c.footPrintClient.GetFootPrints(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, footPrints)
}

func (c *FootPrintController) UnreadCount(ctx *gin.Context) {
	req := &pb.GetUnreadCountReq{
		CurrentUserUid: ctx.Value("currentUserUid").(string),
	}
	res, err := c.footPrintClient.GetUnreadCount(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	// NOTE: gRPCのjson responseにomitemptyがあり 0 だとundefinedになるため
	ctx.JSON(http.StatusOK, res.UnreadCount)
}
