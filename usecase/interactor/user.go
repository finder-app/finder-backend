package interactor

import (
	"finder/domain"
	"finder/interface/repository"
	"fmt"
)

type UserInteractor interface {
	GetUsersByUid(uid string) ([]domain.User, error)
	GetUserByUid(uid string) (*domain.User, error)
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

func (i *userInteractor) GetUsersByUid(uid string) ([]domain.User, error) {
	user, _ := i.userRepository.GetUserByUid(uid)
	fmt.Printf("検索するユーザー： %v\n", user.Email)
	fmt.Printf("検索するユーザーのis_male： %v\n", user.IsMale)
	gender := getGenderForSearch(user.IsMale)
	users, err := i.userRepository.GetUsersByGender(gender)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (i *userInteractor) GetUserByUid(uid string) (*domain.User, error) {
	user, err := i.userRepository.GetUserByUid(uid)
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

func getGenderForSearch(isMale bool) bool {
	if isMale {
		return false
	} else {
		return true
	}
}
