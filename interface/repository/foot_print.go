package repository

import (
	"finder/domain"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type FootPrintRepository interface {
	GetFootPrintUsersByUid(uid string) ([]domain.User, error)
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

func (r *footPrintRepository) GetFootPrintUsersByUid(uid string) ([]domain.User, error) {
	// 足跡を残したユーザー一覧を取得する！
	users := []domain.User{}
	query := `SELECT u.* FROM users AS u
		INNER JOIN foot_prints AS fp ON (fp.user_uid = u.uid)
		WHERE fp.visitor_uid = ?`
	if err := r.db.Raw(query, uid).Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
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
