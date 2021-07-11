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
	CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error)
}

type userUsecase struct {
	userRepository      repository.UserRepository
	footPrintRepository repository.FootPrintRepository
	likeRepository      repository.LikeRepository
}

func NewUserUseuserUsecase(
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
	// NOTE: IDから取得したユーザーとログイン中のユーザーの性別が同じならエラーを返す
	// HACK: mockテストを通すために、GetUserのメソッドを2つ用意している。修正したい
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

	// NOTE: 足跡が存在しなければ新しく作成
	footPrint := &domain.FootPrint{
		VisitorUid: req.VisitorUid,
		UserUid:    req.Uid,
		Unread:     true,
	}
	if err := u.footPrintRepository.CreateFootPrint(footPrint); err != nil {
		return nil, err
	}

	// NOTE: 取得したユーザーがいいね済みか取得
	if err := u.likeRepository.Liked(user, visitor.Uid); err != nil {
		return nil, err
	}

	return &pb.GetUserByUidRes{
		User: converter.ConvertUser(user),
	}, nil
}

func (u *userUsecase) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	// TODO: validationを追加したい
	// if err := validations.ValidateUser(user); err != nil {
	// 	return nil, err
	// }
	inputUser := &domain.User{
		Uid:       req.User.Uid,
		Email:     req.User.Email,
		LastName:  req.User.LastName,
		FirstName: req.User.FirstName,
		Gender:    req.User.Gender,
	}
	user, err := u.userRepository.CreateUser(inputUser)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserRes{
		User: converter.ConvertUser(user),
	}, nil
}
