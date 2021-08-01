package usecase

import (
	"grpc/domain"
	"grpc/repository"
)

type MessageUsecase interface {
	GetMessages(roomId uint64, currentUserUid string) ([]*domain.Message, error)
}

type messageUsecase struct {
	messageRepository  repository.MessageRepository
	roomUserRepository repository.RoomUserRepository
}

func NewMessageUsecase(
	messageRepository repository.MessageRepository,
	roomUserRepository repository.RoomUserRepository,
) MessageUsecase {
	return &messageUsecase{
		messageRepository:  messageRepository,
		roomUserRepository: roomUserRepository,
	}
}

func (u *messageUsecase) GetMessages(roomId uint64, currentUserUid string) ([]*domain.Message, error) {
	if err := u.validationGetMessages(roomId, currentUserUid); err != nil {
		return nil, err
	}

	return u.messageRepository.GetMessages(roomId)
}

// NOTE: roomIdにroomUserが含まれているか確認する。存在しなければerrorを返す
func (u *messageUsecase) validationGetMessages(roomId uint64, currentUserUid string) error {
	_, err := u.roomUserRepository.GetRoomUser(roomId, currentUserUid)
	if err != nil {
		return err
	}
	return nil
}
