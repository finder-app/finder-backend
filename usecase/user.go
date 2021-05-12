package usecase

import (
	"finder/domain"
	"finder/interface/repository"
)

type UserUsecase interface {
	GetUsersByUid(uid string) ([]domain.User, error)
	GetUserByUid(uid string, visitorUid string) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
}

type userUsecase struct {
	userRepository      repository.UserRepository
	footPrintRepository repository.FootPrintRepository
}

func NewUserUsecase(ur repository.UserRepository, fpr repository.FootPrintRepository) *userUsecase {
	return &userUsecase{
		userRepository:      ur,
		footPrintRepository: fpr,
	}
}

func (i *userUsecase) GetUsersByUid(uid string) ([]domain.User, error) {
	user, _ := i.userRepository.GetUserByUid(uid)
	gender := getGenderForSearch(user.IsMale)
	users, err := i.userRepository.GetUsersByGender(gender)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (i *userUsecase) GetUserByUid(uid string, visitorUid string) (*domain.User, error) {
	user, err := i.userRepository.GetUserByUid(uid)
	if err != nil {
		return nil, err
	}
	if err := i.footPrintRepository.CreateFootPrint(uid, visitorUid); err != nil {
		return nil, err
	}
	return user, nil
}

func (i *userUsecase) CreateUser(user *domain.User) (*domain.User, error) {
	user, err := i.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func getGenderForSearch(isMale bool) bool {
	// 仮でstruct持たせても良さそう！分かりやすいし。性別とisMaleを持った
	if isMale {
		return false
	} else {
		return true
	}
}
