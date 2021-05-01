package repository

import (
	"context"
	"finder/domain"
)

// UserRepository interface
type UserRepository interface {
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
}
