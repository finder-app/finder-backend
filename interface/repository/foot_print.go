package repository

import (
	"finder/domain"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type FootPrintRepository interface {
	CreateFootPrint(uid string, visitorUid string) error
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
