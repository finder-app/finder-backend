package interactor

import (
	"finder/domain"
)

// UserUsecase interface
type UserInteractor interface {
	GetUserByID(userID int) (*domain.User, error)
}
