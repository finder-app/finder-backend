package mocks

import (
	"finder/domain"
	"finder/shared"

	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (r *UserRepository) GetUsersByGender(genderToSearch string) ([]domain.User, error) {
	arguments := r.Called(genderToSearch)

	users := []domain.User{}
	if arguments.Get(0) != nil {
		users = arguments.Get(0).([]domain.User)
	}
	err := shared.MockArgumentsError(arguments, 1)
	return users, err
}

func (r *UserRepository) GetUserByUid(uid string) (domain.User, error) {
	arguments := r.Called(uid)

	user := domain.User{}
	if arguments.Get(0) != nil {
		user = arguments.Get(0).(domain.User)
	}
	err := shared.MockArgumentsError(arguments, 1)
	return user, err
}

func (r *UserRepository) GetUserByVisitorUid(visitorUid string) (domain.User, error) {
	arguments := r.Called(visitorUid)

	user := domain.User{}
	if arguments.Get(0) != nil {
		user = arguments.Get(0).(domain.User)
	}
	err := shared.MockArgumentsError(arguments, 1)
	return user, err
}

func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	arguments := r.Called(user)

	if arguments.Get(0) != nil {
		user = arguments.Get(0).(*domain.User)
	}
	err := shared.MockArgumentsError(arguments, 1)
	return user, err
}

func (r *UserRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	arguments := r.Called(user)

	if arguments.Get(0) != nil {
		user = arguments.Get(0).(*domain.User)
	}
	err := shared.MockArgumentsError(arguments, 1)
	return user, err
}
