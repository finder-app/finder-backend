package infrastructure

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DBを盛ったstructには、validateも含ませよう。
// ライブラリだからinfra層に隔離

func NewGormConnect() *gorm.DB {
	driver := os.Getenv("DB_DRIVER")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	switch env := os.Getenv("ENV"); env {
	case "development":
		dbHost := os.Getenv("DB_HOST")
		databaseUrl := user + ":" + password + "@" + dbHost + "/" + dbName + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
		db, err := gorm.Open(driver, databaseUrl)
		if err != nil {
			log.Fatal(err)
		}
		return db
	}
	return nil
}

// func DatabaseInit() {
// 	var err error
// 	db, err = gorm.Open(config.SQLDriver, config.DatabaseURL)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	// バリデーション
// 	validate = validator.New()

// 	// SQLログ出力先
// 	file, err := os.OpenFile("./db/sql_log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	log.SetOutput(file)
// 	db.LogMode(true)
// 	db.SetLogger(log.New(file, "", 0))
// }
