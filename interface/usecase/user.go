package usecase

import (
	"context"
	"finder/domain"
)

// UserUsecase interface
type UserUsecase interface {
	GetUserByID(ctx context.Context, userID int64) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, *domain.User, error)
	DeleteUser(ctx context.Context, user *domain.User) (bool, error)
}
