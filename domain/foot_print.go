package domain

import (
	"time"
)

type FootPrint struct {
	VisitorUid string
	UserUid    string
	Unread     bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
