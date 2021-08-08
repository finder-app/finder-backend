package usecase

import (
	"errors"

	"github.com/finder-app/finder-backend/grpc/domain"
	"github.com/finder-app/finder-backend/grpc/repository"
)

type UserUsecase interface {
	GetUsers(currentUserUid string) ([]*domain.User, error)
	GetUserByUid(uid string, visitorUid string) (*domain.User, error)
	CreateUser(inputUser *domain.User) (*domain.User, error)
}

type userUsecase struct {
	userRepository      repository.UserRepository
	footPrintRepository repository.FootPrintRepository
	likeRepository      repository.LikeRepository
}

func NewUserUsecase(
	userRepository repository.UserRepository,
	footPrintRepository repository.FootPrintRepository,
	likeRepository repository.LikeRepository,
) UserUsecase {
	return &userUsecase{
		userRepository:      userRepository,
		footPrintRepository: footPrintRepository,
		likeRepository:      likeRepository,
	}
}

func (u *userUsecase) GetUsers(currentUserUid string) ([]*domain.User, error) {
	user, _ := u.userRepository.GetUserByUid(currentUserUid)
	genderToSearch := getGenderForSearch(user.Gender)
	return u.userRepository.GetUsersByGender(genderToSearch)
}

// NOTE: 男性なら女性を、女性なら男性のユーザーを検索するように
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

func (u *userUsecase) GetUserByUid(uid string, visitorUid string) (*domain.User, error) {
	// NOTE: IDから取得したユーザーとログイン中のユーザーの性別が同じならエラーを返す
	// HACK: mockテストを通すために、GetUserのメソッドを2つ用意している。修正したい
	user, err := u.userRepository.GetUserByUid(uid)
	if err != nil {
		return nil, err
	}
	visitor, err := u.userRepository.GetUserByVisitorUid(visitorUid)
	if err != nil {
		return nil, err
	}
	if user.Gender == visitor.Gender {
		return nil, errors.New("該当するユーザーは表示できません")
	}

	// NOTE: 足跡が存在しなければ新しく作成
	footPrint := &domain.FootPrint{
		VisitorUid: visitorUid,
		UserUid:    uid,
		Unread:     true,
	}
	if err := u.footPrintRepository.CreateFootPrint(footPrint); err != nil {
		return nil, err
	}

	// NOTE: 取得したユーザーがいいね済みか取得
	if err := u.likeRepository.Liked(user, visitor.Uid); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) CreateUser(inputUser *domain.User) (*domain.User, error) {
	// TODO: validationを追加したい
	// if err := validations.ValidateUser(user); err != nil {
	// 	return nil, err
	// }
	return u.userRepository.CreateUser(inputUser)
}
