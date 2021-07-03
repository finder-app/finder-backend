package controller

import (
	"finder/domain"
	"finder/pb"
	"finder/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecase.UserUsecase
	userClient  pb.UserServiceClient
}

func NewUserController(uu usecase.UserUsecase, userClient pb.UserServiceClient) *UserController {
	return &UserController{
		userUsecase: uu,
		userClient:  userClient,
	}
}

func (c *UserController) Index(ctx *gin.Context) {
	// NOTE: gRPCに移行
	// currentUserUid := ctx.Value("currentUserUid").(string)
	// users, err := c.userUsecase.GetUsersByUid(currentUserUid)
	req := &pb.GetUsersReq{
		Uid: ctx.Value("currentUserUid").(string),
	}
	users, err := c.userClient.GetUsers(ctx, req)
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
		ErrorResponse(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
