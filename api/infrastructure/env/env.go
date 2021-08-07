package env

import "os"

var (
	GOOGLE_SERVICE_ACCOUNT_JSON string
	ENV                         string
	DB_DRIVER                   string
	DB_USER                     string
	DB_PASSWORD                 string
	DB_HOST                     string
	DB_NAME                     string
	PORT                        string
	GRPC_SERVER_NAME            string
	GRPC_SERVER_PORT            string
	AWS_ACCESS_KEY              string
	AWS_PRIVATE_KEY             string
	AWS_REGION                  string
	LOCALSTACK_ENDPOINT         string
	AWS_BUCKET_NAME             string
	AWS_S3_BUCKET               string
)

func init() {
	GOOGLE_SERVICE_ACCOUNT_JSON = os.Getenv("GOOGLE_SERVICE_ACCOUNT_JSON")
	ENV = os.Getenv("ENV")
	DB_DRIVER = os.Getenv("DB_DRIVER")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	PORT = os.Getenv("PORT")
	GRPC_SERVER_NAME = os.Getenv("GRPC_SERVER_NAME")
	GRPC_SERVER_PORT = os.Getenv("GRPC_SERVER_PORT")
	AWS_ACCESS_KEY = os.Getenv("AWS_ACCESS_KEY")
	AWS_PRIVATE_KEY = os.Getenv("AWS_PRIVATE_KEY")
	AWS_REGION = os.Getenv("AWS_REGION")
	LOCALSTACK_ENDPOINT = os.Getenv("LOCALSTACK_ENDPOINT")
	AWS_BUCKET_NAME = os.Getenv("AWS_BUCKET_NAME")
	AWS_S3_BUCKET = os.Getenv("AWS_S3_BUCKET")
}
