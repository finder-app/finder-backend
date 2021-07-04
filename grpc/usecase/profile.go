package usecase

import (
	"context"
	"grpc/domain"
	"grpc/interface/converter"
	"grpc/pb"
	"grpc/repository"
)

type ProfileUsecase interface {
	GetProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.GetProfileRes, error)
	UpdateProfile(ctx context.Context, req *pb.UpdateProfileReq) (*pb.UpdateProfileRes, error)
}

type profileUsecase struct {
	userRepository repository.UserRepository
}

func NewProfileUsecase(userRepository repository.UserRepository) ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
	}
}

func (u *profileUsecase) GetProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.GetProfileRes, error) {
	user, err := u.userRepository.GetUserByUid(req.CurrentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.GetProfileRes{
		User: converter.ConvertUser(user),
	}, nil
}

func (u *profileUsecase) UpdateProfile(ctx context.Context, req *pb.UpdateProfileReq) (*pb.UpdateProfileRes, error) {
	// NOTE: 変更して良いカラムだけ書く uidは更新時にwhereするのに必要
	inputUser := &domain.User{
		Uid:       req.User.Uid,
		LastName:  req.User.LastName,
		FirstName: req.User.FirstName,
	}
	// TODO: update前にvalidationしたい
	user, err := u.userRepository.UpdateUser(inputUser)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProfileRes{
		User: converter.ConvertUser(user),
	}, nil
}
