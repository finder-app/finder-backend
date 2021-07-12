package usecase

import (
	"context"
	"grpc/domain"
	"grpc/interface/converter"
	"grpc/pb"
	"grpc/repository"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LikeUsecase interface {
	CreateLike(ctx context.Context, req *pb.CreateLikeReq) (*pb.CreateLikeRes, error)
	GetOldestLike(ctx context.Context, req *pb.GetOldestLikeReq) (*pb.GetOldestLikeRes, error)
	Skip(ctx context.Context, req *pb.SkipReq) (*empty.Empty, error)
	// Consent(recievedUserUid string, sentUesrUid string) (domain.Like, domain.Room, error)
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

func (u *likeUsecase) GetOldestLike(ctx context.Context, req *pb.GetOldestLikeReq) (*pb.GetOldestLikeRes, error) {
	currentUserUid := req.CurrentUserUid
	like, err := u.likeRepository.GetOldestLikeByUid(currentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.GetOldestLikeRes{
		Like: converter.ConvertLike(like),
	}, nil
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

// func (u *likeUsecase) Consent(recievedUserUid string, sentUesrUid string) (domain.Like, domain.Room, error) {
// 	tx := u.likeRepository.Begin()
// 	like := domain.Like{
// 		RecievedUserUid: recievedUserUid,
// 		SentUserUid:     sentUesrUid,
// 		Consented:       true,
// 	}
// 	if err := u.likeRepository.Consent(tx, &like); err != nil {
// 		tx.Rollback()
// 		return domain.Like{}, domain.Room{}, err
// 	}
// 	room := domain.Room{}
// 	if err := u.roomRepository.CreateRoom(tx, &room); err != nil {
// 		tx.Rollback()
// 		return domain.Like{}, domain.Room{}, err
// 	}
// 	roomUser1 := domain.RoomUser{
// 		RoomId:  room.Id,
// 		UserUid: recievedUserUid,
// 	}
// 	if err := u.roomUserRepository.CreateRoomUser(tx, roomUser1); err != nil {
// 		tx.Rollback()
// 		return domain.Like{}, domain.Room{}, err
// 	}
// 	roomUser2 := domain.RoomUser{
// 		RoomId:  room.Id,
// 		UserUid: sentUesrUid,
// 	}
// 	if err := u.roomUserRepository.CreateRoomUser(tx, roomUser2); err != nil {
// 		tx.Rollback()
// 		return domain.Like{}, domain.Room{}, err
// 	}
// 	tx.Commit()
// 	return like, room, nil
// }
