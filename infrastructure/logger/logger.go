package logger

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func NewLogger(db *gorm.DB) {
	// NOTE: ログ出力先
	file, err := os.OpenFile("./infrastructure/logger/logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	// logにファイルを出力するようにする
	log.SetOutput(file)
	db.LogMode(true)
	db.SetLogger(log.New(file, "", 0))
}
