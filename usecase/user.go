package usecase

import (
	"finder/domain"
	"finder/infrastructure/repository"
)

type UserUsecase interface {
	GetUsersByUid(uid string) ([]domain.User, error)
	GetUserByUid(uid string, visitorUid string) (domain.User, error)
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
	genderToSearch := getGenderForSearch(user.Gender)
	return i.userRepository.GetUsersByGender(genderToSearch)
}

func (i *userUsecase) GetUserByUid(uid string, visitorUid string) (domain.User, error) {
	if err := i.footPrintRepository.CreateFootPrint(uid, visitorUid); err != nil {
		return domain.User{}, err
	}
	return i.userRepository.GetUserByUid(uid)
}

func (i *userUsecase) CreateUser(user *domain.User) (*domain.User, error) {
	return i.userRepository.CreateUser(user)
}

func getGenderForSearch(userGender string) string {
	male, female := "男性", "女性"
	switch userGender {
	case male:
		return female
	case female:
		return male
	default:
		panic("不正な値です")
	}
}
