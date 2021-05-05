package repository

import (
	"finder/domain"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type FootPrintRepository interface {
	GetFootPrintsByUid(uid string) ([]domain.FootPrint, error)
	CreateFootPrint(uid string, visitorUid string) error
	UpdateFootPrint(footPrints *[]domain.FootPrint) error
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

func (r *footPrintRepository) GetFootPrintsByUid(visitorUid string) ([]domain.FootPrint, error) {
	footPrints := []domain.FootPrint{}
	if err := r.db.Model(domain.FootPrint{}).Where("visitor_uid = ?", visitorUid).Preload("User").Find(&footPrints).Error; err != nil {
		return nil, err
	}
	return footPrints, nil
}

func (r *footPrintRepository) CreateFootPrint(uid string, visitorUid string) error {
	footPrint := &domain.FootPrint{
		VisitorUid: visitorUid,
		UserUid:    uid,
		// FIX: UnreadはMySQL側でdefault '1'にしてるけど、何故か'0'でレコードが作られるため。要検証
		Unread: true,
	}
	// EXISTSで判別したかったけど、gormで再現できないから諦め。レコードあってもboolがfalseになる
	// var result bool
	// r.db.Raw(`
	// SELECT EXISTS (
	// 	SELECT * FROM foot_prints
	// 	WHERE visitor_uid = ?	AND user_uid = ?
	// )`, footPrint.VisitorUid, footPrint.UserUid).Scan(&result)
	// fmt.Println(result)

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

func (r *footPrintRepository) UpdateFootPrint(footPrints *[]domain.FootPrint) error {
	tx := r.db.Begin()
	for _, footPrint := range *footPrints {
		footPrint.Unread = false
		if err := tx.Table("foot_prints").Update(&footPrint).Error; err != nil {
			fmt.Println("UpdateFootPrint error")
			fmt.Println(err)
			tx.Rollback()
			return err
		}
		fmt.Println(footPrint.Unread)
	}
	tx.Commit()
	return nil
}
