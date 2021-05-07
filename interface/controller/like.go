package controller

import (
	"errors"
	"finder/usecase/interactor"
	"net/http"

	"github.com/jinzhu/gorm"
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
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, like)
}

func (c *likeController) Next(ctx Context) {
	recievedUserUid := ctx.Value("currentUserUid").(string)
	sentUesrUid := ctx.Param("sent_uesr_uid")
	like, err := c.likeInteractor.NextUserByUid(recievedUserUid, sentUesrUid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrorResponse(ctx, http.StatusNotFound, err)
			return
		}
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusOK, like)
}
