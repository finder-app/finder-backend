package usecase

import (
	"finder/domain"
	"finder/interface/repository"
)

type ProfileUsecase interface {
	GetProfileByUid(currentUserUid string) (*domain.User, error)
	UpdateUser(currentUserUid string, updateUser *domain.UpdateUser) (*domain.User, error)
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
func (i *profileUsecase) UpdateUser(currentUserUid string, updateUser *domain.UpdateUser) (*domain.User, error) {
	// userに移し替え
	user := &domain.User{
		LastName:  updateUser.LastName,
		FirstName: updateUser.FirstName,
	}
	return i.userRepository.UpdateUser(currentUserUid, user)
}
