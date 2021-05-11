package repository_test

import (
	"finder/domain"
	"finder/interface/repository"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func NewGormConnectMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	db, err := gorm.Open("mysql", sqlDB)
	if err != nil {
		return nil, nil, err
	}
	return db, mock, nil
}

func setMockUsers() []domain.User {
	mockUsers := []domain.User{
		{
			Uid:       "Uid",
			Email:     "ohishikaito@gmail.com",
			LastName:  "大石",
			FirstName: "海渡",
			IsMale:    true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Uid:       "Uid2",
			Email:     "ohishikaito2@gmail.com",
			LastName:  "きじま",
			FirstName: "あすか",
			IsMale:    false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return mockUsers
}

func TestGetUsersByGender(t *testing.T) {
	db, mock, err := NewGormConnectMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockUsers := setMockUsers()

	maleRows := sqlmock.NewRows([]string{
		"uid",
		"email",
		"last_name",
		"first_name",
		"is_male",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		mockUsers[0].Uid,
		mockUsers[0].Email,
		mockUsers[0].LastName,
		mockUsers[0].FirstName,
		mockUsers[0].IsMale,
		mockUsers[0].CreatedAt,
		mockUsers[0].UpdatedAt,
		mockUsers[0].DeletedAt,
	)
	femaleRows := sqlmock.NewRows([]string{
		"uid",
		"email",
		"last_name",
		"first_name",
		"is_male",
		"created_at",
		"updated_at",
	}).AddRow(
		mockUsers[1].Uid,
		mockUsers[1].Email,
		mockUsers[1].LastName,
		mockUsers[1].FirstName,
		mockUsers[1].IsMale,
		mockUsers[1].CreatedAt,
		mockUsers[1].UpdatedAt,
	)

	male := true
	female := false
	query := regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((is_male = ?))")
	mock.ExpectQuery(query).WithArgs(male).WillReturnRows(maleRows)
	mock.ExpectQuery(query).WithArgs(female).WillReturnRows(femaleRows)

	validate := validator.New()
	userRepository := repository.NewUserRepository(db, validate)

	maleUsers, err := userRepository.GetUsersByGender(male)
	assert.NoError(t, err)
	assert.Len(t, maleUsers, 1)

	femaleUsers, err := userRepository.GetUsersByGender(female)
	assert.NoError(t, err)
	assert.Len(t, femaleUsers, 1)
}

func TestGetUserByUid(t *testing.T) {
	db, mock, err := NewGormConnectMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockUsers := setMockUsers()
	rows := sqlmock.NewRows([]string{
		"uid",
		"email",
		"last_name",
		"first_name",
		"is_male",
		"created_at",
		"updated_at",
	}).AddRow(
		mockUsers[0].Uid,
		mockUsers[0].Email,
		mockUsers[0].LastName,
		mockUsers[0].FirstName,
		mockUsers[0].IsMale,
		mockUsers[0].CreatedAt,
		mockUsers[0].UpdatedAt,
	)
	query := regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((uid = ?)) LIMIT 1")
	uid := mockUsers[0].Uid
	mock.ExpectQuery(query).WithArgs(uid).WillReturnRows(rows)

	validate := validator.New()
	userRepository := repository.NewUserRepository(db, validate)
	user, err := userRepository.GetUserByUid(uid)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.Uid, mockUsers[0].Uid)
}