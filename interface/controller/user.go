package controller

import (
	"finder/domain"
	"finder/usecase/interactor"
	"net/http"
	"strconv"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

func NewUserController(uc interactor.UserInteractor) *userController {
	return &userController{uc}
}

func (uc *userController) Index(ctx Context) {
	user, err := uc.userInteractor.GetUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *userController) Create(ctx Context) {
	user := &domain.User{}
	ctx.BindJSON(user)
	user, err := uc.userInteractor.CreateUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (uc *userController) Show(ctx Context) {
	userID, _ := strconv.Atoi(ctx.Param("id"))
	user, err := uc.userInteractor.GetUserByID(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
