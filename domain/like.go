package domain

import (
	"time"
)

type Like struct {
	SentUesrUid     string
	RecievedUserUid string
	Unread          bool
	Consented       bool
	SentUesr        User `gorm:"foreignKey:SentUesrUid"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
