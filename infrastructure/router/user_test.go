package router

import (
	"finder/domain/mocks"
	"finder/interface/controller"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
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
	assert.Equal(t, 200, response.Code)
}

func TestGetUserById(t *testing.T) {
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
	assert.Equal(t, 200, response.Code)
}
