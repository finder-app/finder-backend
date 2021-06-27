package mocks

import (
	"finder/domain"
	"finder/shared"

	"github.com/stretchr/testify/mock"
)

type UserUsecase struct {
	mock.Mock
}

func (u *UserUsecase) GetUsersByUid(uid string) ([]*domain.User, error) {
	// TODO: このCalledで引数uidがいらない理由を調べる
	// arguments := u.Called(uid)
	arguments := u.Called()

	var users []*domain.User
	// NOTE: お手本のgo_clean_archに書いてあるけど呼び出されないのでコメントアウト
	// if rf, ok := arguments.Get(0).(func(uid string) []domain.User); ok {
	// 	users = rf(uid)
	// }
	if arguments.Get(0) != nil {
		users = arguments.Get(0).([]*domain.User)
	}
	err := shared.MockArgumentsError(arguments, 1)
	return users, err
}

func (u *UserUsecase) GetUserByUid(uid string, visitorUid string) (*domain.User, error) {
	arguments := u.Called(uid, visitorUid)

	user := &domain.User{}
	if arguments.Get(0) != nil {
		user = arguments.Get(0).(*domain.User)
	}
	err := shared.MockArgumentsError(arguments, 1)
	return user, err
}

func (u *UserUsecase) CreateUser(user *domain.User) (*domain.User, error) {
	arguments := u.Called(user)

	if arguments.Get(0) != nil {
		user = arguments.Get(0).(*domain.User)
	}
	err := shared.MockArgumentsError(arguments, 1)
	return user, err
}
