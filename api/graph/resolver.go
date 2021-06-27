package graph

import (
	"finder/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userUsecase usecase.UserUsecase
}

func NewResolver(
	userUsecase usecase.UserUsecase,
) *Resolver {
	return &Resolver{
		userUsecase,
	}
}