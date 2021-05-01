package infrastructure

import (
	"finder/infrastructure/config"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewGormConnect() *gorm.DB {
	switch env := os.Getenv("ENV"); env {
	case "production":
		// NOTE: 実装中
		return nil
	default:
		db, err := gorm.Open(config.SqlDriver, config.DatabaseUrl)
		if err != nil {
			panic(err.Error())
		}

		// NOTE: SQLログ出力先
		file, err := os.OpenFile("./logger/sql_log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err.Error())
		}
		log.SetOutput(file)
		db.LogMode(true)
		db.SetLogger(log.New(file, "", 0))

		return db
	}
}
