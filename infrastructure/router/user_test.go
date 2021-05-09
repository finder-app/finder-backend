package router

import (
	"encoding/json"
	"finder/domain"
	"finder/domain/mocks"
	"finder/interface/controller"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	mockUseCase := new(mocks.UserUsecase)
	userController := controller.NewUserController(mockUseCase)

	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	router := NewRouter()
	router.Users(userController)

	ctx.Request, _ = http.NewRequest("GET", "/users", nil)
	idToken := os.Getenv("IDTOKEN")
	ctx.Request.Header.Set("Authorization", idToken)
	router.Engine.ServeHTTP(response, ctx.Request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestShow(t *testing.T) {
	var mockUser domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	mockUseCase := new(mocks.UserUsecase)
	userController := controller.NewUserController(mockUseCase)

	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	router := NewRouter()
	router.Users(userController)

	ctx.Request, _ = http.NewRequest("GET", "/users/"+mockUser.Uid, nil)
	idToken := os.Getenv("IDTOKEN")
	ctx.Request.Header.Set("Authorization", idToken)
	router.Engine.ServeHTTP(response, ctx.Request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestCreate(t *testing.T) {
	mockUseCase := new(mocks.UserUsecase)
	userController := controller.NewUserController(mockUseCase)

	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	router := NewRouter()
	router.Users(userController)

	// create User
	mockUser := domain.User{
		Uid:       "Uid",
		Email:     "ohishikaito@gmail.com",
		LastName:  "大石",
		FirstName: "海渡",
		IsMale:    true,
	}
	json, err := json.Marshal(mockUser)
	assert.NoError(t, err)
	body := strings.NewReader(string(json))

	ctx.Request, _ = http.NewRequest("POST", "/users", body)
	idToken := os.Getenv("IDTOKEN")
	ctx.Request.Header.Set("Authorization", idToken)
	router.Engine.ServeHTTP(response, ctx.Request)
	assert.Equal(t, http.StatusCreated, response.Code)
}
