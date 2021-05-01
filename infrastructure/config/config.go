package config

import "os"

var (
	SqlDriver   string
	DatabaseUrl string
)

func init() {
	SqlDriver = os.Getenv("DB_DRIVER")

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	database := user + ":" + password + "@" + host + "/" + dbName + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	DatabaseUrl = database
}
