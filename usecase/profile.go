package usecase

import (
	"errors"
	"finder/domain"
	"finder/interface/repository"
)

type ProfileUsecase interface {
	GetProfileByUid(currentUserUid string) (*domain.User, error)
	UpdateUser(currentUserUid string, user *domain.User) (*domain.User, error)
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
func (i *profileUsecase) UpdateUser(currentUserUid string, user *domain.User) (*domain.User, error) {
	if user.Uid != currentUserUid {
		return nil, errors.New("illegal value")
	}
	return i.userRepository.UpdateUser(user)
}
