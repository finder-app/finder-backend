package domain

import (
	"time"
)

type Room struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
