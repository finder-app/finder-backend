package mocks

import (
	"finder/domain"

	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func argumentError(arguments mock.Arguments, index int) (err error) {
	if arguments.Error(index) != nil {
		err = arguments.Error(index)
	}
	return
}

func (u *UserRepository) GetUsersByGender(genderToSearch string) ([]domain.User, error) {
	arguments := u.Called(genderToSearch)

	users := []domain.User{}
	if arguments.Get(0) != nil {
		users = arguments.Get(0).([]domain.User)
	}
	err := argumentError(arguments, 1)
	return users, err
}

func (u *UserRepository) GetUserByUid(uid string) (domain.User, error) {
	arguments := u.Called(uid)

	user := domain.User{}
	if arguments.Get(0) != nil {
		user = arguments.Get(0).(domain.User)
	}
	err := argumentError(arguments, 1)
	return user, err
}

func (u *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	arguments := u.Called(user)

	if arguments.Get(0) != nil {
		user = arguments.Get(0).(*domain.User)
	}
	err := argumentError(arguments, 1)
	return user, err
}

func (u *UserRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	arguments := u.Called(user)

	if arguments.Get(0) != nil {
		user = arguments.Get(0).(*domain.User)
	}
	err := argumentError(arguments, 1)
	return user, err
}
