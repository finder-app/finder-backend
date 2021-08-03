package controller

import (
	"api/interface/request_helper"
	"api/pb"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

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
	stream       pb.ProfileService_TestImageClient
}

func NewProfileController(
	profileClint pb.ProfileServiceClient,
	stream pb.ProfileService_TestImageClient,
) *ProfileController {
	return &ProfileController{
		profileClint: profileClint,
		stream:       stream,
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

	file, _, _ := ctx.Request.FormFile("thumbnail")
	// defer file.Close()
	streaaaaaam(c.stream, file)
	// これをコメントインすると画像が投稿されちゃう
	// preSignS3(file)

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

func streaaaaaam(stream pb.ProfileService_TestImageClient, file multipart.File) error {
	fmt.Println("streaaaaaam")
	buf := make([]byte, 1024)
	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		stream.Send(&pb.TestImageReq{Image: buf})
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Println("respresprespresp", resp)
	return nil
}

func preSignS3(file multipart.File) {
	var (
		accessKey  = os.Getenv("AWS_ACCESS_KEY")
		privateKey = os.Getenv("AWS_PRIVATE_KEY")
		region     = os.Getenv("AWS_REGION")
		bucketName = os.Getenv("AWS_BUCKET_NAME")
	)

	creds := credentials.NewStaticCredentials(accessKey, privateKey, "")
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(region),
	}))
	uploader := s3manager.NewUploader(sess)
	s3uploadOutput, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(uuid.New().String()),
		Body:        file,
		ACL:         aws.String(s3.BucketCannedACLPublicRead),
		ContentType: aws.String("image/jpeg"),
	})
	// fmt.Println(s3uploadOutput)
	fmt.Println(s3uploadOutput.Location)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("done")
}
