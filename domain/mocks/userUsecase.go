package mocks

import (
	"finder/domain"

	"github.com/stretchr/testify/mock"
)

type UserUsecase struct {
	mock.Mock
}

func (_m *UserUsecase) GetUsersByUid(uid string) ([]domain.User, error) {
	users := []domain.User{}
	return users, nil
}

func (_m *UserUsecase) GetUserByUid(uid string, visitorUid string) (*domain.User, error) {
	user := &domain.User{}
	return user, nil
}

func (_m *UserUsecase) CreateUser(user *domain.User) (*domain.User, error) {
	return user, nil
}
