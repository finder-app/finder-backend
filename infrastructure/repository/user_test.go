package repository_test

import (
	"database/sql/driver"
	"finder/domain"
	"finder/infrastructure/repository"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func NewMockGormConnect(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db, err := gorm.Open("mysql", sqlDB)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func setMockUsers() (mockUsers []domain.User) {
	mockMaleUser := domain.User{}
	faker.FakeData(&mockMaleUser)
	mockMaleUser.Gender = "男性"
	mockUsers = append(mockUsers, mockMaleUser)

	femockMaleUser := domain.User{}
	faker.FakeData(&femockMaleUser)
	mockMaleUser.Gender = "女性"
	mockUsers = append(mockUsers, femockMaleUser)
	return
}

func getRows(mockUser domain.User) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{
		"uid",
		"email",
		"last_name",
		"first_name",
		"gender",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		mockUser.Uid,
		mockUser.Email,
		mockUser.LastName,
		mockUser.FirstName,
		mockUser.Gender,
		mockUser.CreatedAt,
		mockUser.UpdatedAt,
		mockUser.DeletedAt,
	)
	return rows
}

func TestGetUsersByGender(t *testing.T) {
	db, mock := NewMockGormConnect(t)
	mockUsers := setMockUsers()

	query := regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((gender = ?))")
	male, female := "男性", "女性"
	maleRows := getRows(mockUsers[0])
	femaleRows := getRows(mockUsers[1])
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
	db, mock := NewMockGormConnect(t)
	mockUser := setMockUsers()[0]

	query := regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((uid = ?)) LIMIT 1")
	uid := mockUser.Uid
	rows := getRows(mockUser)
	mock.ExpectQuery(query).WithArgs(uid).WillReturnRows(rows)

	validate := validator.New()
	userRepository := repository.NewUserRepository(db, validate)
	user, err := userRepository.GetUserByUid(uid)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.Uid, mockUser.Uid)
}

func TestCreateUser(t *testing.T) {
	db, mock := NewMockGormConnect(t)
	mockUser := setMockUsers()[0]

	mock.ExpectBegin()
	query := regexp.QuoteMeta("INSERT INTO `users` (`uid`,`email`,`last_name`,`first_name`,`gender`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?)")
	mock.ExpectExec(query).WithArgs(
		mockUser.Uid,
		mockUser.Email,
		mockUser.LastName,
		mockUser.FirstName,
		mockUser.Gender,
		mockUser.CreatedAt,
		mockUser.UpdatedAt,
		mockUser.DeletedAt,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	validate := validator.New()
	userRepository := repository.NewUserRepository(db, validate)
	user, err := userRepository.CreateUser(&mockUser)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

// NOTE: UpdateUserのテストで使用する
// https://github.com/DATA-DOG/go-sqlmock#matching-arguments-like-timetime
type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestUpdateUser(t *testing.T) {
	db, mock := NewMockGormConnect(t)
	mockUser := setMockUsers()[0]
	mockUser.LastName = "updateLastName"

	mock.ExpectBegin()
	query := regexp.QuoteMeta("UPDATE `users` SET `first_name` = ?, `last_name` = ?, `updated_at` = ? WHERE `users`.`deleted_at` IS NULL AND ((uid = ?))")

	mock.ExpectExec(query).WithArgs(
		mockUser.FirstName,
		mockUser.LastName,
		AnyTime{},
		mockUser.Uid,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	validate := validator.New()
	userRepository := repository.NewUserRepository(db, validate)
	user, err := userRepository.UpdateUser(&mockUser)
	assert.NoError(t, err)
	assert.Equal(t, mockUser.LastName, user.LastName)
}
