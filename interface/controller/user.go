package controller

import (
	"errors"
	"finder/domain"
	"finder/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) *UserController {
	return &UserController{
		userUsecase: uu,
	}
}

func (c *UserController) Index(ctx *gin.Context) {
	currentUserUid := ctx.Value("currentUserUid").(string)
	users, err := c.userUsecase.GetUsersByUid(currentUserUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) Create(ctx *gin.Context) {
	user := &domain.User{}
	if err := ctx.BindJSON(user); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	user, err := c.userUsecase.CreateUser(user)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) Show(ctx *gin.Context) {
	VisitorUid := ctx.Value("currentUserUid").(string)
	uid := ctx.Param("uid")
	user, err := c.userUsecase.GetUserByUid(uid, VisitorUid)
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
