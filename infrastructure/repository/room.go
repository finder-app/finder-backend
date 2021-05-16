package repository

import (
	"finder/domain"

	"github.com/jinzhu/gorm"
)

type RoomRepository interface {
	CreateRoom(tx *gorm.DB) (domain.Room, error)
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *roomRepository {
	return &roomRepository{
		db: db,
	}
}

func (r *roomRepository) CreateRoom(tx *gorm.DB) (domain.Room, error) {
	room := domain.Room{}
	if err := tx.Create(&room).Error; err != nil {
		return domain.Room{}, err
	}
	return room, nil
}
