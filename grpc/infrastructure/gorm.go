package infrastructure

import (
	"grpc/infrastructure/config"
	"log"

	"grpc/infrastructure/env"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewGormConnect() *gorm.DB {
	switch env.ENV {
	case "production":
		// NOTE: 実装中
		return nil
	default:
		// db, err := gorm.Open(config.SqlDriver, config.DatabaseUrl)
		db, err := gorm.Open(config.SqlDriver, config.DatabaseUrl)
		if err != nil {
			log.Fatal(err.Error())
		}
		return db
	}
}
