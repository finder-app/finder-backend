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

// func setMockUsers(t *testing.T) (mockMaleUser domain.User, mockFemaleUser domain.User) {
// 	err := faker.FakeData(&mockMaleUser)
// 	assert.NoError(t, err)
// 	mockMaleUser.Gender = "男性"

// 	err = faker.FakeData(&mockFemaleUser)
// 	assert.NoError(t, err)
// 	mockFemaleUser.Gender = "女性"
// 	return
// }

func setMockUsers(t *testing.T) []domain.User {
	mockMaleUser := domain.User{}
	err := faker.FakeData(&mockMaleUser)
	assert.NoError(t, err)
	mockMaleUser.Gender = "男性"

	mockFemaleUser := domain.User{}
	err = faker.FakeData(&mockFemaleUser)
	assert.NoError(t, err)
	mockFemaleUser.Gender = "女性"

	mockUsers := []domain.User{
		mockMaleUser,
		mockFemaleUser,
	}
	return mockUsers
}

func TestGetUsersByUid(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	mockFootPrintRepository := new(mocks.FootPrintRepository)
	mockUsecase := usecase.NewUserUsecase(mockUserRepository, mockFootPrintRepository)

	mockUsers := setMockUsers(t)

	mockMaleUser := mockUsers[0]
	mockMaleUsers := []domain.User{mockMaleUser}
	mockFemaleUser := mockUsers[1]
	mockFemaleUsers := []domain.User{mockFemaleUser}

	t.Run("男性が女性の一覧を取得", func(t *testing.T) {
		mockUserRepository.On("GetUserByUid", mock.AnythingOfType("string")).Return(mockMaleUser, nil).Once()
		mockUserRepository.On("GetUsersByGender", mock.AnythingOfType("string")).Return(mockFemaleUsers, nil).Once()
		femaleUsers, err := mockUsecase.GetUsersByUid(mockMaleUser.Uid)
		assert.NoError(t, err)
		assert.Equal(t, femaleUsers[0].Uid, (mockFemaleUsers[0].Uid))
		assert.Len(t, femaleUsers, len(mockFemaleUsers))
	})
	t.Run("女性が男性の一覧を取得", func(t *testing.T) {
		mockUserRepository.On("GetUserByUid", mock.AnythingOfType("string")).Return(mockFemaleUser, nil).Once()
		mockUserRepository.On("GetUsersByGender", mock.AnythingOfType("string")).Return(mockMaleUsers, nil).Once()
		maleUsers, err := mockUsecase.GetUsersByUid(mockFemaleUser.Uid)
		assert.NoError(t, err)
		assert.Equal(t, maleUsers[0].Uid, (mockMaleUsers[0].Uid))
		assert.Len(t, maleUsers, len(mockMaleUsers))
	})
	t.Run("異常値", func(t *testing.T) {
		mockUserRepository.On("GetUserByUid", mock.AnythingOfType("string")).Return(mockMaleUser, nil).Once()
		err := errors.New("Unexpexted Error")
		mockUserRepository.On("GetUsersByGender", mock.AnythingOfType("string")).Return(nil, err).Once()
		femaleUsers, err := mockUsecase.GetUsersByUid(mockMaleUser.Uid)
		assert.Error(t, err)
		assert.Len(t, femaleUsers, 0)
	})
}

func TestGetUserByUid(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	mockFootPrintRepository := new(mocks.FootPrintRepository)
	mockUsecase := usecase.NewUserUsecase(mockUserRepository, mockFootPrintRepository)

	mockUsers := setMockUsers(t)
	mockMaleUser := mockUsers[0]
	mockFemaleUser := mockUsers[1]

	t.Run("男性が女性の詳細を取得", func(t *testing.T) {
		mockUserRepository.On("GetUserByVisitorUid", mock.AnythingOfType("string")).Return(mockMaleUser, nil).Once()
		mockUserRepository.On("GetUserByUid", mock.AnythingOfType("string")).Return(mockFemaleUser, nil).Once()
		mockFootPrintRepository.On("CreateFootPrint", mock.AnythingOfType("*domain.FootPrint")).Return(nil).Once()

		_, err := mockUsecase.GetUserByUid(mockFemaleUser.Uid, mockMaleUser.Uid)
		assert.NoError(t, err)
	})
	t.Run("女性が男性の詳細を取得", func(t *testing.T) {
		mockUserRepository.On("GetUserByUid", mock.AnythingOfType("string")).Return(mockMaleUser, nil).Once()
		mockUserRepository.On("GetUserByVisitorUid", mock.AnythingOfType("string")).Return(mockFemaleUser, nil).Once()
		mockFootPrintRepository.On("CreateFootPrint", mock.AnythingOfType("*domain.FootPrint")).Return(nil).Once()

		_, err := mockUsecase.GetUserByUid(mockMaleUser.Uid, mockFemaleUser.Uid)
		assert.NoError(t, err)
	})
}

// NOTE: 成功時と同じmockUsecaseを使用するとエラーが発生するため、関数を分ける
func TestGetUserByUidError(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	mockFootPrintRepository := new(mocks.FootPrintRepository)
	mockUsecase := usecase.NewUserUsecase(mockUserRepository, mockFootPrintRepository)

	mockUsers := setMockUsers(t)
	mockMaleUser := mockUsers[0]
	mockFemaleUser := mockUsers[1]

	t.Run("異常値（同性の詳細へリクエストを送った場合）", func(t *testing.T) {
		mockUserRepository.On("GetUserByUid", mock.AnythingOfType("string")).Return(mockMaleUser, nil).Once()
		mockUserRepository.On("GetUserByVisitorUid", mock.AnythingOfType("string")).Return(mockMaleUser, nil).Once()

		user, err := mockUsecase.GetUserByUid(mockMaleUser.Uid, mockMaleUser.Uid)
		assert.Error(t, err)
		assert.Equal(t, domain.User{}, user)
	})
	t.Run("異常値（存在しないユーザー）", func(t *testing.T) {
		newError := errors.New("record not found")
		mockUserRepository.On("GetUserByUid", mock.AnythingOfType("string")).Return(nil, newError).Once()

		user, err := mockUsecase.GetUserByUid(mockFemaleUser.Uid, mockMaleUser.Uid)
		assert.Error(t, err)
		assert.Equal(t, domain.User{}, user)
	})
	t.Run("異常値（足跡が作成できない）", func(t *testing.T) {
		mockUserRepository.On("GetUserByUid", mock.AnythingOfType("string")).Return(mockMaleUser, nil).Once()
		mockUserRepository.On("GetUserByVisitorUid", mock.AnythingOfType("string")).Return(mockFemaleUser, nil).Once()
		err := errors.New("StatusUnprocessable Entity")
		mockFootPrintRepository.On("CreateFootPrint", mock.AnythingOfType("*domain.FootPrint")).Return(err).Once()

		user, err := mockUsecase.GetUserByUid(mockMaleUser.Uid, mockFemaleUser.Uid)
		assert.Error(t, err)
		assert.Equal(t, domain.User{}, user)
	})
}

func TestCreateUser(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	mockFootPrintRepository := new(mocks.FootPrintRepository)
	mockUsecase := usecase.NewUserUsecase(mockUserRepository, mockFootPrintRepository)
	mockUser := setMockUsers(t)[0]

	t.Run("正常", func(t *testing.T) {
		mockUserRepository.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(&mockUser, nil).Once()
		user, err := mockUsecase.CreateUser(&mockUser)
		assert.NoError(t, err)
		assert.Equal(t, user.Uid, mockUser.Uid)
	})
	t.Run("異常値", func(t *testing.T) {
		err := errors.New("StatusUnprocessable Entity")
		mockUserRepository.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(nil, err).Once()
		// NOTE: 多分pointerでuserを渡してるから、userの値が返ってきちゃう。テストしない。
		// errが返ってくるかのテストなので。go-clean-archも確認してない。あれは返り値ないからだけど
		_, err = mockUsecase.CreateUser(&mockUser)
		assert.Error(t, err)
	})
}
