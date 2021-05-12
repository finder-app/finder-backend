package controller_test

import (
	"encoding/json"
	"finder/domain"
	"finder/infrastructure/router"
	"finder/interface/controller"
	"finder/usecase/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func getUsersRouter(t *testing.T) *router.Router {
	mockUseCase := new(mocks.UserUsecase)
	userController := controller.NewUserController(mockUseCase)
	router := &router.Router{
		Engine: gin.Default(),
	}
	router.Engine.Use(setCurrentUserUid(t))
	router.Users(userController)
	return router
}

func setCurrentUserUid(t *testing.T) gin.HandlerFunc {
	return func(c *gin.Context) {
		var mockUser domain.User
		err := faker.FakeData(&mockUser)
		assert.NoError(t, err)
		currentUserUid := mockUser.Uid
		c.Set("currentUserUid", currentUserUid)
		c.Next()
	}
}

func TestIndex(t *testing.T) {
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	ctx.Request, _ = http.NewRequest("GET", "/users", nil)

	router := getUsersRouter(t)
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

	router := getUsersRouter(t)
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

	router := getUsersRouter(t)
	router.Engine.ServeHTTP(response, ctx.Request)
	assert.Equal(t, http.StatusCreated, response.Code)
}
