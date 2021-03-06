package domain

import (
	"time"
)

type User struct {
	Uid       string `gorm:"primaryKey"`
	Email     string `validate:"required"`
	LastName  string `validate:"required"`
	FirstName string `validate:"required"`
	Gender    string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Thumbnail string

	// NOTE: カラムを持たないfield
	Liked bool
}

func (user *User) FullName() string {
	return user.LastName + user.FirstName
}
