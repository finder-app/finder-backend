package domain

import (
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	gorm.Model
	Name  string `validate:"required"`
	Email string `validate:"required"`
}
