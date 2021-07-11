package controller

import (
	"finder/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userClient pb.UserServiceClient
}

func NewUserController(userClient pb.UserServiceClient) *UserController {
	return &UserController{
		userClient: userClient,
	}
}

func (c *UserController) Index(ctx *gin.Context) {
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
	req := &pb.GetUserByUidReq{
		Uid:        ctx.Param("uid"),
		VisitorUid: ctx.Value("currentUserUid").(string),
	}
	res, err := c.userClient.GetUserByUid(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusNotFound, err)
		return
	}
	// NOTE: gRPCが生成したstructにomitemptyがあり、Likedがfalseだとkeyが返せない
	// なのでフロントでuser.likedが取得できず、いいねした時にいいね済みにならない
	// そのため、自前でstructを作って返す
	user := map[string]interface{}{
		"Uid":       res.User.Uid,
		"LastName":  res.User.LastName,
		"FirstName": res.User.FirstName,
		"FullName":  res.User.FullName,
		"Email":     res.User.Email,
		"Liked":     res.User.Liked,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
