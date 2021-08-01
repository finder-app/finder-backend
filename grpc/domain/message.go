package domain

import "time"

type Message struct {
	Id        uint64
	RoomId    uint64
	UserUid   string
	Text      string
	Unread    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
