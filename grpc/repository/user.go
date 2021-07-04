package repository

import (
	"grpc/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUserByUid(uid string) (*domain.User, error)
	GetUsersByGender(genderToSearch string) ([]*domain.User, error)
	GetUserByVisitorUid(visitorUid string) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
	UpdateUser(inputUser *domain.User) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByUid(uid string) (*domain.User, error) {
	return r.getUserByUid(uid)
}

// NOTE: testを通すために、GetUserByUidと全く同じメソッドを作成する。
func (r *userRepository) GetUserByVisitorUid(visitorUid string) (*domain.User, error) {
	return r.getUserByUid(visitorUid)
}

func (r *userRepository) getUserByUid(uid string) (*domain.User, error) {
	user := &domain.User{}
	if err := r.db.Where("uid = ?", uid).Take(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUsersByGender(genderToSearch string) ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.Where("gender = ?", genderToSearch).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(inputUser *domain.User) (*domain.User, error) {
	result := r.db.Model(domain.User{}).Where("uid = ?", inputUser.Uid).Update(inputUser)
	if err := result.Error; err != nil {
		return nil, err
	}
	return inputUser, nil
}
