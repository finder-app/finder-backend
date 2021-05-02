package controller

import (
	"finder/interface/interactor"
	"net/http"
	"strconv"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

func NewUserController(uc interactor.UserInteractor) *userController {
	return &userController{uc}
}

func (uc *userController) Show(c Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	user, err := uc.userInteractor.GetUserByID(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
