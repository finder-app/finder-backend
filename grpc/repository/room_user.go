package repository

import (
	"grpc/domain"

	"github.com/jinzhu/gorm"
)

type RoomUserRepository interface {
	CreateRoomUser(tx *gorm.DB, roomUser domain.RoomUser) error
}

type roomUserRepository struct {
	db *gorm.DB
}

func NewRoomUserRepository(db *gorm.DB) *roomUserRepository {
	return &roomUserRepository{
		db: db,
	}
}

func (r *roomUserRepository) CreateRoomUser(tx *gorm.DB, roomUser domain.RoomUser) error {
	if err := tx.Table("rooms_users").Create(&roomUser).Error; err != nil {
		return err
	}
	return nil
}
