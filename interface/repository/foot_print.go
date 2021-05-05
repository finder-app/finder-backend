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
	// Preloadつかうとエラー吐く　モデル指定してるのにー
	if err := r.db.Model(domain.FootPrint{}).Where("user_uid = ?", currentUserUid).Preload("User").Find(&footPrints).Error; err != nil {
		return nil, err
	}

	// resultをstructで作ってscanして、後からfoot_printsに追加するw
	// query := `SELECT foot_prints.created_at, users.* FROM foot_prints
	// INNER JOIN users on users.uid = foot_prints.visitor_uid
	// WHERE foot_prints.user_uid = ?`
	// rows, err := r.db.Raw(query, currentUserUid).Rows()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	footPrint := domain.FootPrint{}
	// 	rows.Scan(
	// 		&footPrint.CreatedAt,
	// 		&footPrint.User.Uid,
	// 		&footPrint.User.LastName,
	// 		&footPrint.User.FirstName,
	// 	)
	// 	footPrints = append(footPrints, footPrint)
	// }

	// if err := r.db.Raw(query, currentUserUid).Scan(&footPrints).Error; err != nil {
	// 	return nil, err
	// }
	return footPrints, nil
}

func (r *footPrintRepository) CreateFootPrint(currentUserUid string, visitorUid string) error {
	footPrint := &domain.FootPrint{
		VisitorUid: visitorUid,
		UserUid:    currentUserUid,
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

func (r *footPrintRepository) UpdateToAlreadyRead(currentUserUid string) error {
	// 後でトランザクションする
	// tx := r.db.Begin()
	var unreadFootPrintIds []int
	r.db.Raw(`SELECT id FROM foot_prints
		WHERE visitor_uid = ? AND unread = 1`, currentUserUid).Scan(&unreadFootPrintIds)
	fmt.Println(unreadFootPrintIds)

	// footPrint.Unread = false
	// if err := r.db.Model(domain.FootPrint{}).Where("id = ?", footPrint.ID).Update("").Error; err != nil {
	// 	fmt.Println("UpdateToAlreadyRead error")
	// 	fmt.Println(err)
	// 	// r.db.Rollback()
	// 	return err
	// }
	// fmt.Printf("footPrintのUnreadは%v", footPrint.Unread)

	// tx.Commit()
	// TODO: このアクションに現在の未読数を返すobjectを追加したい。
	// headerで表示してるはずなので、それを更新できるように。state.UnreadFootPrintCountを変える？
	return nil
}
