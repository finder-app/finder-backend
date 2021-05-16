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

func (u *userUsecase) GetUsersByUid(uid string) ([]domain.User, error) {
	user, _ := u.userRepository.GetUserByUid(uid)
	genderToSearch := getGenderForSearch(user.Gender)
	return u.userRepository.GetUsersByGender(genderToSearch)
}

func (u *userUsecase) GetUserByUid(uid string, visitorUid string) (domain.User, error) {
	if err := u.footPrintRepository.CreateFootPrint(uid, visitorUid); err != nil {
		return domain.User{}, err
	}
	return u.userRepository.GetUserByUid(uid)
}

func (u *userUsecase) CreateUser(user *domain.User) (*domain.User, error) {
	return u.userRepository.CreateUser(user)
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
