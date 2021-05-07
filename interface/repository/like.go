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
	like := &domain.Like{}
	query := `SELECT * FROM likes
		WHERE recieved_user_uid = ?
			AND skipped = 0
			AND consented = 0
		ORDER BY CAST(created_at AS DATE) ASC
		LIMIT 1`
	if err := r.db.Raw(query, currentUserUid).Scan(like).Error; err != nil {
		return nil, err
	}
	// NOTE: 生SQLを発行してるからか、上だとSentUserを取得できないので改めて取得。リファクタしたい
	if err := r.db.Model(domain.Like{}).Preload("SentUser").Take(like).Error; err != nil {
		return nil, err
	}
	return like, nil
}

func (r *likeRepository) NopeUserByUid(recievedUserUid string, sentUesrUid string) error {
	fmt.Println(recievedUserUid, sentUesrUid)
	query := `UPDATE likes SET skipped = 1
		WHERE recieved_user_uid = ? AND sent_user_uid = ?`
	result := r.db.Exec(query, recievedUserUid, sentUesrUid)
	if err := result.Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
