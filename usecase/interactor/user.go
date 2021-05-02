package interactor

import (
	"finder/domain"
	"finder/usecase/repository"
)

type userInteractor struct {
	userRepository repository.UserRepository
}

func NewUserInteractor(ur repository.UserRepository) *userInteractor {
	return &userInteractor{ur}
}

func (i *userInteractor) GetUsers() ([]domain.User, error) {
	result, err := i.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i *userInteractor) GetUserByID(userID int) (*domain.User, error) {
	result, err := i.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
