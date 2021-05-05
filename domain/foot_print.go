package domain

import (
	"time"
)

type FootPrint struct {
	ID         uint   `gorm:"primaryKey"`
	VisitorUid string `gorm:"references:Uid foreignKey:UserUid"`
	UserUid    string `gorm:"references:Uid foreignKey:UserUid"`
	Unread     bool
	// これは取れるやつ。自分自身のIDがね
	User User `gorm:"foreignKey:UserUid"`

	// User User `gorm:"foreignKey:VisitorUid"`
	// User User `gorm:"foreignKey:VisitorUid references:Uid"`

	// これが理想だけど、できない！なぜ！？
	// User User `gorm:"references:Uid foreignKey:UserUid"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
