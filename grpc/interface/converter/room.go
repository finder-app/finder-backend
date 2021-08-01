package converter

import (
	"grpc/domain"
	"grpc/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertRoom(room *domain.Room) *pb.Room {
	pbRoom := &pb.Room{
		Id:        room.Id,
		CreatedAt: timestamppb.New(room.CreatedAt),
		UpdatedAt: timestamppb.New(room.UpdatedAt),
	}
	return pbRoom
}
