package controller

import (
	"context"
	"grpc/domain"
	"grpc/finder-protocol-buffers/pb"
	"grpc/interface/converter"
	"grpc/usecase"
)

type ProfileController struct {
	profileUsecase usecase.ProfileUsecase
}

func NewProfileController(profileUsecase usecase.ProfileUsecase) *ProfileController {
	return &ProfileController{
		profileUsecase: profileUsecase,
	}
}

func (c *ProfileController) GetProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.GetProfileRes, error) {
	user, err := c.profileUsecase.GetProfile(req.CurrentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.GetProfileRes{
		User: converter.ConvertUser(user),
	}, nil
}

func (c *ProfileController) UpdateProfile(ctx context.Context, req *pb.UpdateProfileReq) (*pb.UpdateProfileRes, error) {
	// NOTE: 変更して良いカラムだけ書く uidは更新時にwhereするのに必要
	inputUser := &domain.User{
		Uid:       req.User.Uid,
		LastName:  req.User.LastName,
		FirstName: req.User.FirstName,
		Thumbnail: req.User.Thumbnail,
	}
	user, err := c.profileUsecase.UpdateProfile(inputUser)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProfileRes{
		User: converter.ConvertUser(user),
	}, nil
}
