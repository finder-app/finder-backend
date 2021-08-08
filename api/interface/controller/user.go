package controller

import (
	"net/http"

	"github.com/finder-app/finder-backend/api/finder-protocol-buffers/pb"

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
	pbUser := &pb.User{}
	if err := ctx.BindJSON(pbUser); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	req := &pb.CreateUserReq{
		Uid:       pbUser.Uid,
		Email:     pbUser.Email,
		LastName:  pbUser.LastName,
		FirstName: pbUser.FirstName,
		Gender:    pbUser.Gender,
		Thumbnail: pbUser.Thumbnail,
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
	// keyがないとフロントでuser.likedが取得できず、いいねした時にいいね済みにならない
	// そのため、自前でstructを作って返す
	type responseUser struct {
		*pb.User
		// TODO: Likedのoptionがどこまで必要か要検証。
		Liked bool `json:"liked"`
		// Liked bool `protobuf:"varint,9,opt,name=liked,proto3" json:"liked"`
	}
	user := responseUser{
		User:  res.User,
		Liked: res.User.Liked,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
