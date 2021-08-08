package repository

import (
	"github.com/finder-app/finder-backend/grpc/domain"

	"github.com/jinzhu/gorm"
)

type LikeRepository interface {
	CreateLike(like *domain.Like) (*domain.Like, error)
	Liked(user *domain.User, visitorUid string) error
	GetOldestLikeByUid(currentUserUid string) (*domain.Like, error)
	Skip(like *domain.Like) error
	Begin() *gorm.DB
	GetLike(like *domain.Like) error
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

func (r *likeRepository) Liked(user *domain.User, visitorUid string) error {
	query := `SELECT EXISTS (
							SELECT *
							FROM likes
							WHERE recieved_user_uid = ?
							AND sent_user_uid = ?
						) AS 'liked'
						FROM likes`
	row := r.db.Raw(query, user.Uid, visitorUid).Row()
	if err := row.Err(); err != nil {
		return err
	}
	row.Scan(&user.Liked)
	return nil
}

func (r *likeRepository) GetOldestLikeByUid(currentUserUid string) (*domain.Like, error) {
	like := &domain.Like{}
	result := r.db.Model(like).
		Where(`recieved_user_uid = ? AND skipped = 0 AND consented = 0`, currentUserUid).
		Order("CAST(created_at AS DATE) ASC").
		Preload("SentUser").
		Take(like)
	if err := result.Error; err != nil {
		return nil, err
	}
	return like, nil
}

func (r *likeRepository) Skip(like *domain.Like) error {
	query := `UPDATE likes SET skipped = true
		WHERE sent_user_uid = ? AND recieved_user_uid = ?`
	if err := r.db.Exec(query, like.SentUserUid, like.RecievedUserUid).Error; err != nil {
		return err
	}
	return nil
}

func (r *likeRepository) Begin() *gorm.DB {
	return r.db.Begin()
}

func (r *likeRepository) GetLike(like *domain.Like) error {
	query := `SELECT *
						FROM likes
						WHERE recieved_user_uid = ?
						AND sent_user_uid = ?`
	// HACK: Raw.Scanで値が取れる理由が確定できていない。Rawで1行全体が取得できるから？
	row := r.db.Raw(query, like.RecievedUserUid, like.SentUserUid)
	if err := row.Error; err != nil {
		return err
	}
	if err := row.Scan(like).Error; err != nil {
		return err
	}
	return nil
}

func (r *likeRepository) Consent(tx *gorm.DB, like *domain.Like) error {
	query := `UPDATE likes SET consented = true
	WHERE recieved_user_uid = ? AND sent_user_uid = ?`
	if err := tx.Exec(query, like.RecievedUserUid, like.SentUserUid).Error; err != nil {
		return err
	}
	return nil
}
