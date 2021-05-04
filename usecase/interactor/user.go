package interactor

import (
	"finder/domain"
	"finder/interface/repository"
	"fmt"
)

type UserInteractor interface {
	GetUsers(uid string) ([]domain.User, error)
	GetUserByID(uid string) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
}

type userInteractor struct {
	userRepository repository.UserRepository
}

func NewUserInteractor(ur repository.UserRepository) *userInteractor {
	return &userInteractor{
		userRepository: ur,
	}
}

func (i *userInteractor) GetUsers(uid string) ([]domain.User, error) {
	user, _ := i.userRepository.GetUserByID(uid)
	fmt.Printf("検索するユーザー： %v\n", user.Email)
	fmt.Printf("検索するユーザーのis_male： %v\n", user.IsMale)
	var genderToSearch bool
	if user.IsMale == true {
		genderToSearch = false
	} else {
		genderToSearch = true
	}
	users, err := i.userRepository.GetUsers(genderToSearch)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (i *userInteractor) GetUserByID(uid string) (*domain.User, error) {
	user, err := i.userRepository.GetUserByID(uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (i *userInteractor) CreateUser(user *domain.User) (*domain.User, error) {
	user, err := i.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
