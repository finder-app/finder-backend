package repository

import (
	"grpc/domain"

	"github.com/jinzhu/gorm"
)

type LikeRepository interface {
	CreateLike(like *domain.Like) (*domain.Like, error)
	GetOldestLikeByUid(currentUserUid string) (*domain.Like, error)
	NopeUserByUid(recievedUserUid string, sentUesrUid string) error
	Begin() *gorm.DB
	Consent(tx *gorm.DB, like *domain.Like) error
}

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *likeRepository {
	return &likeRepository{
		db: db,
	}
}

func (r *likeRepository) CreateLike(like *domain.Like) (*domain.Like, error) {
	if err := r.db.Create(like).Error; err != nil {
		return nil, err
	}
	return like, nil
}

func (r *likeRepository) GetOldestLikeByUid(currentUserUid string) (*domain.Like, error) {
	like := &domain.Like{}
	result := r.db.Model(domain.Like{}).
		Where(`recieved_user_uid = ? AND skipped = 0 AND consented = 0`, currentUserUid).
		Order("CAST(created_at AS DATE) ASC").
		Preload("SentUser").
		Take(like)
	if err := result.Error; err != nil {
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

func (r *likeRepository) Begin() *gorm.DB {
	return r.db.Begin()
}

func (r *likeRepository) Consent(tx *gorm.DB, like *domain.Like) error {
	query := `UPDATE likes SET consented = ?
	WHERE recieved_user_uid = ? AND sent_user_uid = ?`
	if err := tx.Exec(query, like.Consented, like.RecievedUserUid, like.SentUserUid).Error; err != nil {
		return err
	}
	return nil
}
