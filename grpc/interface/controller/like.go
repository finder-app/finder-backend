package controller

import (
	"context"
	"grpc/interface/converter"
	"grpc/pb"
	"grpc/usecase"

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
	return nil, nil
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

func (c *LikeController) Skip(ctx context.Context, req *pb.SkipReq) (*emptypb.Empty, error) {
	return nil, nil
}

func (c *LikeController) ConsentLike(ctx context.Context, req *pb.ConsentLikeReq) (*pb.ConsentLikeRes, error) {
	recievedUserUid := req.RecievedUserUid
	sentUserUid := req.SentUserUid
	like, _, err := c.likeUsecase.Consent(recievedUserUid, sentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.ConsentLikeRes{
		Like: converter.ConvertLike(like),
		Room: converter.ConvertLike(like),
	}, nil
}
