package controller

import (
	"errors"
	"finder/domain"
	"finder/usecase/interactor"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserController struct {
	userInteractor interactor.UserInteractor
}

func NewUserController(ui interactor.UserInteractor) *UserController {
	return &UserController{
		userInteractor: ui,
	}
}

func (c *UserController) Index(ctx *gin.Context) {
	currentUserUid := ctx.Request.Header.Get("currentUserUid")
	user, err := c.userInteractor.GetUsersByUid(currentUserUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) Create(ctx Context) {
	user := &domain.User{}
	if err := ctx.BindJSON(user); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
	}
	user, err := c.userInteractor.CreateUser(user)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) Show(ctx Context) {
	VisitorUid := ctx.Value("currentUserUid").(string)
	uid := ctx.Param("uid")
	user, err := c.userInteractor.GetUserByUid(uid, VisitorUid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrorResponse(ctx, http.StatusNotFound, err)
			return
		}
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
