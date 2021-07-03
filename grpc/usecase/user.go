package usecase

import (
	"context"
	"errors"
	"grpc/domain"
	"grpc/interface/converter"
	"grpc/pb"
	"grpc/repository"
)

type UserUsecase interface {
	GetUsers(ctx context.Context, req *pb.GetUsersReq) (*pb.GetUsersRes, error)
	GetUserByUid(ctx context.Context, req *pb.GetUserByUidReq) (*pb.GetUserByUidRes, error)
}

type userUsecase struct {
	userRepository      repository.UserRepository
	footPrintRepository repository.FootPrintRepository
}

func NewUserUseuserUsecase(
	userRepository repository.UserRepository,
	footPrintRepository repository.FootPrintRepository,
) UserUsecase {
	return &userUsecase{
		userRepository:      userRepository,
		footPrintRepository: footPrintRepository,
	}
}

func (u *userUsecase) GetUsers(ctx context.Context, req *pb.GetUsersReq) (*pb.GetUsersRes, error) {
	user, _ := u.userRepository.GetUserByUid(req.CurrentUserUid)
	genderToSearch := getGenderForSearch(user.Gender)
	users, err := u.userRepository.GetUsersByGender(genderToSearch)
	if err != nil {
		return nil, err
	}
	return &pb.GetUsersRes{
		Users: converter.ConvertUsers(users),
	}, nil
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

func (u *userUsecase) GetUserByUid(ctx context.Context, req *pb.GetUserByUidReq) (*pb.GetUserByUidRes, error) {
	user, err := u.userRepository.GetUserByUid(req.Uid)
	if err != nil {
		return nil, err
	}
	visitor, err := u.userRepository.GetUserByVisitorUid(req.VisitorUid)
	if err != nil {
		return nil, err
	}
	if user.Gender == visitor.Gender {
		return nil, errors.New("該当するユーザーは表示できません")
	}

	footPrint := &domain.FootPrint{
		VisitorUid: req.VisitorUid,
		UserUid:    req.Uid,
		Unread:     true,
	}
	if err := u.footPrintRepository.CreateFootPrint(footPrint); err != nil {
		return nil, err
	}
	return &pb.GetUserByUidRes{
		User: converter.ConvertUser(user),
	}, nil
}
