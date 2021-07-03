package usecase

import (
	"context"
	"grpc/interface/converter"
	"grpc/pb"
	"grpc/repository"
)

type UserUsecase interface {
	GetUsers(ctx context.Context, req *pb.GetUsersReq) (*pb.GetUsersRes, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUseuserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) GetUsers(ctx context.Context, req *pb.GetUsersReq) (*pb.GetUsersRes, error) {
	user, _ := u.userRepository.GetUserByUid(req.Uid)
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
