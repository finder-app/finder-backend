package router_test

import (
	"encoding/json"
	"finder/domain"
	"finder/domain/mocks"
	"finder/infrastructure/router"
	"finder/interface/controller"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func getUsersRouter() *router.Router {
	mockUseCase := new(mocks.UserUsecase)
	userController := controller.NewUserController(mockUseCase)
	router := router.NewRouter()
	router.Users(userController)
	return router
}

func setCurrentUserUid(req *http.Request, t *testing.T) {
	var mockUser domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	currentUserUid := mockUser.Uid
	req.Header.Set("currentUserUid", currentUserUid)
}

func TestIndex(t *testing.T) {
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)

	ctx.Request, _ = http.NewRequest("GET", "/users", nil)
	setCurrentUserUid(ctx.Request, t)

	router := getUsersRouter()
	router.Engine.ServeHTTP(response, ctx.Request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestShow(t *testing.T) {
	var mockUser domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	ctx.Request, _ = http.NewRequest("GET", "/users/"+mockUser.Uid, nil)
	setCurrentUserUid(ctx.Request, t)

	router := getUsersRouter()
	router.Engine.ServeHTTP(response, ctx.Request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestCreate(t *testing.T) {
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

	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)

	ctx.Request, _ = http.NewRequest("POST", "/users", body)
	setCurrentUserUid(ctx.Request, t)

	router := getUsersRouter()
	router.Engine.ServeHTTP(response, ctx.Request)
	assert.Equal(t, http.StatusCreated, response.Code)
}
