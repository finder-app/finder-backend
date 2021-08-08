package converter

import (
	"github.com/finder-app/finder-backend/grpc/domain"
	"github.com/finder-app/finder-backend/grpc/finder-protocol-buffers/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertRooms(rooms []*domain.Room) []*pb.Room {
	var pbRooms []*pb.Room
	for _, room := range rooms {
		pbRoom := ConvertRoom(room)
		pbRooms = append(pbRooms, pbRoom)
	}
	return pbRooms
}

func ConvertRoom(room *domain.Room) *pb.Room {
	pbRoom := &pb.Room{
		Id:        room.Id,
		CreatedAt: timestamppb.New(room.CreatedAt),
		UpdatedAt: timestamppb.New(room.UpdatedAt),
		// LastMessage: room.LastMessage,
	}
	return pbRoom
}
