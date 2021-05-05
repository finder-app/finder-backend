package domain

import (
	"time"
)

type FootPrint struct {
	ID         uint
	VisitorUid string
	UserUid    string
	Unread     bool
	// これは取れるやつ
	User User `gorm:"foreignKey:UserUid"`

	// User User `gorm:"foreignKey:VisitorUid"`
	// User User `gorm:"foreignKey:VisitorUid references:Uid"`

	// これが理想だけど、できない！なぜ！？
	// User User `gorm:"references:Uid"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
