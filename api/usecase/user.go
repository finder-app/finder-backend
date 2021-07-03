package usecase

import (
	"finder/domain"
	"finder/infrastructure/repository"
)

type UserUsecase interface {
	// GetUsersByUid(uid string) ([]*domain.User, error)
	// GetUserByUid(uid string, visitorUid string) (*domain.User, error)
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

// func (u *userUsecase) GetUsersByUid(uid string) ([]*domain.User, error) {
// 	user, _ := u.userRepository.GetUserByUid(uid)
// 	genderToSearch := getGenderForSearch(user.Gender)
// 	return u.userRepository.GetUsersByGender(genderToSearch)
// }

// func (u *userUsecase) GetUserByUid(uid string, visitorUid string) (*domain.User, error) {
// 	user, err := u.userRepository.GetUserByUid(uid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	visitor, err := u.userRepository.GetUserByVisitorUid(visitorUid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if user.Gender == visitor.Gender {
// 		return nil, errors.New("該当するユーザーは表示できません")
// 	}

// 	footPrint := &domain.FootPrint{
// 		VisitorUid: visitorUid,
// 		UserUid:    uid,
// 		Unread:     true,
// 	}
// 	if err := u.footPrintRepository.CreateFootPrint(footPrint); err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

func (u *userUsecase) CreateUser(user *domain.User) (*domain.User, error) {
	return u.userRepository.CreateUser(user)
}

// NOTE: 男性なら女性を、女性なら男性のユーザーを検索するように
// func getGenderForSearch(userGender string) string {
// 	male, female := "男性", "女性"
// 	switch userGender {
// 	case male:
// 		return female
// 	case female:
// 		return male
// 	default:
// 		panic("不正な値です")
// 	}
// }
