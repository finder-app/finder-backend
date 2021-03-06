package domain

import (
	"time"
)

type FootPrint struct {
	Id         uint64
	VisitorUid string
	UserUid    string
	Unread     bool
	// NOTE: 理由は説明できないけど何か取得できた。Visitorっていう仮想のstructを作ったイメージ
	Visitor *User `gorm:"foreignKey:VisitorUid"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
