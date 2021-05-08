package repository

import (
	"finder/domain"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type LikeRepository interface {
	CreateLike(sentUesrUid string, recievedUserUid string) (*domain.Like, error)
	GetOldestLikeByUid(currentUserUid string) (*domain.Like, error)
	NopeUserByUid(recievedUserUid string, sentUesrUid string) error
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
	}
	if err := r.db.Create(like).Error; err != nil {
		return nil, err
	}
	return like, nil
}

func (r *likeRepository) GetOldestLikeByUid(currentUserUid string) (*domain.Like, error) {
	query := `SELECT id FROM likes
	WHERE recieved_user_uid = ?
	AND skipped = 0
	AND consented = 0
	ORDER BY CAST(created_at AS DATE) ASC
	LIMIT 1`
	var result struct {
		ID int
	}
	if err := r.db.Raw(query, currentUserUid).Scan(&result).Error; err != nil {
		return nil, err
	}
	like := &domain.Like{}
	if err := r.db.Model(domain.Like{}).Preload("SentUser").Take(like, result.ID).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return like, nil
}

func (r *likeRepository) NopeUserByUid(recievedUserUid string, sentUesrUid string) error {
	query := `UPDATE likes SET skipped = 1
		WHERE recieved_user_uid = ? AND sent_user_uid = ?`
	if err := r.db.Exec(query, recievedUserUid, sentUesrUid).Error; err != nil {
		return err
	}
	return nil
}
