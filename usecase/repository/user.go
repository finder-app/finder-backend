package repository

import (
	"finder/domain"
)

type UserRepository interface {
	GetUserByID(id int) (*domain.User, error)
}
