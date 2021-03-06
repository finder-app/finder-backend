package controller

import (
	"context"

	"github.com/finder-app/finder-backend/grpc/domain"
	"github.com/finder-app/finder-backend/grpc/finder-protocol-buffers/pb"
	"github.com/finder-app/finder-backend/grpc/interface/converter"
	"github.com/finder-app/finder-backend/grpc/usecase"

	"google.golang.org/protobuf/types/known/emptypb"
)

type LikeController struct {
	likeUsecase usecase.LikeUsecase
}

func NewLikeController(likeUsecase usecase.LikeUsecase) *LikeController {
	return &LikeController{
		likeUsecase: likeUsecase,
	}
}

func (c *LikeController) CreateLike(ctx context.Context, req *pb.CreateLikeReq) (*pb.CreateLikeRes, error) {
	like := &domain.Like{
		SentUserUid:     req.SentUserUid,
		RecievedUserUid: req.RecievedUserUid,
	}
	if _, err := c.likeUsecase.CreateLike(like); err != nil {
		return nil, err
	}
	return &pb.CreateLikeRes{
		Like: converter.ConvertLike(like),
	}, nil
}

func (c *LikeController) GetOldestLike(ctx context.Context, req *pb.GetOldestLikeReq) (*pb.GetOldestLikeRes, error) {
	currentUserUid := req.CurrentUserUid
	like, err := c.likeUsecase.GetOldestLike(currentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.GetOldestLikeRes{
		Like: converter.ConvertLike(like),
	}, nil
}

// SkipLikeにしたい
func (c *LikeController) SkipLike(ctx context.Context, req *pb.SkipLikeReq) (*emptypb.Empty, error) {
	like := &domain.Like{
		SentUserUid:     req.SentUserUid,
		RecievedUserUid: req.RecievedUserUid,
	}
	if err := c.likeUsecase.Skip(like); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (c *LikeController) ConsentLike(ctx context.Context, req *pb.ConsentLikeReq) (*pb.ConsentLikeRes, error) {
	recievedUserUid := req.RecievedUserUid
	sentUserUid := req.SentUserUid
	like, room, err := c.likeUsecase.Consent(recievedUserUid, sentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.ConsentLikeRes{
		Like: converter.ConvertLike(like),
		Room: converter.ConvertRoom(room),
	}, nil
}
