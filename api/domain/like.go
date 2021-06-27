package domain

import (
	"time"
)

type Like struct {
	SentUserUid     string
	RecievedUserUid string
	Skipped         bool
	Consented       bool
	SentUser        User `gorm:"foreignKey:SentUserUid"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
