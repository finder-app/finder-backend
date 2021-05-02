package domain

import (
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	gorm.Model
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"min=6,max=75"`
	TelephoneNumber string `json:"telephone_number" validate:"required"`
	Gender          int    `json:"gender" validate:"min=1,max=3"`
}
