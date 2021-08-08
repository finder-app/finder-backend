package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	infrastructure "github.com/finder-app/finder-backend/api/infrastructure"
	"github.com/finder-app/finder-backend/api/interface/controller"
	mocks "github.com/finder-app/finder-backend/api/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"gopkg.in/go-playground/assert.v1"
)

func newRouter() *infrastructure.Router {
	router := &infrastructure.Router{
		Engine: gin.Default(),
	}
	router.Engine.Use(setMockUserUid())
	return router
}

func setMockUserUid() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("currentUserUid", "mock_user_uid")
		c.Next()
	}
}

func TestIndex(t *testing.T) {
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	ctx.Request, _ = http.NewRequest("GET", "/users", nil)

	mockUserClient := new(mocks.UserClient)
	// NOTE: context.Contextはmock.Anythingにする必要ありそう
	// FIXME: mockのmethod内で型を参照できないため、mock.Anythingで対応
	mockUserClient.On("GetUsers", mock.Anything, mock.Anything).
		Return(mock.Anything, nil).Once()
	// Return(mock.AnythingOfType("*pb.GetUsersRes"), nil).Once()

	router := newRouter()
	userController := controller.NewUserController(mockUserClient)
	router.Users(userController)
	router.Engine.ServeHTTP(response, ctx.Request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestIndexError(t *testing.T) {
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	ctx.Request, _ = http.NewRequest("GET", "/users", nil)

	mockUserClient := new(mocks.UserClient)
	err := errors.New("Unexpexted Error")
	mockUserClient.On("GetUsers", mock.Anything, mock.Anything).Return(nil, err).Once()

	router := newRouter()
	userController := controller.NewUserController(mockUserClient)
	router.Users(userController)
	router.Engine.ServeHTTP(response, ctx.Request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

// TODO: showとcreateのテストは今後実装する

// func TestShow(t *testing.T) {
// 	mockUseCase := new(mocks.UserUsecase)
// 	mockUser := setMockUsers(t)[0]
// 	mockUseCase.On("GetUserByUid", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockUser, nil).Once()

// 	response := httptest.NewRecorder()
// 	ctx, _ := gin.CreateTestContext(response)
// 	ctx.Request, _ = http.NewRequest("GET", "/users/"+mockUser.Uid, nil)

// 	router := getUsersRouter(mockUseCase)
// 	router.Engine.ServeHTTP(response, ctx.Request)
// 	assert.Equal(t, http.StatusOK, response.Code)
// }

// func TestShowError(t *testing.T) {
// 	mockUseCase := new(mocks.UserUsecase)
// 	err := errors.New("record not found")
// 	mockUseCase.On("GetUserByUid", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, err).Once()

// 	response := httptest.NewRecorder()
// 	ctx, _ := gin.CreateTestContext(response)
// 	ctx.Request, _ = http.NewRequest("GET", "/users/"+"nil", nil)

// 	router := getUsersRouter(mockUseCase)
// 	router.Engine.ServeHTTP(response, ctx.Request)
// 	assert.Equal(t, http.StatusNotFound, response.Code)
// }

// func TestCreate(t *testing.T) {
// 	mockUseCase := new(mocks.UserUsecase)
// 	mockUser := domain.User{}
// 	faker.FakeData(&mockUser)
// 	mockUseCase.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(&mockUser, nil).Once()

// 	json, err := json.Marshal(mockUser)
// 	assert.NoError(t, err)
// 	body := strings.NewReader(string(json))

// 	response := httptest.NewRecorder()
// 	ctx, _ := gin.CreateTestContext(response)
// 	ctx.Request, _ = http.NewRequest("POST", "/users", body)

// 	router := getUsersRouter(mockUseCase)
// 	router.Engine.ServeHTTP(response, ctx.Request)
// 	assert.Equal(t, http.StatusCreated, response.Code)
// }

// func TestCreateError(t *testing.T) {
// 	mockUseCase := new(mocks.UserUsecase)
// 	err := errors.New("status unprocessable entity ")
// 	mockUseCase.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(nil, err).Once()

// 	mockUser := domain.User{}
// 	faker.FakeData(&mockUser)
// 	json, err := json.Marshal(mockUser)
// 	assert.NoError(t, err)
// 	body := strings.NewReader(string(json))

// 	response := httptest.NewRecorder()
// 	ctx, _ := gin.CreateTestContext(response)
// 	ctx.Request, _ = http.NewRequest("POST", "/users", body)

// 	router := getUsersRouter(mockUseCase)
// 	router.Engine.ServeHTTP(response, ctx.Request)
// 	assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
// }
