package controller

import (
	"context"
	"grpc/finder-protocol-buffers/pb"
	"grpc/interface/converter"
	"grpc/usecase"
)

type RoomController struct {
	roomUsecase usecase.RoomUsecase
}

func NewRoomController(roomUsecase usecase.RoomUsecase) *RoomController {
	return &RoomController{
		roomUsecase: roomUsecase,
	}
}

func (c *RoomController) GetRooms(ctx context.Context, req *pb.GetRoomsReq) (*pb.GetRoomsRes, error) {
	rooms, err := c.roomUsecase.GetRooms(req.CurrentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.GetRoomsRes{
		Rooms: converter.ConvertRooms(rooms),
	}, nil
}
