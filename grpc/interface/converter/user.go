package converter

import (
	"grpc/domain"
	"grpc/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertUsers(users []*domain.User) []*pb.User {
	var pbUsers []*pb.User
	for _, user := range users {
		pbUser := ConvertUser(user)
		pbUsers = append(pbUsers, pbUser)
	}
	return pbUsers
}

func ConvertUser(user *domain.User) *pb.User {
	pbUser := &pb.User{
		Uid:       user.Uid,
		Email:     user.Email,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Gender:    user.Gender,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
	return pbUser
}
