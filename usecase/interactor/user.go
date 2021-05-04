package interactor

import (
	"finder/domain"
	"finder/interface/repository"
)

type UserInteractor interface {
	GetUsers() ([]domain.User, error)
	GetUserByID(userID int) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
}

type userInteractor struct {
	userRepository repository.UserRepository
}

func NewUserInteractor(r repository.UserRepository) *userInteractor {
	return &userInteractor{
		userRepository: r,
	}
}

func (i *userInteractor) GetUsers() ([]domain.User, error) {
	users, err := i.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (i *userInteractor) GetUserByID(userID int) (*domain.User, error) {
	user, err := i.userRepository.GetUserByID(userID)
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
