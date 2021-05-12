package usecase

import (
	"finder/domain"
	"finder/interface/repository"
)

type ProfileUsecase interface {
	GetProfileByUid(currentUserUid string) (*domain.User, error)
}

type profileUsecase struct {
	userRepository repository.UserRepository
}

func NewProfileUsecase(ur repository.UserRepository) *profileUsecase {
	return &profileUsecase{
		userRepository: ur,
	}
}

func (i *profileUsecase) GetProfileByUid(currentUserUid string) (*domain.User, error) {
	user, err := i.userRepository.GetUserByUid(currentUserUid)
	if err != nil {
		return nil, err
	}
	return user, nil
}
