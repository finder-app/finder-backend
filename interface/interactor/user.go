package interactor

import (
	"finder/domain"
)

type UserInteractor interface {
	GetUsers() ([]domain.User, error)
	GetUserByID(userID int) (*domain.User, error)
}
