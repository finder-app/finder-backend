package repository

import (
	"grpc/domain"

	"github.com/jinzhu/gorm"
)

type RoomUserRepository interface {
	CreateRoomUsers(tx *gorm.DB, roomUser1 *domain.RoomUser, roomUser2 *domain.RoomUser) error
	GetRoomUser(roomId uint64, currentUserUid string) (*domain.RoomUser, error)
}

type roomUserRepository struct {
	db *gorm.DB
}

func NewRoomUserRepository(db *gorm.DB) *roomUserRepository {
	return &roomUserRepository{
		db: db,
	}
}

func (r *roomUserRepository) CreateRoomUsers(tx *gorm.DB, roomUser1 *domain.RoomUser, roomUser2 *domain.RoomUser) error {
	if err := createRoomUser(tx, roomUser1); err != nil {
		return err
	}
	if err := createRoomUser(tx, roomUser2); err != nil {
		return err
	}
	return nil
}

func createRoomUser(tx *gorm.DB, roomUser *domain.RoomUser) error {
	if err := tx.Table("rooms_users").Create(roomUser).Error; err != nil {
		return err
	}
	return nil
}

func (r *roomUserRepository) GetRoomUser(roomId uint64, currentUserUid string) (*domain.RoomUser, error) {
	roomUser := &domain.RoomUser{}
	query := `SELECT * FROM rooms_users
						WHERE room_id = ? AND user_uid = ?`
	if err := r.db.Raw(query, roomId, currentUserUid).Scan(roomUser).Error; err != nil {
		return nil, err
	}
	return roomUser, nil
}
