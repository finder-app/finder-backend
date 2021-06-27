package usecase

import (
	"finder/domain"
	"finder/infrastructure/repository"
)

type LikeUsecase interface {
	CreateLike(sentUesrUid string, recievedUserUid string) (*domain.Like, error)
	GetOldestLikeByUid(currentUserUid string) (*domain.Like, error)
	GetNextUserByUid(recievedUserUid string, sentUesrUid string) (*domain.Like, error)
	Consent(recievedUserUid string, sentUesrUid string) (domain.Like, domain.Room, error)
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
	like := domain.Like{
		SentUserUid:     sentUesrUid,
		RecievedUserUid: recievedUserUid,
	}
	return u.likeRepository.CreateLike(&like)
}

func (u *likeUsecase) GetOldestLikeByUid(currentUserUid string) (*domain.Like, error) {
	return u.likeRepository.GetOldestLikeByUid(currentUserUid)
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

func (u *likeUsecase) Consent(recievedUserUid string, sentUesrUid string) (domain.Like, domain.Room, error) {
	tx := u.likeRepository.Begin()
	like := domain.Like{
		RecievedUserUid: recievedUserUid,
		SentUserUid:     sentUesrUid,
		Consented:       true,
	}
	if err := u.likeRepository.Consent(tx, &like); err != nil {
		tx.Rollback()
		return domain.Like{}, domain.Room{}, err
	}
	room := domain.Room{}
	if err := u.roomRepository.CreateRoom(tx, &room); err != nil {
		tx.Rollback()
		return domain.Like{}, domain.Room{}, err
	}
	roomUser1 := domain.RoomUser{
		RoomId:  room.Id,
		UserUid: recievedUserUid,
	}
	if err := u.roomUserRepository.CreateRoomUser(tx, roomUser1); err != nil {
		tx.Rollback()
		return domain.Like{}, domain.Room{}, err
	}
	roomUser2 := domain.RoomUser{
		RoomId:  room.Id,
		UserUid: sentUesrUid,
	}
	if err := u.roomUserRepository.CreateRoomUser(tx, roomUser2); err != nil {
		tx.Rollback()
		return domain.Like{}, domain.Room{}, err
	}
	tx.Commit()
	return like, room, nil
}
