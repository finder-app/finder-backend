package controller

import (
	"net/http"

	"github.com/finder-app/finder-backend/api/finder-protocol-buffers/pb"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	roomClint pb.RoomServiceClient
}

func NewRoomController(roomClint pb.RoomServiceClient) *RoomController {
	return &RoomController{
		roomClint: roomClint,
	}
}

func (c *RoomController) Index(ctx *gin.Context) {
	req := &pb.GetRoomsReq{
		CurrentUserUid: ctx.Value("currentUserUid").(string),
	}
	res, err := c.roomClint.GetRooms(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}
