package interactor

import "finder/usecase/repository"

type userInteractor struct {
	userRepository repository.UserRepository
}

func NewUserInteractor(ur repository.UserRepository) *userInteractor {
	return &userInteractor{ur}
}
