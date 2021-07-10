package domain

import (
	"time"
)

type Like struct {
	Id              uint64
	SentUserUid     string
	SentUser        User `gorm:"foreignKey:SentUserUid"`
	RecievedUserUid string
	Skipped         bool
	Consented       bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
