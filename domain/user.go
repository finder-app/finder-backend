package domain

import (
	"time"
)

type User struct {
	Uid       string `gorm:"primaryKey"`
	Email     string `validate:"required"`
	LastName  string `validate:"required"`
	FirstName string `validate:"required"`
	IsMale    bool   `validate:"required"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
