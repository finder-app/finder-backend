package repository

import (
	"finder/domain"
	"finder/infrastructure/validations"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUsersByGender(genderToSearch string) ([]domain.User, error)
	GetUserByUid(uid string) (domain.User, error)
	GetUserByVisitorUid(visitorUid string) (domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUsersByGender(genderToSearch string) ([]domain.User, error) {
	users := []domain.User{}
	if err := r.db.Where("gender = ?", genderToSearch).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByUid(uid string) (domain.User, error) {
	return r.getUserByUid(uid)
}

// NOTE: testを通すために、GetUserByUidと全く同じメソッドを作成する。
func (r *userRepository) GetUserByVisitorUid(visitorUid string) (domain.User, error) {
	return r.getUserByUid(visitorUid)
}

func (r *userRepository) getUserByUid(uid string) (domain.User, error) {
	user := domain.User{}
	if err := r.db.Where("uid = ?", uid).Take(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	if err := validations.ValidateUser(user); err != nil {
		return nil, err
	}
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	if err := validations.ValidateUser(user); err != nil {
		return nil, err
	}
	result := r.db.Model(domain.User{}).Where("uid = ?", user.Uid).Update(
		user.LastName,
		user.FirstName,
	)
	if err := result.Error; err != nil {
		return nil, err
	}
	return user, nil
}
