package converter

import (
	"github.com/finder-app/finder-backend/grpc/domain"
	"github.com/finder-app/finder-backend/grpc/finder-protocol-buffers/pb"

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
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		Gender:    user.Gender,
		FullName:  user.FullName(),
		Thumbnail: user.Thumbnail,

		// NOTE: カラムを持たないfield
		Liked: user.Liked,
	}
	return pbUser
}
