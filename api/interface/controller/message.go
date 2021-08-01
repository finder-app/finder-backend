package controller

import (
	"api/pb"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	messageClint pb.MessageServiceClient
}

func NewMessageController(messageClint pb.MessageServiceClient) *MessageController {
	return &MessageController{
		messageClint: messageClint,
	}
}

func (c *MessageController) Index(ctx *gin.Context) {
	roomId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	req := &pb.GetMessagesReq{
		RoomId:         roomId,
		CurrentUserUid: ctx.Value("currentUserUid").(string),
	}
	res, err := c.messageClint.GetMessages(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	// res := roomId
	ctx.JSON(http.StatusOK, res)
}