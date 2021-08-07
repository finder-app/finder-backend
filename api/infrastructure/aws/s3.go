package aws

import (
	"mime/multipart"
	"strings"

	"api/infrastructure/env"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

var (
	uploader *s3manager.Uploader
)

type (
	S3uploader interface {
		Upload(file multipart.File) (location string, err error)
	}
	s3uploader struct {
		uploader *s3manager.Uploader
	}
)

// NOTE: initでuploaderを初期化
func init() {
	accessKey := env.AWS_ACCESS_KEY
	privateKey := env.AWS_PRIVATE_KEY

	creds := credentials.NewStaticCredentials(accessKey, privateKey, "")
	sess := session.Must(session.NewSession(newAwsConfig(creds)))
	uploader = s3manager.NewUploader(sess)
}

// NOTE: init()で生成したs3managerUploaderを代入
func NewS3uplodaer() S3uploader {
	return &s3uploader{
		uploader: uploader,
	}
}

func newAwsConfig(creds *credentials.Credentials) *aws.Config {
	region := env.AWS_REGION

	switch env.ENV {
	// switch env := "production"; env {
	case "production":
		return &aws.Config{
			Credentials: creds,
			Region:      aws.String(region),
		}
	default:
		// NOTE: dev環境のみ使用
		endPoint := env.LOCALSTACK_ENDPOINT
		return &aws.Config{
			Credentials:      creds,
			Region:           aws.String(region),
			Endpoint:         aws.String(endPoint),
			S3ForcePathStyle: aws.Bool(true),
		}
	}
}

func (s *s3uploader) Upload(file multipart.File) (location string, err error) {
	uploadInput := s.newUploadInput(file)
	s3uploadOutput, err := s.uploader.Upload(uploadInput)
	if err != nil {
		return "", err
	}

	if env.ENV == "development" {
		return s.replaceLocation(s3uploadOutput.Location), nil
	}
	return s3uploadOutput.Location, nil
}

func (s *s3uploader) newUploadInput(file multipart.File) *s3manager.UploadInput {
	// bucketNameが実際のs3、bucketがlocalstack
	// bucketName = env.AWS_BUCKET_NAME
	bucketName := env.AWS_S3_BUCKET
	// NOTE: keyはs3のobject名になる。被らせたくないのでuuidにする。230京分の1で同じuuidが生成される
	key := uuid.New().String()
	contentType := "image/png"

	return &s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        file,
		ACL:         aws.String(s3.BucketCannedACLPublicRead),
		ContentType: aws.String(contentType),
	}
}

// NOTE: dev環境だったらURLのlocalstackをlocalhostに置換
// s3output.Locationがhttp://localstackで生成されるため
func (s *s3uploader) replaceLocation(location string) string {
	location = strings.Replace(location, "localstack", "localhost", 1)
	return location
}
