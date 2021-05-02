package controller

import (
	"finder/interface/usecase"
	"strconv"
)

type userController struct {
	userInteractor usecase.UserUsecase
}

func NewUserController(uc usecase.UserUsecase) *userController {
	return &userController{uc}
}

func (uc *userController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := uc.userInteractor.GetUserByID(id)
	if err != nil {
		// c.AbortWithJSON使いたい〜
		c.JSON()
		return nil, err
	}
	return user, nil
}
