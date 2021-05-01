package db

import (
	"app/app/model"
	"app/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Migrate() {
	fmt.Println("------------ migrate database... ------------")
	db, err := gorm.Open(config.SQLDriver, config.DatabaseURL)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("migrate...")
	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.BookDetail{})
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.Language{})

	db.Model(&model.Comment{}).AddForeignKey("book_id", "books(id)", "CASCADE", "CASCADE")
	db.Model(&model.Comment{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&model.BookDetail{}).AddForeignKey("book_id", "books(id)", "CASCADE", "CASCADE")
	// db.Model(&model.Comment{}).RemoveForeignKey("book_id", "books(id)")
	fmt.Println("------------ finish migrate! ----------------")
}
