package usecase_test

import (
	"errors"
	"finder/domain"
	"finder/infrastructure/repository/mocks"
	"finder/usecase"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setMockUsers(t *testing.T) (mockUsers []domain.User) {
	mockMaleUser := domain.User{}
	err := faker.FakeData(&mockMaleUser)
	assert.NoError(t, err)
	mockMaleUser.Gender = "男性"
	mockUsers = append(mockUsers, mockMaleUser)

	mockFemaleUser := domain.User{}
	err = faker.FakeData(&mockFemaleUser)
	assert.NoError(t, err)
	mockFemaleUser.Gender = "女性"
	mockUsers = append(mockUsers, mockFemaleUser)
	return
}

func TestGetUsersByUid(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	mockFootPrintRepository := new(mocks.FootPrintRepository)
	// foot_printsは仮で作った
	mockUsecase := usecase.NewUserUsecase(mockUserRepository, mockFootPrintRepository)
	mockUsers := setMockUsers(t)

	mockMaleUsers := []domain.User{}
	mockMaleUser := mockUsers[0]
	mockMaleUsers = append(mockMaleUsers, mockMaleUser)

	mockFemaleUsers := []domain.User{}
	mockFemaleUser := mockUsers[1]
	mockFemaleUsers = append(mockFemaleUsers, mockFemaleUser)

	t.Run("男性が女性の一覧を取得", func(t *testing.T) {
		mockUserRepository.On("GetUserByUid", mock.AnythingOfType("string")).Return(mockMaleUser, nil).Once()
		mockUserRepository.On("GetUsersByGender", mock.AnythingOfType("string")).Return(mockFemaleUsers, nil).Once()
		maleUsers, err := mockUsecase.GetUsersByUid(mockMaleUser.Uid)
		assert.NoError(t, err)
		assert.Len(t, maleUsers, len(mockMaleUsers))
	})
	t.Run("女性が男性の一覧を取得", func(t *testing.T) {
		mockUserRepository.On("GetUserByUid", mock.AnythingOfType("string")).Return(mockFemaleUser, nil).Once()
		mockUserRepository.On("GetUsersByGender", mock.AnythingOfType("string")).Return(mockMaleUsers, nil).Once()
		maleUsers, err := mockUsecase.GetUsersByUid(mockFemaleUser.Uid)
		assert.NoError(t, err)
		assert.Len(t, maleUsers, len(mockMaleUsers))
	})
	t.Run("異常値", func(t *testing.T) {
		mockUserRepository.On("GetUserByUid", mock.AnythingOfType("string")).Return(mockMaleUser, nil).Once()
		err := errors.New("Unexpexted Error")
		mockUserRepository.On("GetUsersByGender", mock.AnythingOfType("string")).Return(nil, err).Once()
		maleUsers, err := mockUsecase.GetUsersByUid(mockMaleUser.Uid)
		assert.Error(t, err)
		assert.Len(t, maleUsers, 0)
	})
}

// func TestGetUserByUid(t *testing.T) {
// 	// foot_prints作る！
// 	// response
// }

// func TestCreateUser(t *testing.T) {
// 	// 普通にcreateする
// }
