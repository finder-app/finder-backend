package controller

import (
	"finder/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type likeController struct {
	likeUsecase usecase.LikeUsecase
}

func NewLikeController(lu usecase.LikeUsecase) *likeController {
	return &likeController{
		likeUsecase: lu,
	}
}

func (c *likeController) Create(ctx *gin.Context) {
	sentUesrUid := ctx.Value("currentUserUid").(string)
	recievedUserUid := ctx.Param("uid")
	like, err := c.likeUsecase.CreateLike(sentUesrUid, recievedUserUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusCreated, like)
}

func (c *likeController) Index(ctx *gin.Context) {
	currentUserUid := ctx.Value("currentUserUid").(string)
	like, err := c.likeUsecase.GetOldestLikeByUid(currentUserUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, like)
}

func (c *likeController) Consent(ctx *gin.Context) {
	recievedUserUid := ctx.Value("currentUserUid").(string)
	sentUesrUid := ctx.Param("sent_uesr_uid")
	like, room, err := c.likeUsecase.Consent(recievedUserUid, sentUesrUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"like": like,
		"room": room,
	})
}

func (c *likeController) Next(ctx *gin.Context) {
	recievedUserUid := ctx.Value("currentUserUid").(string)
	sentUesrUid := ctx.Param("sent_uesr_uid")
	like, err := c.likeUsecase.GetNextUserByUid(recievedUserUid, sentUesrUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, like)
}
