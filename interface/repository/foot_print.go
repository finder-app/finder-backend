package repository

import (
	"finder/domain"

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
	}
	if err := r.db.Create(footPrint).Error; err != nil {
		return err
	}
	return nil
}
