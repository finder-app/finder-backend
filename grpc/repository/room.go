package repository

import (
	"grpc/domain"

	"github.com/jinzhu/gorm"
)

type RoomRepository interface {
	CreateRoom(tx *gorm.DB, room *domain.Room) error
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *roomRepository {
	return &roomRepository{
		db: db,
	}
}

func (r *roomRepository) CreateRoom(tx *gorm.DB, room *domain.Room) error {
	if err := tx.Create(room).Error; err != nil {
		return err
	}
	return nil
}
