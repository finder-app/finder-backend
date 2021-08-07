package config

import "grpc/infrastructure/env"

var (
	SqlDriver   string
	DatabaseUrl string
)

func init() {
	SqlDriver = env.DB_DRIVER

	user := env.DB_USER
	password := env.DB_PASSWORD
	host := env.DB_HOST
	dbName := env.DB_NAME
	options := "charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	database := user + ":" + password + "@" + host + "/" + dbName + "?" + options
	DatabaseUrl = database
}
