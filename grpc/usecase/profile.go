package usecase

import (
	"github.com/finder-app/finder-backend/grpc/domain"
	"github.com/finder-app/finder-backend/grpc/repository"
)

type ProfileUsecase interface {
	GetProfile(currentUserUid string) (*domain.User, error)
	UpdateProfile(inputUser *domain.User) (*domain.User, error)
}

type profileUsecase struct {
	userRepository repository.UserRepository
}

func NewProfileUsecase(userRepository repository.UserRepository) ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
	}
}

func (u *profileUsecase) GetProfile(currentUserUid string) (*domain.User, error) {
	return u.userRepository.GetUserByUid(currentUserUid)
}

func (u *profileUsecase) UpdateProfile(inputUser *domain.User) (*domain.User, error) {
	// TODO: update前にvalidationしたい
	return u.userRepository.UpdateUser(inputUser)
}
