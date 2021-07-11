package controller

import (
	"api/pb"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileClint pb.ProfileServiceClient
}

func NewProfileController(profileClint pb.ProfileServiceClient) *ProfileController {
	return &ProfileController{
		profileClint: profileClint,
	}
}

func (c *ProfileController) Index(ctx *gin.Context) {
	req := &pb.GetProfileReq{
		CurrentUserUid: ctx.Value("currentUserUid").(string),
	}
	res, err := c.profileClint.GetProfile(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *ProfileController) Update(ctx *gin.Context) {
	user := &pb.User{}
	if err := ctx.BindJSON(user); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	// NOTE: requestのuidとcurrent user uidが一致しなければreturn
	if user.Uid != ctx.Value("currentUserUid").(string) {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("bad request"))
		return
	}
	req := &pb.UpdateProfileReq{
		User: user,
	}
	res, err := c.profileClint.UpdateProfile(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}
