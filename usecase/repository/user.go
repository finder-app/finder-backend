package repository

import (
	"context"
	"finder/domain"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
}
