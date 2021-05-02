package repository

import (
	"finder/domain"

	"github.com/jinzhu/gorm"
)

type (
	userRepository struct {
		DB *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(id int) (*domain.User, error) {
	user := domain.User{}
	if err := r.DB.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
