package controller

import (
	"api/finder-protocol-buffers/pb"
	"api/infrastructure/aws"
	"api/interface/request_helper"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileClint pb.ProfileServiceClient
	s3uploader   aws.S3uploader
}

func NewProfileController(
	profileClint pb.ProfileServiceClient,
	s3uploader aws.S3uploader,
) *ProfileController {
	return &ProfileController{
		profileClint: profileClint,
		s3uploader:   s3uploader,
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
	// NOTE: フロントからFormDataで送っているのでBindJSONだと受け取れない
	requestUser := request_helper.NewRequestUser()
	if err := ctx.Bind(&requestUser); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	// NOTE: requestのuidとcurrentUserUidが一致しなければreturn
	if requestUser.Uid != ctx.Value("currentUserUid").(string) {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("bad request: invalid uid"))
		return
	}

	// NOTE: 画像が存在すればS3にアップしてuser.thumbnailにS3のURLを代入
	file, _, _ := ctx.Request.FormFile("thumbnail")
	defer file.Close()
	if file != nil {
		location, err := c.s3uploader.Upload(file)
		if err != nil {
			ErrorResponse(ctx, http.StatusInternalServerError, err)
			return
		}
		requestUser.Thumbnail = location
	}

	pbUser := request_helper.NewPbUser(requestUser)
	req := &pb.UpdateProfileReq{
		User: pbUser,
	}
	res, err := c.profileClint.UpdateProfile(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}
