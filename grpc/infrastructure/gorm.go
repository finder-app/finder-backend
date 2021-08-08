package infrastructure

import (
	"log"

	"github.com/finder-app/finder-backend/grpc/infrastructure/env"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewGormConnect() *gorm.DB {
	user := env.DB_USER
	password := env.DB_PASSWORD
	host := env.DB_HOST
	dbName := env.DB_NAME
	options := "charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	databaseUrl := user + ":" + password + "@" + host + "/" + dbName + "?" + options

	switch env.ENV {
	case "production":
		// NOTE: 実装中
		return nil
	default:
		// db, err := gorm.Open(config.SqlDriver, config.DatabaseUrl)
		db, err := gorm.Open(env.DB_DRIVER, databaseUrl)
		if err != nil {
			log.Fatal(err.Error())
		}
		return db
	}
}
