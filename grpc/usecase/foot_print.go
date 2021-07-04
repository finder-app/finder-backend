package usecase

import (
	"context"
	"grpc/interface/converter"
	"grpc/pb"
	"grpc/repository"
)

type FootPrintUsecase interface {
	GetFootPrints(ctx context.Context, req *pb.GetFootPrintsReq) (*pb.GetFootPrintsRes, error)
	GetUnreadCount(ctx context.Context, req *pb.GetUnreadCountReq) (*pb.GetUnreadCountRes, error)
}

type footPrintUsecase struct {
	footPrintRepository repository.FootPrintRepository
}

func NewFootPrintUsecase(footPrintRepository repository.FootPrintRepository) FootPrintUsecase {
	return &footPrintUsecase{
		footPrintRepository: footPrintRepository,
	}
}

func (u *footPrintUsecase) GetFootPrints(ctx context.Context, req *pb.GetFootPrintsReq) (*pb.GetFootPrintsRes, error) {
	if err := u.footPrintRepository.UpdateToAlreadyRead(req.CurrentUserUid); err != nil {
		return nil, err
	}
	footPrints, err := u.footPrintRepository.GetFootPrintsByUid(req.CurrentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.GetFootPrintsRes{
		FootPrints: converter.ConvertFootPrints(footPrints),
	}, nil
}

func (u *footPrintUsecase) GetUnreadCount(ctx context.Context, req *pb.GetUnreadCountReq) (*pb.GetUnreadCountRes, error) {
	unreadCount, err := u.footPrintRepository.GetUnreadCount(req.CurrentUserUid)
	if err != nil {
		return nil, err
	}
	return &pb.GetUnreadCountRes{
		UnreadCount: int64(unreadCount),
	}, nil
}
