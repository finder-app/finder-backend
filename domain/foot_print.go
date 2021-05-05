package domain

import (
	"time"
)

type FootPrint struct {
	ID         uint
	VisitorUid string
	UserUid    string
	Unread     bool
	User       User `gorm:"foreignKey:UserUid"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
