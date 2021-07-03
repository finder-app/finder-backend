package usecase

import (
	"context"
	"grpc/pb"
)

type userUsecase struct {
}

func NewUserUseuserUsecase() *userUsecase {
	return &userUsecase{}
}

func (c *userUsecase) GetUsers(ctx context.Context, req *pb.GetUsersReq) (*pb.GetUsersRes, error) {
	var pbUsers []*pb.User
	return &pb.GetUsersRes{
		Users: pbUsers,
	}, nil
}
