package interactor

import (
	"finder/domain"
)

type UserInteractor interface {
	GetUserByID(userID int) (*domain.User, error)
}
