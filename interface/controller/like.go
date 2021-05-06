package controller

import (
	"finder/usecase/interactor"
	"net/http"
)

type likeController struct {
	likeInteractor interactor.LikeInteractor
}

func NewLikeController(ri interactor.LikeInteractor) *likeController {
	return &likeController{
		likeInteractor: ri,
	}
}

func (c *likeController) Create(ctx Context) {
	SentUesrUid := ctx.Value("currentUserUid").(string)
	recievedUserUid := ctx.Param("uid")
	like, err := c.likeInteractor.CreateLike(SentUesrUid, recievedUserUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, like)
}
