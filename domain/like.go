package domain

import (
	"time"
)

type Like struct {
	SentUserUid     string
	RecievedUserUid string
	Unread          bool
	Consented       bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
