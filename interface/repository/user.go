package repository

import (
	"finder/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUsers() ([]domain.User, error)
	GetUserByID(id int) (*domain.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUsers() ([]domain.User, error) {
	users := []domain.User{}
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByID(id int) (*domain.User, error) {
	user := domain.User{}
	if err := r.DB.Find(&user, id).Error; err != nil {
		return nil, err
	}
	// ここ実態で返したらどうなる？
	return &user, nil
}
