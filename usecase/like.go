package usecase

import (
	"finder/domain"
	"finder/infrastructure/repository"
)

type LikeUsecase interface {
	CreateLike(sentUesrUid string, recievedUserUid string) (*domain.Like, error)
	GetOldestLikeByUid(currentUserUid string) (*domain.Like, error)
	GetNextUserByUid(recievedUserUid string, sentUesrUid string) (*domain.Like, error)
	Consent(recievedUserUid string, sentUesrUid string) error
}

type likeUsecase struct {
	likeRepository     repository.LikeRepository
	roomRepository     repository.RoomRepository
	roomUserRepository repository.RoomUserRepository
}

func NewLikeUsecase(
	lr repository.LikeRepository,
	rr repository.RoomRepository,
	rur repository.RoomUserRepository,
) *likeUsecase {
	return &likeUsecase{
		likeRepository:     lr,
		roomRepository:     rr,
		roomUserRepository: rur,
	}
}

func (u *likeUsecase) CreateLike(sentUesrUid string, recievedUserUid string) (*domain.Like, error) {
	like, err := u.likeRepository.CreateLike(sentUesrUid, recievedUserUid)
	if err != nil {
		return nil, err
	}
	return like, nil
}

func (u *likeUsecase) GetOldestLikeByUid(currentUserUid string) (*domain.Like, error) {
	like, err := u.likeRepository.GetOldestLikeByUid(currentUserUid)
	if err != nil {
		return nil, err
	}
	return like, nil
}

func (u *likeUsecase) GetNextUserByUid(recievedUserUid string, sentUesrUid string) (*domain.Like, error) {
	if err := u.likeRepository.NopeUserByUid(recievedUserUid, sentUesrUid); err != nil {
		return nil, err
	}
	like, err := u.likeRepository.GetOldestLikeByUid(recievedUserUid)
	if err != nil {
		return nil, err
	}
	return like, nil
}

func (u *likeUsecase) Consent(recievedUserUid string, sentUesrUid string) error {
	tx := u.likeRepository.Begin()
	if err := u.likeRepository.Consent(tx, recievedUserUid, sentUesrUid); err != nil {
		tx.Rollback()
		return err
	}
	room := domain.Room{}
	if err := u.roomRepository.CreateRoom(tx, &room); err != nil {
		tx.Rollback()
		return err
	}
	roomUser1 := domain.RoomUser{
		RoomId:  room.Id,
		UserUid: recievedUserUid,
	}
	if err := u.roomUserRepository.CreateRoomUser(tx, roomUser1); err != nil {
		tx.Rollback()
		return err
	}
	roomUser2 := domain.RoomUser{
		RoomId:  room.Id,
		UserUid: sentUesrUid,
	}
	if err := u.roomUserRepository.CreateRoomUser(tx, roomUser2); err != nil {
		tx.Rollback()
		return err
	}
	// tx.Rollback()
	tx.Commit()
	// roomIdをresponseにかえして、部屋に移動できるように！
	return nil
}
