package repository

import (
	"github.com/finder-app/finder-backend/grpc/domain"

	"github.com/jinzhu/gorm"
)

type MessageRepository interface {
	GetMessages(roomId uint64) ([]*domain.Message, error)
	// CreateMessage(tx *gorm.DB, message *domain.Message) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{
		db: db,
	}
}

func (r *messageRepository) GetMessages(roomId uint64) ([]*domain.Message, error) {
	var messages []*domain.Message
	query := `SELECT * FROM messages WHERE room_id = ?`
	if err := r.db.Raw(query, roomId).Scan(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

// func (r *messageRepository) CreateMessage(tx *gorm.DB, message *domain.Message) error {
// 	if err := tx.Create(message).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
