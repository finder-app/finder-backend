package repository

import (
	"finder/domain"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUsers() ([]domain.User, error)
	GetUserByID(id int) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
}

type userRepository struct {
	db       *gorm.DB
	validate *validator.Validate
}

func NewUserRepository(db *gorm.DB, validate *validator.Validate) *userRepository {
	return &userRepository{
		db:       db,
		validate: validate,
	}
}

func (r *userRepository) GetUsers() ([]domain.User, error) {
	users := []domain.User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByID(id int) (*domain.User, error) {
	user := domain.User{}
	if err := r.db.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	if err := r.validate.Struct(user); err != nil {
		return nil, err
	}
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
