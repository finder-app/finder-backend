package usecase

import (
	"grpc/domain"
)

type RoomUsecase interface {
	GetRooms(currentRoomUid string) ([]*domain.Room, error)
}

type roomUsecase struct {
	// userRepository repository.RoomRepository
}

// func NewRoomUsecase(roomRepository repository.RoomRepository) RoomUsecase {
func NewRoomUsecase() RoomUsecase {
	return &roomUsecase{
		// userRepository: userRepository,
	}
}

func (u *roomUsecase) GetRooms(currentRoomUid string) ([]*domain.Room, error) {
	// return u.userRepository.GetRoomByUid(currentRoomUid)
	return nil, nil
}
