package usecase

import (
	"grpc/domain"
	"grpc/repository"
)

type RoomUsecase interface {
	GetRooms(currentUserUid string) ([]*domain.Room, error)
}

type roomUsecase struct {
	roomRepository repository.RoomRepository
}

func NewRoomUsecase(roomRepository repository.RoomRepository) RoomUsecase {
	return &roomUsecase{
		roomRepository: roomRepository,
	}
}

func (u *roomUsecase) GetRooms(currentUserUid string) ([]*domain.Room, error) {
	return u.roomRepository.GetRooms(currentUserUid)
}
