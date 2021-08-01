package repository

import (
	"grpc/domain"

	"github.com/jinzhu/gorm"
)

type RoomRepository interface {
	CreateRoom(tx *gorm.DB, room *domain.Room) error
	GetRooms(currentUserUid string) ([]*domain.Room, error)
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

func (r *roomRepository) GetRooms(currentUserUid string) ([]*domain.Room, error) {
	var rooms []*domain.Room
	query := `SELECT * FROM rooms
						INNER JOIN rooms_users ON rooms_users.room_id = rooms.id
						WHERE rooms_users.user_uid = ?`
	if err := r.db.Raw(query, currentUserUid).Scan(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}
