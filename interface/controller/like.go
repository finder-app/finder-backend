package controller

import (
	"finder/usecase/interactor"
	"net/http"
)

type likeController struct {
	likeInteractor interactor.LikeInteractor
}

func NewLikeController(li interactor.LikeInteractor) *likeController {
	return &likeController{
		likeInteractor: li,
	}
}

func (c *likeController) Create(ctx Context) {
	sentUesrUid := ctx.Value("currentUserUid").(string)
	recievedUserUid := ctx.Param("uid")
	like, err := c.likeInteractor.CreateLike(sentUesrUid, recievedUserUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusCreated, like)
}

func (c *likeController) Index(ctx Context) {
	currentUserUid := ctx.Value("currentUserUid").(string)
	like, err := c.likeInteractor.GetOldestLikeByUid(currentUserUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, like)
}

func (c *likeController) Next(ctx Context) {
	recievedUserUid := ctx.Value("currentUserUid").(string)
	sentUesrUid := ctx.Param("sent_uesr_uid")
	like, err := c.likeInteractor.GetNextUserByUid(recievedUserUid, sentUesrUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, like)
}

func (c *likeController) Update(ctx Context) {
	recievedUserUid := ctx.Value("currentUserUid").(string)
	sentUesrUid := ctx.Param("sent_uesr_uid")
	err := c.likeInteractor.Consent(recievedUserUid, sentUesrUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "update!",
		"data":    "data",
	})
}
