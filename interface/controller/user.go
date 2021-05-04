package controller

import (
	"finder/domain"
	"finder/usecase/interactor"
	"net/http"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

func NewUserController(c interactor.UserInteractor) *userController {
	return &userController{
		userInteractor: c,
	}
}

func (c *userController) Index(ctx Context) {
	currentUserUid := ctx.Value("currentUserUid").(string)
	user, err := c.userInteractor.GetUsersByUid(currentUserUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *userController) Create(ctx Context) {
	user := &domain.User{}
	ctx.BindJSON(user)
	user, err := c.userInteractor.CreateUser(user)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (c *userController) Show(ctx Context) {
	uid := ctx.Param("uid")
	user, err := c.userInteractor.GetUserByUid(uid)
	if err != nil {
		ErrorResponse(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
