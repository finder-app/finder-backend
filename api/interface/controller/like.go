package controller

import (
	"api/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	likeClinet pb.LikeServiceClient
}

func NewLikeController(likeClinet pb.LikeServiceClient) *LikeController {
	return &LikeController{
		likeClinet: likeClinet,
	}
}

func (c *LikeController) Create(ctx *gin.Context) {
	req := &pb.CreateLikeReq{
		SentUserUid:     ctx.Value("currentUserUid").(string),
		RecievedUserUid: ctx.Param("uid"),
	}
	like, err := c.likeClinet.CreateLike(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusCreated, like)
}

func (c *LikeController) Index(ctx *gin.Context) {
	// currentUserUid := ctx.Value("currentUserUid").(string)
	// like, err := c.likeUsecase.GetOldestLikeByUid(currentUserUid)
	// if err != nil {
	// 	ErrorResponse(ctx, http.StatusNotFound, err)
	// 	return
	// }
	// ctx.JSON(http.StatusOK, like)
}

func (c *LikeController) Consent(ctx *gin.Context) {
	// recievedUserUid := ctx.Value("currentUserUid").(string)
	// sentUesrUid := ctx.Param("sent_uesr_uid")
	// like, room, err := c.likeUsecase.Consent(recievedUserUid, sentUesrUid)
	// if err != nil {
	// 	ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	// ctx.JSON(http.StatusCreated, gin.H{
	// 	"like": like,
	// 	"room": room,
	// })
}

func (c *LikeController) Next(ctx *gin.Context) {
	// recievedUserUid := ctx.Value("currentUserUid").(string)
	// sentUesrUid := ctx.Param("sent_uesr_uid")
	// like, err := c.likeUsecase.GetNextUserByUid(recievedUserUid, sentUesrUid)
	// if err != nil {
	// 	ErrorResponse(ctx, http.StatusNotFound, err)
	// 	return
	// }
	// ctx.JSON(http.StatusOK, like)
}
