package usecase

import (
	"context"
	"grpc/domain"
	"grpc/interface/converter"
	"grpc/pb"
	"grpc/repository"
	"grpc/usecase/validation"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LikeUsecase interface {
	CreateLike(ctx context.Context, req *pb.CreateLikeReq) (*pb.CreateLikeRes, error)
	GetOldestLike(currentUserUid string) (*domain.Like, error)
	Skip(ctx context.Context, req *pb.SkipReq) (*empty.Empty, error)
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

func (u *likeUsecase) CreateLike(ctx context.Context, req *pb.CreateLikeReq) (*pb.CreateLikeRes, error) {
	like := &domain.Like{
		SentUserUid:     req.SentUserUid,
		RecievedUserUid: req.RecievedUserUid,
	}
	if _, err := u.likeRepository.CreateLike(like); err != nil {
		return nil, err
	}
	return &pb.CreateLikeRes{
		Like: converter.ConvertLike(like),
	}, nil
}

func (u *likeUsecase) GetOldestLike(currentUserUid string) (*domain.Like, error) {
	return u.likeRepository.GetOldestLikeByUid(currentUserUid)
}

func (u *likeUsecase) Skip(ctx context.Context, req *pb.SkipReq) (*empty.Empty, error) {
	like := &domain.Like{
		SentUserUid:     req.SentUserUid,
		RecievedUserUid: req.RecievedUserUid,
	}
	if err := u.likeRepository.Skip(like); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
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

	tx := u.likeRepository.Begin()
	if err := u.likeRepository.Consent(tx, like); err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	room := &domain.Room{}
	if err := u.roomRepository.CreateRoom(tx, room); err != nil {
		tx.Rollback()
		return nil, nil, err
	}
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
