package repository

import (
	"finder/domain"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type FootPrintRepository interface {
	GetFootPrintUsersByUid(currentUserUid string) ([]domain.FootPrint, error)
	CreateFootPrint(currentUserUid string, visitorUid string) error
	UpdateToAlreadyRead(currentUserUid string) error
}

type footPrintRepository struct {
	db       *gorm.DB
	validate *validator.Validate
}

func NewFootPrintRepository(db *gorm.DB, validate *validator.Validate) *footPrintRepository {
	return &footPrintRepository{
		db:       db,
		validate: validate,
	}
}

// 足跡をつけた時間が欲しい！
func (r *footPrintRepository) GetFootPrintUsersByUid(currentUserUid string) ([]domain.FootPrint, error) {
	footPrints := []domain.FootPrint{}
	result := r.db.Model(domain.FootPrint{}).Where("user_uid = ?", currentUserUid).Preload("Visitor").Find(&footPrints)
	if err := result.Error; err != nil {
		return nil, err
	}
	return footPrints, nil
}

func (r *footPrintRepository) CreateFootPrint(currentUserUid string, visitorUid string) error {
	footPrint := &domain.FootPrint{
		VisitorUid: visitorUid,
		UserUid:    currentUserUid,
		Unread:     true,
	}
	where := domain.FootPrint{
		VisitorUid: footPrint.VisitorUid,
		UserUid:    footPrint.UserUid,
	}
	if err := r.db.FirstOrCreate(footPrint, where).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *footPrintRepository) UpdateToAlreadyRead(currentUserUid string) error {
	result := r.db.Exec("UPDATE foot_prints SET unread = 0 WHERE unread = 1 AND user_uid = ?", currentUserUid)
	if err := result.Error; err != nil {
		return err
	}
	// TODO: このアクションに現在の未読数を返すobjectを追加したい。
	// headerで表示してるはずなので、それを更新できるように。state.UnreadFootPrintCountを変える？仕様は検討する
	return nil
}
