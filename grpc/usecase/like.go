package usecase

import (
	"grpc/domain"
	"grpc/repository"
	"grpc/usecase/validation"
)

type LikeUsecase interface {
	CreateLike(like *domain.Like) (*domain.Like, error)
	GetOldestLike(currentUserUid string) (*domain.Like, error)
	Skip(like *domain.Like) error
	Consent(recievedUserUid string, sentUesrUid string) (*domain.Like, *domain.Room, error)
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
) LikeUsecase {
	return &likeUsecase{
		likeRepository:     lr,
		roomRepository:     rr,
		roomUserRepository: rur,
	}
}

func (u *likeUsecase) CreateLike(like *domain.Like) (*domain.Like, error) {
	return u.likeRepository.CreateLike(like)
}

func (u *likeUsecase) GetOldestLike(currentUserUid string) (*domain.Like, error) {
	return u.likeRepository.GetOldestLikeByUid(currentUserUid)
}

func (u *likeUsecase) Skip(like *domain.Like) error {
	return u.likeRepository.Skip(like)
}

func (u *likeUsecase) Consent(recievedUserUid string, sentUesrUid string) (*domain.Like, *domain.Room, error) {
	like := &domain.Like{
		RecievedUserUid: recievedUserUid,
		SentUserUid:     sentUesrUid,
	}
	// NOTE: 既存のいいねを取得していいね済み or スキップ済みか確認する 返り値がtrueならerrorをreturnする
	if err := u.likeRepository.GetLike(like); err != nil {
		return nil, nil, err
	}
	if err := validation.ValidateLike(like); err != nil {
		return nil, nil, err
	}

	// NOTE: いいねを行う
	tx := u.likeRepository.Begin()
	if err := u.likeRepository.Consent(tx, like); err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	// NOTE: roomを作成
	room := domain.NewRoom()
	if err := u.roomRepository.CreateRoom(tx, room); err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	// NOTE: 作成したroomを元に男女のroomUserを作成
	roomUser1, roomUser2 := initializeRoomUsers(room, like)
	if err := u.roomUserRepository.CreateRoomUsers(tx, roomUser1, roomUser2); err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	tx.Commit()
	return like, room, nil
}

func initializeRoomUsers(room *domain.Room, like *domain.Like) (*domain.RoomUser, *domain.RoomUser) {
	roomUser1 := &domain.RoomUser{
		RoomId:  room.Id,
		UserUid: like.RecievedUserUid,
	}
	roomUser2 := &domain.RoomUser{
		RoomId:  room.Id,
		UserUid: like.SentUserUid,
	}
	return roomUser1, roomUser2
}
