package repository

import (
	"finder/domain"

	"github.com/jinzhu/gorm"
)

type FootPrintRepository interface {
	GetFootPrintsByUid(currentUserUid string) ([]domain.FootPrint, error)
	CreateFootPrint(footPrint *domain.FootPrint) error
	UpdateToAlreadyRead(currentUserUid string) error
	GetUnreadCount(currentUserUid string) (int, error)
}

type footPrintRepository struct {
	db *gorm.DB
}

func NewFootPrintRepository(db *gorm.DB) *footPrintRepository {
	return &footPrintRepository{
		db: db,
	}
}

func (r *footPrintRepository) GetFootPrintsByUid(currentUserUid string) ([]domain.FootPrint, error) {
	footPrints := []domain.FootPrint{}
	result := r.db.Model(domain.FootPrint{}).Where("user_uid = ?", currentUserUid).Preload("Visitor").Find(&footPrints)
	if err := result.Error; err != nil {
		return nil, err
	}
	return footPrints, nil
}

func (r *footPrintRepository) CreateFootPrint(footPrint *domain.FootPrint) error {
	where := domain.FootPrint{
		VisitorUid: footPrint.VisitorUid,
		UserUid:    footPrint.UserUid,
	}
	if err := r.db.FirstOrCreate(footPrint, where).Error; err != nil {
		return err
	}
	return nil
}

func (r *footPrintRepository) UpdateToAlreadyRead(currentUserUid string) error {
	result := r.db.Exec("UPDATE foot_prints SET unread = 0 WHERE unread = 1 AND user_uid = ?", currentUserUid)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (r *footPrintRepository) GetUnreadCount(currentUserUid string) (int, error) {
	var unreadCount int
	query := `SELECT count(*) AS unreadCount FROM foot_prints WHERE user_uid = ? AND unread = 1`
	row := r.db.Raw(query, currentUserUid).Row()
	if err := row.Err(); err != nil {
		return 0, err
	}
	row.Scan(&unreadCount)
	return unreadCount, nil
}