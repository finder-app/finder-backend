package config

import (
	"os"
)

var (
	SQLDriver   string
	DatabaseURL string
)

func init() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	// dockerのネットワークを経由してアクセスするから、protocolはコンテナ名にしろ
	protocol := os.Getenv("DB_PROTOCOL")
	dbName := os.Getenv("DB_NAME")
	database := username + ":" + password + "@" + protocol + "/" + dbName + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	DatabaseURL = database
	SQLDriver = os.Getenv("DB_DRIVER")
}
