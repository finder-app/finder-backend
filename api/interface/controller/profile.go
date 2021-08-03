package controller

import (
	"api/interface/request_helper"
	"api/pb"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	// NOTE:
	file, _, _ := ctx.Request.FormFile("thumbnail")
	defer file.Close()
	if file != nil {
		ThumbnailUrl, err := s3upload(file)
		if err != nil {
			ErrorResponse(ctx, http.StatusInternalServerError, err)
			return
		}
		requestUser.Thumbnail = ThumbnailUrl
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

// varとかはinitで呼び出してあげるのが良さそう。1回しか呼ばれなさそうだし！
var (
	accessKey  = os.Getenv("AWS_ACCESS_KEY")
	privateKey = os.Getenv("AWS_PRIVATE_KEY")
	region     = os.Getenv("AWS_REGION")
	// bucketNameが実際のs3、bucketがlocalstack
	// bucketName = os.Getenv("AWS_BUCKET_NAME")
	bucketName = os.Getenv("AWS_S3_BUCKET")
	endPoint   = os.Getenv("AWS_S3_ENDPOINT")

	// NOTE: s3のobject名になる。被らせたくないのでuuidにする。230京分の1で同じuuidが生成される
	key         = uuid.New().String()
	contentType = "image/png"
)

func s3upload(file multipart.File) (location string, err error) {
	creds := credentials.NewStaticCredentials(accessKey, privateKey, "")
	sess := session.Must(session.NewSession(newAwsConfig(creds)))
	uploader := s3manager.NewUploader(sess)
	uploadInput := newUploadInput(file)
	s3uploadOutput, err := uploader.Upload(uploadInput)
	if err != nil {
		fmt.Println("errerrerrerr", err)
		return "", err
	}

	switch env := os.Getenv("ENV"); env {
	case "production":
		return s3uploadOutput.Location, nil
	default:
		// http://localstack:4566/bucket/8bd48331-3953-4d08-9e9b-dc543fc415bd
		// location = s3uploadOutput.Location
		location := strings.Replace(s3uploadOutput.Location, "localstack", "localhost", 1)
		fmt.Println("location", location)
		return location, nil
	}
}

func newAwsConfig(creds *credentials.Credentials) *aws.Config {
	switch env := os.Getenv("ENV"); env {
	// switch env := "production"; env {
	case "production":
		return &aws.Config{
			Credentials: creds,
			Region:      aws.String(region),
		}
	default:
		return &aws.Config{
			Credentials:      creds,
			Region:           aws.String(region),
			Endpoint:         aws.String(endPoint),
			S3ForcePathStyle: aws.Bool(true),
		}
	}
}

func newUploadInput(file multipart.File) *s3manager.UploadInput {
	return &s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        file,
		ACL:         aws.String(s3.BucketCannedACLPublicRead),
		ContentType: aws.String(contentType),
	}
}
