package repository

import (
	"context"
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

func (r *userRepository) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	user := domain.User{}
	if err := r.DB.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
