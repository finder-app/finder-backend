package repository

import (
	"finder/domain"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type LikeRepository interface {
	CreateLike(sentUesrUid string, recievedUserUid string) (*domain.Like, error)
	// GetLikesByUid(sentUesrUid string) ([]domain.Like, error)
	// UpdateToAlreadyRead(sentUesrUid string) error
}

type likeRepository struct {
	db       *gorm.DB
	validate *validator.Validate
}

func NewLikeRepository(db *gorm.DB, validate *validator.Validate) *likeRepository {
	return &likeRepository{
		db:       db,
		validate: validate,
	}
}

func (r *likeRepository) CreateLike(sentUesrUid string, recievedUserUid string) (*domain.Like, error) {
	like := &domain.Like{
		SentUserUid:     sentUesrUid,
		RecievedUserUid: recievedUserUid,
		Unread:          true,
	}
	fmt.Println(like)
	if err := r.db.Create(like).Error; err != nil {
		return nil, err
	}
	return like, nil
}
