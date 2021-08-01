package domain

import (
	"time"
)

type Room struct {
	Id        uint64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewRoom() *Room {
	return &Room{}
}
