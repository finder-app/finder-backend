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
	// result := r.db.Model(domain.FootPrint{}).Where("user_uid = ?", currentUserUid).Preload("User").Find(&footPrints)

	// とれるw foreign_keyはuser_uid。逆の組み合わせってこと
	// result := r.db.Model(domain.FootPrint{}).Where("visitor_uid = ?", currentUserUid).Preload("User").Find(&footPrints)
	// result := r.db.Model(domain.FootPrint{}).Where("user_uid = ?", currentUserUid).Preload("User").Find(&footPrints)

	// 現状の回答
	// result := r.db.Model(domain.FootPrint{}).Where("user_uid = ?", currentUserUid).Find(&footPrints)
	result := r.db.Model(domain.FootPrint{}).Where("user_uid = ?", currentUserUid).Preload("Visitor").Find(&footPrints)
	if err := result.Error; err != nil {
		return nil, err
	}
	// if err := result.Joins("users").Preload("User").Find(&footPrints).Error; err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }

	// 別解を試す。footPrintだけを先に取得。次にユーザーを取得して、
	// footPrintsのUserに突っ込む。取れない気がしてきた
	// if err := r.db.Model(domain.FootPrint{}).Where("user_uid = ?", currentUserUid).Preload("User").Find(&footPrints).Error; err != nil {
	// 	return nil, err
	// }

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
	// result := r.db.Exec("UPDATE foot_prints SET unread = 1 WHERE unread = 0 AND user_uid = ?", currentUserUid)
	if err := result.Error; err != nil {
		return err
	}
	// TODO: このアクションに現在の未読数を返すobjectを追加したい。
	// headerで表示してるはずなので、それを更新できるように。state.UnreadFootPrintCountを変える？仕様は検討する
	return nil
}
