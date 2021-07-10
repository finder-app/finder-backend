package controller_test

// import (
// 	"encoding/json"
// 	"errors"
// 	"finder/domain"
// 	"finder/infrastructure"
// 	"finder/interface/controller"
// 	"finder/usecase/mocks"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/bxcodec/faker"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // FIXME: ここmockUsecaseを引数に渡すの微妙な気がする
// func getUsersRouter(mockUseCase *mocks.UserUsecase) *infrastructure.Router {
// 	userController := controller.NewUserController(mockUseCase)
// 	router := &infrastructure.Router{
// 		Engine: gin.Default(),
// 	}
// 	router.Engine.Use(setCurrentUserUid())
// 	router.Users(userController)
// 	return router
// }

// func setCurrentUserUid() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		mockUser := domain.User{}
// 		faker.FakeData(&mockUser)
// 		currentUserUid := mockUser.Uid
// 		c.Set("currentUserUid", currentUserUid)
// 		c.Next()
// 	}
// }

// func setMockUsers(t *testing.T) []*domain.User {
// 	mockUser := &domain.User{}
// 	err := faker.FakeData(&mockUser)
// 	assert.NoError(t, err)
// 	mockUsers := []*domain.User{mockUser}
// 	return mockUsers
// }

// func TestIndex(t *testing.T) {
// 	response := httptest.NewRecorder()
// 	ctx, _ := gin.CreateTestContext(response)
// 	ctx.Request, _ = http.NewRequest("GET", "/users", nil)

// 	mockUseCase := new(mocks.UserUsecase)
// 	mockUsers := setMockUsers(t)
// 	mockUseCase.On("GetUsersByUid", mock.AnythingOfType("string")).Return(mockUsers, nil).Once()

// 	router := getUsersRouter(mockUseCase)
// 	router.Engine.ServeHTTP(response, ctx.Request)
// 	assert.Equal(t, http.StatusOK, response.Code)
// }

// func TestIndexError(t *testing.T) {
// 	response := httptest.NewRecorder()
// 	ctx, _ := gin.CreateTestContext(response)
// 	ctx.Request, _ = http.NewRequest("GET", "/users", nil)

// 	mockUseCase := new(mocks.UserUsecase)
// 	err := errors.New("Unexpexted Error")
// 	mockUseCase.On("GetUsersByUid", mock.AnythingOfType("string")).Return(nil, err).Once()

// 	router := getUsersRouter(mockUseCase)
// 	router.Engine.ServeHTTP(response, ctx.Request)
// 	assert.Equal(t, http.StatusInternalServerError, response.Code)
// }

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
