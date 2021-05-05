package domain

import (
	"time"
)

type FootPrint struct {
	ID         uint
	VisitorUid string
	UserUid    string
	Unread     bool
	// これは取れるやつ。自分自身のIDがね
	// User User `gorm:"foreignKey:UserUid"`

	// User User `gorm:"foreignKey:VisitorUid"`
	// User User `gorm:"foreignKey:VisitorUid references:Uid"`

	// error
	// User User `gorm:"references:Uid"`

	//
	// Visitor User
	Visitor User `gorm:"foreignKey:VisitorUid"`

	// これが理想だけど、できない！なぜ！？
	// User User `gorm:"foreignKey:VisitorUid"`
	// User User `gorm:"foreignKey:VisitorUid references:Uid"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
