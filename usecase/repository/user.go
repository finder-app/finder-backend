package repository

import (
	"finder/domain"
)

type UserRepository interface {
	GetUsers() ([]domain.User, error)
	GetUserByID(id int) (*domain.User, error)
}
