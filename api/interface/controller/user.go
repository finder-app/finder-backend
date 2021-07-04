package controller

import (
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
		CurrentUserUid: ctx.Value("currentUserUid").(string),
	}
	users, err := c.userClient.GetUsers(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) Create(ctx *gin.Context) {
	reqUser := &pb.User{}
	if err := ctx.BindJSON(reqUser); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	req := &pb.CreateUserReq{
		User: reqUser,
	}
	user, err := c.userClient.CreateUser(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) Show(ctx *gin.Context) {
	// user, err := c.userUsecase.GetUserByUid(uid, VisitorUid)
	req := &pb.GetUserByUidReq{
		Uid:        ctx.Param("uid"),
		VisitorUid: ctx.Value("currentUserUid").(string),
	}
	user, err := c.userClient.GetUserByUid(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
